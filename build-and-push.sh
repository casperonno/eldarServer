#!/bin/bash

# Install the AWS CLI
apt-get update && \
apt-get install -y awscli

# Authenticate with ECR
$(aws ecr get-login --no-include-email --region ${AWS_REGION})

# Build the Docker image
docker build -t ${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${ECR_REPO} .

# Push the Docker image to ECR
docker push ${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${ECR_REPO}