import requests
import matplotlib.pyplot as plt
from io import BytesIO
import boto3

def lambda_handler(event, context):
    query_params = event.get('queryStringParameters') or {}
    day = query_params.get('day', 'day')
    url = f'https://rate.bot.com.tw/xrt/flcsv/0/{day}'
    
    if day == 'day':
        day_string = 'Today'
    else:
        day_string = day
    
    rate = requests.get(url)
    rate.encoding = 'utf-8'
    rt = rate.text
    rts = rt.split('\n')
    
    currency_rates = {}
    for i in rts:
        try:
            a = i.split(',')
            currency = a[0]
            rate = float(a[12])
            currency_rates[currency] = rate
        except (IndexError, ValueError):
            continue
    
    sorted_rates = {k: v for k, v in sorted(currency_rates.items(), key=lambda item: item[1], reverse=False)}
    
    plt.barh(list(sorted_rates.keys()), list(sorted_rates.values()), color='blue')
    plt.title('Currency Rates')
    plt.xlabel('Rate')
    plt.ylabel('Currency')
    plt.annotate(f"Date: {day_string}", xy=(0.85, 1.05), xycoords='axes fraction')  # 新增的行，將日期顯示在右上角
    
    for index, value in enumerate(sorted_rates.values()):
        plt.text(value, index, str(value))

    img_data = BytesIO()
    plt.savefig(img_data, format='png')
    img_data.seek(0)
    
    s3 = boto3.client('s3')
    bucket_name = 'your-bucket-name'
    object_key = f'currency_rate_chart_{day}.png'
    
    s3.put_object(Body=img_data.read(), Bucket=bucket_name, Key=object_key, ContentType='image/png')
    
    presigned_url = s3.generate_presigned_url(
        'get_object',
        Params={
            'Bucket': bucket_name, 
            'Key': object_key
        },
        ExpiresIn=3600
    )
    
    plt.close()  # 在這裡關閉當前的 figure

    return {
        'statusCode': 302,
        'headers': {
            'Location': presigned_url
        },
    }