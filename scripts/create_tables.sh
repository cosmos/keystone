#!/bin/sh

# Create dynamodb tables on LOCAL DynamoDB instance and first delete
# them if they already exist

aws dynamodb delete-table \
    --table-name Resolved \
    --endpoint-url http://localhost:8000

aws dynamodb delete-table \
    --table-name Promised \
    --endpoint-url http://localhost:8000

aws dynamodb create-table \
    --table-name Resolved \
    --attribute-definitions \
    AttributeName=Id,AttributeType=S \
    --key-schema \
    AttributeName=Id,KeyType=HASH \
    --billing-mode PAY_PER_REQUEST \
    --endpoint-url http://localhost:8000

aws dynamodb create-table \
    --table-name Promised \
    --attribute-definitions \
        AttributeName=Id,AttributeType=S \
    --key-schema \
        AttributeName=Id,KeyType=HASH \
    --billing-mode PAY_PER_REQUEST \
    --endpoint-url http://localhost:8000

