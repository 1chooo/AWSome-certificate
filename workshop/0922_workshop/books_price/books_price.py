import boto3
import matplotlib.pyplot as plt
from io import BytesIO
import requests
import json
from bs4 import BeautifulSoup

# S3 Bucket 名稱
S3_BUCKET_NAME = 'your-bucket-name'


def lambda_handler(event, context):
    # 從 API Gateway 事件中取出 ISBN，若不存在則回傳錯誤
    query_params = event.get('queryStringParameters') or {}
    isbn = query_params.get('isbn')
    
    if isbn is None:
        return {
            'statusCode': 400,
            'body': 'Missing isbn parameter'
        }
    
    # 網站 1
    url1 = f"https://athena.eslite.com/api/v2/search?q={isbn}&final_price=0,&sort=weight&size=20&start=0"
    
    # 網站 2
    url2 = f"https://www.kingstone.com.tw/search/key/{isbn}?"
    
    headers = {
        'User-Agent': 'Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1'
    }
    
    response1 = requests.get(url1, headers=headers)
    data1 = json.loads(response1.text)
    final_price_list1 = [hit["fields"]["final_price"] for hit in data1["hits"]["hit"]][:1]
    
    response2 = requests.get(url2, headers=headers)
    soup = BeautifulSoup(response2.text, 'html.parser')
    price_tags = soup.find_all('span')
    final_price_list2 = []

    for tag in price_tags:
        if tag.find('b') and '特價' in tag.text:
            final_price_list2.append(int(tag.find('b').text))
            
    final_price_list2 = final_price_list2[:1]
    
    
    # 數據視覺化
    labels1 = ["Eslite Price" for i in range(len(final_price_list1))]
    labels2 = ["Kingstone Price" for i in range(len(final_price_list2))]
    labels = labels1 + labels2

    prices1 = [int(price) for price in final_price_list1]
    prices2 = [int(price) for price in final_price_list2]
    prices = prices1 + prices2

    colors = ['blue'] * len(prices1) + ['yellow'] * len(prices2)

    plt.bar(labels, prices, color=colors)
    
    
    # 在每個長條上顯示數字
    for i, price in enumerate(prices):
        plt.text(i, price, str(price), ha='center', va='bottom')

    
    plt.xlabel('Price Index')
    plt.ylabel('Final Price')
    plt.title(f'Final Prices from Multiple Websites\nISBN: {isbn}')
    
    # Save to a BytesIO object
    s3 = boto3.client('s3')
    buffer = BytesIO()
    plt.savefig(buffer, format="png")
    buffer.seek(0)
    
    # Upload to S3
    s3.put_object(Body=buffer, Bucket=S3_BUCKET_NAME, Key=f'price_comparison_{isbn}.png', ContentType='image/png')

    # Generate pre-signed URL
    presigned_url = s3.generate_presigned_url(
        'get_object',
        Params={
            'Bucket': S3_BUCKET_NAME, 
            'Key': f'price_comparison_{isbn}.png'
        },
        ExpiresIn=3600,
    )
                                             
    plt.close()
    
    # Return the pre-signed URL
    return {
        'statusCode': 302,
        'headers': {
            'Location': presigned_url
        }
    }