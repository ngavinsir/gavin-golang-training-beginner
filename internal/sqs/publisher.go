package sqs

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Publisher struct {
	SQS      *sqs.SQS
	QueueUrl *string
}

func NewPublisher(s *session.Session, queue string) (*Publisher, error) {
	var p Publisher
	sqsClient := sqs.New(s)

	res, err := sqsClient.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &queue,
	})

	if err != nil {
		return &p, err
	}

	p.SQS = sqsClient
	p.QueueUrl = res.QueueUrl

	return &p, nil
}

func (p Publisher) Publish(msg interface{}) (err error) {
	messageBody, _ := json.Marshal(msg)

	_, err = p.SQS.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(messageBody)),
		QueueUrl:    p.QueueUrl,
	})

	return
}