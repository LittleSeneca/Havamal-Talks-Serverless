# Havamal-Talks-Serverless

## Overview

This project is designed to gather new content from the NIST National Vulnerability Database (NVD) using the NVD CVE API. It automatically deploys AWS Lambda functions to process various severity levels of vulnerabilities and then pushes them into an Amazon Simple Queue Service (SQS) queue. This setup helps in monitoring and analyzing security vulnerabilities as reported by NIST.

## Getting a NIST API Key

To access the NIST NVD CVE API, you need an API key. Follow these steps to request a new API key:

1. Visit the [NIST API Key Request page](https://nvd.nist.gov/developers/request-an-api-key).
2. Fill in the required details, including your name and email address.
3. After submission, NIST will send the API key to your email. This key is required to authenticate your requests to the NVD CVE API.

## Prerequisites

- An AWS account with permissions for Lambda, S3, CloudFormation, SNS, and SQS.
- A GitHub account.
- Familiarity with AWS services and GitHub Actions.

## Forking the Repository

1. Go to the GitHub repository URL.
2. Click the 'Fork' button at the top right to fork it into your account.

## Setting Up AWS Credentials

1. In the AWS Console, navigate to IAM.
2. Create a new IAM user with programmatic access.
3. Attach policies for managing Lambda, S3, CloudFormation, SNS, and SQS.
4. Note the Access Key ID and Secret Access Key.

## Configuring GitHub Actions Secrets

1. In your forked GitHub repository, go to 'Settings' > 'Secrets'.
2. Click 'New repository secret' and add the following:
   - `AWS_ACCESS_KEY_ID`: Your AWS IAM Access Key ID.
   - `AWS_SECRET_ACCESS_KEY`: Your AWS IAM Secret Access Key.
   - `NVD_API_KEY`: Your NIST NVD API Key.
   - `S3_BUCKET_NAME`: The S3 bucket name for Lambda deployments.

## Deployment Process

1. **S3 Bucket Setup**: Deployed via `s3-template.yml`.
2. **Lambda Functions**: Deployed via `lambda-fetch-template.yml`, creating four functions for different severity levels (critical, high, medium, low).
3. **GitHub Actions**: The workflow in `.github/workflows/lambda_deploy.yml` automates the deployment process.

## Triggering Deployments

- Deployments are triggered by pushes or pull requests to the `main` branch.

## Monitoring Deployments

- View progress in the 'Actions' tab of your GitHub repository.

## Notes

- Ensure the S3 bucket name is unique.
- Review AWS IAM policies for security.

## Troubleshooting

- Check GitHub Actions logs for errors.
- Use AWS Console or CLI for further diagnostics.

This README provides an outline for setting up and deploying the project. For detailed instructions, consult the repository documentation and specific AWS/GitHub Actions resources.