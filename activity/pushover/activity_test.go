package pushover

import (
	"io/ioutil"
	"testing"
	"log"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
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

	act := &PushoverActivity{getActivityMetadata()}
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput(ivUserKey, "[YOUR_PUSHOVER_USER_KEY]")
	tc.SetInput(ivAuthToken, "[YOUR_PUSHOVER_APP_TOKEN")
	tc.SetInput(ivMessage, "Go Flogo")

	act.Eval(tc)
	log.Printf("TestEval successfull. Output [%d]", tc.GetOutput("status"))
	//check result attr
}
