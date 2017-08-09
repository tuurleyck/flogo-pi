package pushover

import (
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
)

func TestRegistered(t *testing.T) {
	act := activity.Get("pushover")

	if act == nil {
		t.Error("Activity Not Registered")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	md := activity.NewMetadata(jsonMetadata)
	act := &PushoverActivity{metadata: md}

	tc := test.NewTestActivityContext(md)
	//setup attrs
	tc.SetInput(ivUserKey, "[YOUR_PUSHOVER_USER_KEY]")
	tc.SetInput(ivAuthToken, "[YOUR_PUSHOVER_APP_TOKEN")
	tc.SetInput(ivMessage, "Go Flogo")

	act.Eval(tc)

	//check result attr
}
