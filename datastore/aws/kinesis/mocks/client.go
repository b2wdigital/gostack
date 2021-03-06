// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	kinesis "github.com/aws/aws-sdk-go-v2/service/kinesis"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// BulkPublish provides a mock function with given fields: ctx, messages
func (_m *Client) BulkPublish(ctx context.Context, messages []kinesis.PutRecordInput) error {
	ret := _m.Called(ctx, messages)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []kinesis.PutRecordInput) error); ok {
		r0 = rf(ctx, messages)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
