import json
import boto3
import os

def lambda_handler(event, context):
    # Keywords to check in the message
    keywords = ["HIGH", "CRITICAL", "CHANGED", "macOS", "mac", "windows", "linux", "ubuntu", "rhel", "redhat", "red hat"]

    # Initialize SQS client
    sqs = boto3.client('sqs')
    
    # URLs for the original and filtered queues
    unfiltered_queue_url = os.getenv('UNFILTERED_QUEUE_URL')
    filtered_queue_url = os.getenv('FILTERED_QUEUE_URL')

    # Process each record from the SQS message
    for record in event['Records']:
        message = json.loads(record['body'])
        message_str = json.dumps(message) if not isinstance(message, str) else message

        # Check for important keywords
        if any(keyword in message_str for keyword in keywords):
            # Send to filtered queue
            sqs.send_message(QueueUrl=filtered_queue_url, MessageBody=message_str)
        else:
            # Delete the message from the original queue if no keyword is found
            sqs.delete_message(QueueUrl=unfiltered_queue_url, ReceiptHandle=record['receiptHandle'])

    return {
        'statusCode': 200,
        'body': json.dumps('Message processing completed.')
    }
