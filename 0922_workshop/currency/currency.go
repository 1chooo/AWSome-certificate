package main

import (
	"bytes"
	"fmt"
	"image/color"
	"image/png"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-redis/redis/v8"
	"github.com/PuerkitoBio/goquery"
	"github.com/fogleman/gg"
)

// S3BucketName 存儲 S3 存儲桶名稱
const S3BucketName = "your-bucket-name"

// Redis 地址和端口
const RedisAddress = "your-redis-endpoint:6379"

func handleRequest(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// 從 API Gateway 事件中取出 day 參數，預設為 "day"
	day := event.QueryStringParameters["day"]
	if day == "" {
		day = "day"
	}

	url := fmt.Sprintf("https://rate.bot.com.tw/xrt/flcsv/0/%s", day)

	var dayString string
	if day == "day" {
		dayString = "Today"
	} else {
		dayString = day
	}

	// 發送網路請求獲取匯率數據
	response, err := http.Get(url)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, err
	}
	defer response.Body.Close()

	rateData := response.Body

	// 解析 CSV 格式的匯率數據
	currencyRates := make(map[string]float64)
	scanner := bufio.NewScanner(rateData)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) < 13 {
			continue
		}
		currency := parts[0]
		rateStr := strings.TrimSpace(parts[12])
		rate, err := strconv.ParseFloat(rateStr, 64)
		if err != nil {
			continue
		}
		currencyRates[currency] = rate
	}

	// 將匯率按值排序
	sortedRates := make(map[string]float64)
	var currencies []string
	for currency, rate := range currencyRates {
		currencies = append(currencies, currency)
		sortedRates[currency] = rate
	}
	sort.Slice(currencies, func(i, j int) bool {
		return sortedRates[currencies[i]] < sortedRates[currencies[j]]
	})

	dc := gg.NewContext(400, 300)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	labels := make([]string, len(currencies))
	for i, currency := range currencies {
		labels[i] = currency
	}
	rates := make([]float64, len(currencies))
	for i, currency := range currencies {
		rates[i] = sortedRates[currency]
	}

	colors := make([]color.Color, len(currencies))
	for i := range colors {
		colors[i] = color.RGBA{0, 0, 255, 255} // 藍色
	}

	dc.SetFontFace(gg.NewFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf", 96))
	for i, rate := range rates {
		dc.SetRGB(0, 0, 0)
		dc.DrawStringAnchored(fmt.Sprintf("%.2f", rate), float64(i*100+50), 200, 0.5, 0.5)
	}

	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored(fmt.Sprintf("Currency Rates\nDate: %s", dayString), 200, 50, 0.5, 0.5)

	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored("Rate", 200, 280, 0.5, 0.5)
	dc.DrawStringAnchored("Currency", 20, 150, 0.5, 0.5)

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
		Key:         aws.String(fmt.Sprintf("currency_rate_chart_%s.png", day)),
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
		Key:    aws.String(fmt.Sprintf("currency_rate_chart_%s.png", day)),
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
