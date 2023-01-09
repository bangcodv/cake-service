// Package handler
package handler

import (
	"context"

	"github.com/bangcodv/cake-service/internal/appctx"
	"github.com/bangcodv/cake-service/internal/consts"
	uContract "github.com/bangcodv/cake-service/internal/ucase/contract"
	"github.com/bangcodv/cake-service/pkg/awssqs"
)

// SQSConsumerHandler sqs consumer message processor handler
func SQSConsumerHandler(msgHandler uContract.MessageProcessor) awssqs.MessageProcessorFunc {
	return func(decoder *awssqs.MessageDecoder) error {
		return msgHandler.Serve(context.Background(), &appctx.ConsumerData{
			Body:        []byte(*decoder.Body),
			Key:         []byte(*decoder.MessageId),
			ServiceType: consts.ServiceTypeConsumer,
		})
	}
}
