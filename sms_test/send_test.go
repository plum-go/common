package sms

import (
	"strings"
	"testing"

	"github.com/plum-go/common/sms"
)

func TestSend(t *testing.T) {
	name := "+8618664391697"
	message := sms.Send(name)
	isContains := strings.Contains(message, name)

	if !isContains {
		t.Fatalf(`Send(name) not contains %v, want true`, message)
	}
}

func TestSendMultiple(t *testing.T) {
	name := "+8618664391697"
	nameMaps, error := sms.SendMultiple([]string{name})
	isContains := strings.Contains(nameMaps[name], name)

	if !isContains || error != nil {
		t.Fatalf(`Send([name]) not contains %v, want true`, name)
	}

}
