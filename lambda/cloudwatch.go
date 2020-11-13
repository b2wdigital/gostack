package lambda

import (
	"context"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/b2wdigital/gostack/cloudevents"
	v2 "github.com/cloudevents/sdk-go/v2"
)

func fromCloudWatch(parentCtx context.Context, event Event) []*cloudevents.InOut {

	logger := log.FromContext(parentCtx)
	logger.Info("receiving cloudwatch event")

	lc, _ := lambdacontext.FromContext(parentCtx)

	var inouts []*cloudevents.InOut

	in := v2.NewEvent()

	in.SetType(event.DetailType)
	in.SetID(event.ID)
	in.SetSource(event.Source)

	in.SetExtension("awsRequestID", lc.AwsRequestID)
	in.SetExtension("invokedFunctionArn", lc.InvokedFunctionArn)

	inouts = append(inouts, &cloudevents.InOut{
		In:  &in,
		Err: nil,
	})

	return inouts
}
