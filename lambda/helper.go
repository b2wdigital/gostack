package lambda

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/b2wdigital/gostack/cloudevents"
)

type Helper struct {
	handler *Handler
}

func NewHelper(handler *cloudevents.HandlerWrapper, options *Options) *Helper {

	h := NewHandler(handler, options)

	return &Helper{
		handler: h,
	}
}

func NewDefaultHelper(handler *cloudevents.HandlerWrapper) *Helper {

	opt, err := DefaultOptions()
	if err != nil {
		log.Fatal(err.Error())
	}

	return NewHelper(handler, opt)
}

func (h *Helper) Start() {
	lambda.Start(h.handler.Handle)
}
