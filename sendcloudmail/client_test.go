package sendcloudmail

import (
	"github.com/pkg6/go-notify"
	"reflect"
	"testing"
)

func TestClient_I(t *testing.T) {
	type fields struct {
		ApiUser string
		ApiKey  string
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
				ApiUser: tt.fields.ApiUser,
				ApiKey:  tt.fields.ApiKey,
			}
			if got := c.I(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("I() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Name(t *testing.T) {
	type fields struct {
		ApiUser string
		ApiKey  string
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
				ApiUser: tt.fields.ApiUser,
				ApiKey:  tt.fields.ApiKey,
			}
			if got := c.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Send(t *testing.T) {
	type fields struct {
		ApiUser string
		ApiKey  string
	}
	type args struct {
		message notify.IMessage
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult notify.IResult
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				ApiUser: tt.fields.ApiUser,
				ApiKey:  tt.fields.ApiKey,
			}
			if gotResult := c.Send(tt.args.message); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Send() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
