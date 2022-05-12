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

