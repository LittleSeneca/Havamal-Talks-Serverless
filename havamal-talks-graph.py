import boto3

dynamodb = boto3.resource('dynamodb', region_name='us-east-1')
table = dynamodb.Table('plaintextdisclosures')
response = table.scan()
data = response['Items']
count = 0
for x in data:
    if x['baseSeverity'] == 'CRITICAL' or x['baseSeverity'] == 'HIGH':
        count = count + 1 
        print(x['cveid'])

print("There are " + str(count) + " critical or high severity vulnerabilities in the database.")