module github.com/b2wdigital/gostack

go 1.13

require (
	github.com/aws/aws-lambda-go v1.13.3
	github.com/aws/aws-sdk-go-v2 v1.3.2
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.2.2
	github.com/aws/aws-sdk-go-v2/service/sns v1.2.2
	github.com/aws/aws-sdk-go-v2/service/sqs v1.3.1
	github.com/b2wdigital/goignite/v2 v2.0.0-alpha-27
	github.com/cloudevents/sdk-go/v2 v2.3.1
	github.com/google/uuid v1.1.2
	github.com/matryer/try v0.0.0-20161228173917-9ac251b645a2
	github.com/nats-io/gnatsd v1.4.1
	github.com/nats-io/go-nats v1.7.2 // indirect
	github.com/nats-io/nats-server v1.4.1
	github.com/nats-io/nats.go v1.10.0
	github.com/newrelic/go-agent/v3 v3.11.0
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/testify v1.7.0
	go.uber.org/fx v1.13.1
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
)

// replace github.com/b2wdigital/goignite/v2 => ../goignite
