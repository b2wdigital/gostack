package lambda

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/b2wdigital/goignite/v2/core/errors"
	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/b2wdigital/gostack/cloudevents"
	v2 "github.com/cloudevents/sdk-go/v2"
	"golang.org/x/sync/errgroup"
)

func fromSNS(parentCtx context.Context, event Event) []*cloudevents.InOut {

	logger := log.FromContext(parentCtx)
	logger.Info("receiving SNS event")

	lc, _ := lambdacontext.FromContext(parentCtx)

	var inouts []*cloudevents.InOut

	g, gctx := errgroup.WithContext(parentCtx)

	for _, record := range event.Records {

		record := record

		g.Go(func() error {

			j, _ := json.Marshal(record)
			logger.Debug(string(j))

			var err error

			in := v2.NewEvent()

			if err = json.Unmarshal([]byte(record.SNS.Message), &in); err != nil {

				var data interface{}

				if err = json.Unmarshal([]byte(record.SNS.Message), &data); err != nil {
					err = errors.NewNotValid(err, "could not decode SNS record")
				} else {
					if err = in.SetData(v2.ApplicationJSON, data); err != nil {
						err = errors.NewNotValid(err, "could not set data in event")
					}
				}

			}

			in.SetType(record.SNS.Type)

			if in.ID() == "" {
				in.SetID(record.SNS.MessageID)
			}

			in.SetSource(record.EventSource)

			in.SetExtension("awsRequestID", lc.AwsRequestID)
			in.SetExtension("invokedFunctionArn", lc.InvokedFunctionArn)

			inouts = append(inouts, &cloudevents.InOut{
				In:  &in,
				Err: err,
			})

			return nil
		})

	}

	if err := g.Wait(); err == nil {
		logger.Debug("all events converted")
	}

	gctx.Done()

	return inouts
}
