# ddb-ttl-lambda
Lambda function, which move dynamodb records to firehose in json via DynamoDB Streams.

# Overview
Export DynamoDB records in json to reduce storage.
Setting TTL attribute is effective way to delete records but of cource you can delete records by `DeleteItem`

# Requirement
Enabled DynamoDB Streams(Old images required)
Enabled DynamoDB TTL
Firehose
Lambda
S3

## Workload
- records are deleted by ttl
- changed items appear in Streams
- Lambda which connected with Streams detects `REMOVE` event and sends items in json to firehose
- Firehose delivers them to s3
