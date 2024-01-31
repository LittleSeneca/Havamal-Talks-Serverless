import json
import http.client
from datetime import datetime, timedelta
import boto3
import os

def get_vulnerabilities():
    # Adjust the dates to get vulnerabilities from the last month
    end_date = datetime.now()
    start_date = end_date - timedelta(days=30)

    conn = http.client.HTTPSConnection("services.nvd.nist.gov")
    url_path = f"/rest/json/cves/2.0?resultsPerPage=1000&pubStartDate={start_date.strftime('%Y-%m-%d')}T00:00:00.000&pubEndDate={end_date.strftime('%Y-%m-%d')}T00:00:00.000"

    conn.request("GET", url_path)
    response = conn.getresponse()

    if response.status == 200:
        return json.loads(response.read().decode())
    else:
        print(f"Error fetching data: {response.reason}")
        return None

def lambda_handler(event, context):
    vulnerabilities = get_vulnerabilities()

    if vulnerabilities:
        sqs = boto3.client('sqs')
        unfiltered_queue_url = os.getenv('UNFILTERED_QUEUE_URL')

        for vuln in vulnerabilities.get('vulnerabilities', []):
            # Send each vulnerability as a separate message
            response = sqs.send_message(
                QueueUrl=unfiltered_queue_url,
                MessageBody=json.dumps(vuln)
            )
            print(f"Message sent with ID: {response['MessageId']}")
    else:
        print("No vulnerabilities found or error in request")

    return {
        'statusCode': 200,
        'body': json.dumps('Vulnerabilities processing completed.')
    }
