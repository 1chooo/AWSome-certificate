package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/color"
	"image/png"
	"net/http"
	// "net/url"
	// "os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/fogleman/gg"
)

// S3BucketName 存儲 S3 存儲桶名稱
const S3BucketName = "your-bucket-name"

func handleRequest(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// 從 API Gateway 事件中取出 ISBN，若不存在則回傳錯誤
	isbn := event.QueryStringParameters["isbn"]
	if isbn == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Missing isbn parameter",
		}, nil
	}

	// 網站 1
	url1 := fmt.Sprintf("https://athena.eslite.com/api/v2/search?q=%s&final_price=0,&sort=weight&size=20&start=0", isbn)

	// 網站 2
	url2 := fmt.Sprintf("https://www.kingstone.com.tw/search/key/%s?", isbn)

	// 建立 HTTP Client
	client := &http.Client{}

	// 發送網路請求獲取價格數據
	response1, err := client.Get(url1)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, err
	}
	defer response1.Body.Close()

	var data1 map[string]interface{}
	decoder := json.NewDecoder(response1.Body)
	err = decoder.Decode(&data1)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, err
	}

	hits, ok := data1["hits"].(map[string]interface{})["hit"].([]interface{})
	if !ok {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error parsing response data",
		}, nil
	}

	finalPriceList1 := []float64{}
	for _, hit := range hits {
		fields, ok := hit.(map[string]interface{})["fields"].(map[string]interface{})
		if !ok {
			continue
		}
		finalPrice, ok := fields["final_price"].(float64)
		if ok {
			finalPriceList1 = append(finalPriceList1, finalPrice)
		}
	}

	// 發送網路請求獲取價格數據
	response2, err := client.Get(url2)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, err
	}
	defer response2.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response2.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, err
	}

	finalPriceList2 := []float64{}
	doc.Find("span").Each(func(i int, s *goquery.Selection) {
		b := s.Find("b")
		if strings.Contains(b.Text(), "特價") {
			priceStr := b.Text()
			priceStr = strings.Replace(priceStr, "特價", "", -1)
			priceStr = strings.Replace(priceStr, "元", "", -1)
			price, err := strconv.ParseFloat(priceStr, 64)
			if err == nil {
				finalPriceList2 = append(finalPriceList2, price)
			}
		}
	})

	// 數據視覺化
	dc := gg.NewContext(400, 300)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	labels1 := make([]string, len(finalPriceList1))
	for i := range labels1 {
		labels1[i] = "Eslite Price"
	}
	labels2 := make([]string, len(finalPriceList2))
	for i := range labels2 {
		labels2[i] = "Kingstone Price"
	}
	labels := append(labels1, labels2...)

	prices1 := make([]float64, len(finalPriceList1))
	for i, price := range finalPriceList1 {
		prices1[i] = price
	}
	prices2 := make([]float64, len(finalPriceList2))
	for i, price := range finalPriceList2 {
		prices2[i] = price
	}
	prices := append(prices1, prices2...)

	colors := make([]color.Color, len(prices1)+len(prices2))
	for i := range colors {
		if i < len(prices1) {
			colors[i] = color.RGBA{0, 0, 255, 255} // 藍色
		} else {
			colors[i] = color.RGBA{255, 255, 0, 255} // 黃色
		}
	}

	dc.SetFontFace(gg.NewFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf", 96))
	for i, price := range prices {
		dc.SetRGB(0, 0, 0)
		dc.DrawStringAnchored(fmt.Sprintf("%.0f", price), float64(i*100+50), 200, 0.5, 0.5)
	}

	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored(fmt.Sprintf("Final Prices from Multiple Websites\nISBN: %s", isbn), 200, 50, 0.5, 0.5)

	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored("Price Index", 200, 280, 0.5, 0.5)
	dc.DrawStringAnchored("Final Price", 20, 150, 0.5, 0.5)

	// 儲存圖片到 BytesIO 對象
	var buffer bytes.Buffer
	err = png.Encode(&buffer, dc.Image())
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, err
	}

	// 上傳到 S3
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // 設定您的 AWS 區域
	})
	svc := s3.New(sess)
	_, err = svc.PutObject(&s3.PutObjectInput{
		Body:        bytes.NewReader(buffer.Bytes()),
		Bucket:      aws.String(S3BucketName),
		Key:         aws.String(fmt.Sprintf("price_comparison_%s.png", isbn)),
		ContentType: aws.String("image/png"),
	})
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, err
	}

	// 生成預簽名 URL
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(S3BucketName),
		Key:    aws.String(fmt.Sprintf("price_comparison_%s.png", isbn)),
	})
	presignedURL, err := req.Presign(3600)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, err
	}

	// 回傳預簽名 URL
	return events.APIGatewayProxyResponse{
		StatusCode: 302,
		Headers: map[string]string{
			"Location": presignedURL,
		},
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
