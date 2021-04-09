package lambda

import (
	"context"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/b2wdigital/goignite/v2/core/errors"
	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/b2wdigital/gostack/cloudevents"
)

type Handler struct {
	handler *cloudevents.HandlerWrapper
	options *Options
}

func NewHandler(handler *cloudevents.HandlerWrapper, options *Options) *Handler {
	return &Handler{handler: handler, options: options}
}

// Handler handles a event
func (h *Handler) Handle(ctx context.Context, event Event) error {

	logger := log.FromContext(ctx).WithTypeOf(*h)

	lc, ok := lambdacontext.FromContext(ctx)
	if !ok {
		return errors.Internalf("lambda context not exists")
	}

	logger = logger.WithField("awsrequestid", lc.AwsRequestID)
	ctx = logger.ToContext(ctx)

	if h.options.Skip {
		logger.Info("skipping event")
		return nil
	}

	inouts, err := h.getInOuts(ctx, event)
	if err != nil {
		return err
	}

	if len(inouts) > 0 {

		err := h.handler.Process(ctx, inouts)
		if err != nil {
			logger.Error(errors.ErrorStack(err))
			return err
		}

		logger.Debug("all events called")

		for _, inout := range inouts {
			if inout.Err != nil {
				err := errors.Wrap(inout.Err, errors.New("closing with errors lambda handle"))
				logger.Error(errors.ErrorStack(err))
				return err
			}
		}

	}

	logger.Info("closing lambda handle")

	return nil
}

func (h *Handler) getInOuts(ctx context.Context, event Event) ([]*cloudevents.InOut, error) {

	logger := log.FromContext(ctx)

	var inouts []*cloudevents.InOut

	if len(event.Records) > 0 {

		if event.Records[0].EventSource == "aws:kinesis" {
			inouts = fromKinesis(ctx, event)
		} else if event.Records[0].EventSource == "aws:sqs" {
			inouts = fromSQS(ctx, event)
		} else if event.Records[0].EventSource == "aws:sns" {
			inouts = fromSNS(ctx, event)
		} else {
			return nil, errors.NotImplementedf("the trigger received has not yet been implemented")
		}

	} else {

		if event.Source == "aws.events" {
			inouts = fromCloudWatch(ctx, event)
		} else {
			logger.Warnf("ignoring trigger")
			return nil, nil
		}

	}

	return inouts, nil
}
