#!/bin/bash

# Install AWS CLI
pip install awscli --upgrade --user

# Add AWS CLI executable to the PATH
export PATH=$PATH:$HOME/.local/bin

# Configure AWS CLI with your secrets
aws configure set aws_access_key_id ${{ secrets.AWS_ACCESS_KEY_ID }}
aws configure set aws_secret_access_key ${{ secrets.AWS_SECRET_ACCESS_KEY }}
aws configure set default.region us-east-1

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
