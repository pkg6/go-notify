package alimail

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/pkg6/go-notify"
	"reflect"
	"testing"
)

func TestClient_I(t *testing.T) {
	type fields struct {
		AccessKeyId     string
		AccessKeySecret string
		aliClient       *openapi.Client
	}
	tests := []struct {
		name   string
		fields fields
		want   notify.IClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				AccessKeyId:     tt.fields.AccessKeyId,
				AccessKeySecret: tt.fields.AccessKeySecret,
				aliClient:       tt.fields.aliClient,
			}
			if got := c.I(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("I() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Name(t *testing.T) {
	type fields struct {
		AccessKeyId     string
		AccessKeySecret string
		aliClient       *openapi.Client
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				AccessKeyId:     tt.fields.AccessKeyId,
				AccessKeySecret: tt.fields.AccessKeySecret,
				aliClient:       tt.fields.aliClient,
			}
			if got := c.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Send(t *testing.T) {
	type fields struct {
		AccessKeyId     string
		AccessKeySecret string
		aliClient       *openapi.Client
	}
	type args struct {
		message notify.IMessage
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   notify.IResult
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				AccessKeyId:     tt.fields.AccessKeyId,
				AccessKeySecret: tt.fields.AccessKeySecret,
				aliClient:       tt.fields.aliClient,
			}
			if got := c.Send(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Send() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_commonParams(t *testing.T) {
	type fields struct {
		AccessKeyId     string
		AccessKeySecret string
		aliClient       *openapi.Client
	}
	tests := []struct {
		name   string
		fields fields
		want   *openapi.Params
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				AccessKeyId:     tt.fields.AccessKeyId,
				AccessKeySecret: tt.fields.AccessKeySecret,
				aliClient:       tt.fields.aliClient,
			}
			if got := c.commonParams(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commonParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
