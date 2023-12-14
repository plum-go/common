package sms

import (
	"fmt"
	"testing"

	"github.com/plum-go/common/sms"
)

func TestVerify(t *testing.T) {
	phone := "+8618664391697"
	code := "123456"
	message, error := sms.Verify(phone, code)

	if fmt.Sprintf("Hi, %v. Code %v validate successful!", phone, code) != message || error != nil {
		t.Fatalf(`sms.Verify(phone, code) not contains %v and %v, want true`, phone, code)
	}
}
