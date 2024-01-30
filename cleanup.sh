#!/bin/bash

# Specify your S3 bucket name and CloudFormation stack prefix
S3_BUCKET_NAME="havamal-talks-serverless"
STACK_PREFIX="havamal-talks-serverless"

# Step 1: Remove all files from the S3 bucket
aws s3 rm s3://${S3_BUCKET_NAME}/ --recursive

# Step 2: List CloudFormation stacks starting with the specified prefix
stacks=$(aws cloudformation list-stacks --query "StackSummaries[?starts_with(StackName, '$STACK_PREFIX')].[StackName]" --output text)

# Step 3: Delete CloudFormation stacks
for stack in $stacks; do
    echo "Deleting stack: $stack"
    aws cloudformation delete-stack --stack-name $stack
done
