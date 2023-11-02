package dingtalk

import (
	"github.com/pkg6/go-notify"
	"net/url"
	"reflect"
	"testing"
)

func TestClient_I(t1 *testing.T) {
	type fields struct {
		AccessToken string
		Secret      string
	}
	tests := []struct {
		name   string
		fields fields
		want   notify.IClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Client{
				AccessToken: tt.fields.AccessToken,
				Secret:      tt.fields.Secret,
			}
			if got := t.I(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("I() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Name(t1 *testing.T) {
	type fields struct {
		AccessToken string
		Secret      string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Client{
				AccessToken: tt.fields.AccessToken,
				Secret:      tt.fields.Secret,
			}
			if got := t.Name(); got != tt.want {
				t1.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Send(t1 *testing.T) {
	type fields struct {
		AccessToken string
		Secret      string
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
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Client{
				AccessToken: tt.fields.AccessToken,
				Secret:      tt.fields.Secret,
			}
			if gotResult := t.Send(tt.args.message); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t1.Errorf("Send() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestClient_url(t1 *testing.T) {
	type fields struct {
		AccessToken string
		Secret      string
	}
	tests := []struct {
		name   string
		fields fields
		want   *url.URL
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Client{
				AccessToken: tt.fields.AccessToken,
				Secret:      tt.fields.Secret,
			}
			if got := t.url(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("url() = %v, want %v", got, tt.want)
			}
		})
	}
}
