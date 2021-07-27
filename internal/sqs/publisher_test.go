package sqs_test

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	sqsSDK "github.com/aws/aws-sdk-go/service/sqs"
	"github.com/ngavinsir/golangtraining"
	"github.com/ngavinsir/golangtraining/internal/sqs"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestPublisherPublish(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip the Test for SQS Publisher")
	}

	region := "ap-southeast-1"
	endpoint := "http://localhost:4566/000000000000/payments"
	queue := "payments"

	sess, err := session.NewSession(&aws.Config{
		Region:   &region,
		Endpoint: &endpoint,
	})
	if err != nil {
		panic(err)
	}

	mockPayment := golangtraining.Payment{
		ID: "1",
		PaymentCode: "code",
		TransactionID: "1",
		Name: "name",
		Amount: "12345",
	}

	type resType struct {
		Res golangtraining.Payment
		Err error
	}

	tests := []struct {
		desc        string
		publisher   *sqs.Publisher
		msg         interface{}
		expectedRes resType
	}{
		{
			desc: "publish - success",
			publisher: func() *sqs.Publisher{
				p, err := sqs.NewPublisher(sess, queue)
				if err != nil {
					panic(err)
				}

				return p
			}(),
			msg: mockPayment,
			expectedRes: resType{
				Res: mockPayment,
				Err: nil,
			},
		},
	}
	for _, tC := range tests {
		t.Run(tC.desc, func(t *testing.T) {
			_, err = tC.publisher.SQS.PurgeQueue(&sqsSDK.PurgeQueueInput{
				QueueUrl: tC.publisher.QueueUrl,
			})
			if err != nil {
				t.Error(err)
			}

			err := tC.publisher.Publish(tC.msg)

			if tC.expectedRes.Err != nil {
				require.Equal(t, tC.expectedRes.Err.Error(), errors.Cause(err).Error())
			} else {
				timeout := int64(5)
				msgResult, err := tC.publisher.SQS.ReceiveMessage(&sqsSDK.ReceiveMessageInput{
					AttributeNames: []*string{
						aws.String(sqsSDK.MessageSystemAttributeNameSentTimestamp),
					},
					QueueUrl:            tC.publisher.QueueUrl,
					MaxNumberOfMessages: aws.Int64(10),
					VisibilityTimeout:   &timeout,
				})
				if err != nil {
					t.Error(err)
				}

				var msgPayment golangtraining.Payment
				err = json.Unmarshal([]byte(*msgResult.Messages[0].Body), &msgPayment)
				if err != nil {
					t.Error(err)
				}

				require.Equal(t, tC.expectedRes.Res, msgPayment)
			}

			_, err = tC.publisher.SQS.PurgeQueue(&sqsSDK.PurgeQueueInput{
				QueueUrl: tC.publisher.QueueUrl,
			})
			if err != nil {
				t.Error(err)
			}
		})
	}
}