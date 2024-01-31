import json
import boto3
import os

def format_message(cve_data):
    try:
            cve_id = cve_data['cve']['id']
            publish_date = cve_data['cve']['published']
            description = cve_data['cve']['descriptions'][0]['value']
            base_score = cve_data['cve']['metrics']['cvssMetricV31'][0]['cvssData']['baseScore']
            base_severity = cve_data['cve']['metrics']['cvssMetricV31'][0]['cvssData']['baseSeverity']
            attack_vector = cve_data['cve']['metrics']['cvssMetricV31'][0]['cvssData']['attackVector']
            attack_complexity = cve_data['cve']['metrics']['cvssMetricV31'][0]['cvssData']['attackComplexity']
            privileges_required = cve_data['cve']['metrics']['cvssMetricV31'][0]['cvssData']['privilegesRequired']
            user_interaction = cve_data['cve']['metrics']['cvssMetricV31'][0]['cvssData']['userInteraction']
            scope = cve_data['cve']['metrics']['cvssMetricV31'][0]['cvssData']['scope']
            confidentiality_impact = cve_data['cve']['metrics']['cvssMetricV31'][0]['cvssData']['confidentialityImpact']
            integrity_impact = cve_data['cve']['metrics']['cvssMetricV31'][0]['cvssData']['integrityImpact']
            availability_impact = cve_data['cve']['metrics']['cvssMetricV31'][0]['cvssData']['availabilityImpact']

            formatted_message = (
                f"CVEID: {cve_id}\n"
                f"Publish Date: {publish_date}\n"
                f"Description: {description}\n"
                f"BaseScore: {base_score}\n"
                f"BaseSeverity: {base_severity}\n"
                f"Attack Vector: {attack_vector} | Attack Complexity: {attack_complexity} | "
                f"Privileges Required: {privileges_required} / User Interaction: {user_interaction} | Scope: {scope}\n"
                f"Confidentiality Impact: {confidentiality_impact} | Integrity Impact: {integrity_impact} | Availability Impact: {availability_impact}\n"
            )
            return formatted_message

    except KeyError as e:
        print(f"Key error: {e}")
        return "Error in formatting message."

def lambda_handler(event, context):
    # Initialize SQS client
    sqs = boto3.client('sqs')

    # URLs for the original and destination queues
    filtered_queue_url = os.getenv('FILTERED_QUEUE_URL')
    formated_queue_url = os.getenv('FORMATTED_QUEUE_URL')

    # Process each record from the SQS message
    for record in event['Records']:
        message = json.loads(record['body'])
        formatted_output = format_message(message)

        # Send the formatted message to the destination queue
        sqs.send_message(QueueUrl=formated_queue_url, MessageBody=formatted_output)

        # Delete the message from the original queue after processing
        sqs.delete_message(QueueUrl=filtered_queue_url, ReceiptHandle=record['receiptHandle'])

    return {
        'statusCode': 200,
        'body': json.dumps('Message formatting and forwarding completed.')
    }
