package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"
)

var (
	svc *firehose.Firehose
)

func init() {
	svc = firehose.New(session.New())
}

func main() {
	if os.Getenv("FIREHOSE_STREAM_NAME") == "" {
		log.Fatalln("FIREHOSE_STREAM_NAME missing")
	}
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, e events.DynamoDBEvent) error {
	buf := make([][]byte, 0, len(e.Records))
	for _, v := range e.Records {
		if v.EventName != "REMOVE" {
			continue
		}
		blob, err := json.Marshal(v.Change.OldImage)
		if err != nil {
			log.Println(err)
			continue
		}
		buf = append(buf, blob)
	}
	return putBatch(buf)
}

func putBatch(bb [][]byte) error {
	if len(bb) == 0 {
		return nil
	}
	records := make([]*firehose.Record, 0, len(bb))
	for _, v := range bb {
		record := &firehose.Record{
			Data: append(v, 10),
		}
		records = append(records, record)
	}
	input := &firehose.PutRecordBatchInput{
		DeliveryStreamName: aws.String(os.Getenv("FIREHOSE_STREAM_NAME")),
		Records:            records,
	}
	_, err := svc.PutRecordBatch(input)
	return err
}
