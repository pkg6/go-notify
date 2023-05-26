package notify_test

import (
	"github.com/pkg6/go-notify"
	"github.com/pkg6/go-notify/dingtalk"
	"github.com/pkg6/go-notify/wework"
	"testing"
)

func TestNotify_Send(t *testing.T) {
	weworkClient := &wework.Client{Key: "693axxx6-7aoc-4bc4-97a0-0ec2sifa5aaa"}
	dingTalkClient := &dingtalk.Client{
		AccessToken: "566cc69da782ec******",
	}
	weworkText := wework.NewTextMessage("测试")
	dingTalkText := dingtalk.NewTextMessage("测试")
	tests := []struct {
		name       string
		clientName string
		client     notify.IClient
		message    notify.IMessage
	}{
		{

			name:       "wework",
			clientName: notify.WeworkName,
			client:     weworkClient,
			message:    weworkText,
		},
		{

			name:       "dingtalk",
			clientName: notify.DingTalkName,
			client:     dingTalkClient,
			message:    dingTalkText,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := notify.New(tt.client)
			got := n.Send(tt.message)
			for s, result := range got {
				if s == tt.clientName {
					if result.Status() != notify.RESULT_STATUS_ERROR {
						t.Fatalf(tt.name + " send fail")
					}
				}
			}
		})
	}
}
