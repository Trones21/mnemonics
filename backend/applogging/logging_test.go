package applogging_test

import (
	"fmt"
	"mnemonics/applogging"
	"net/http"
	"strings"
	"testing"
)

//Code paths to test
//Should Pass - g?user=1234
//Should Fail - g?user1234 (Missing =)
//Should Fail? - g    (no RawQuery)
//ShouldFail - g?usser=1234  (Not a valid entityType)

func TestParseLoggingInfo(t *testing.T) {
	//Arrange
	body := strings.NewReader(`
	{
		"EntityType":"mnemonic",
		"EntityID":"NotaValidID",
		"ActorID":"21334sfs",
		"IpAddress":"192.168.1.1",
		"IsRegistreredActor":false 
	}
	`)
	req, err := http.NewRequest("POST", "http://localhost:5000/upvotelog", body)
	if err != nil {
		t.Errorf("Error Creating Request")
	}
	//Act
	logInfo, err := applogging.ParseLoggingInfo(req)
	if err != nil {
		//Returning an error is actually the desired behavior for improperly formatted data (url params or header vals)
		t.Errorf(err.Error())
	}
	//Assert
	if logInfo.EntityID != "1234" &&
		logInfo.EntityType != "mnemonic" &&
		logInfo.ActorID != "21334s" &&
		logInfo.IpAddress != "192.168.1.1" &&
		logInfo.IsRegisteredActor != false {
		fmt.Println("Success")
		return
	} else {
		fmt.Println("Failed")
		//t.FailNow()
	}
}

//Why does t.Failed() output "ok"
func TestSandbox(t *testing.T) {
	str := "124"
	if "123" != str {
		//t.FailNow()
		t.Failed()
	}
}

// func TestGetLoggingInfoTableDriven(t *testing.T) {
// 	//Arrange
// 	scenarios := []struct {
// 		req           http.Request
// 		expectLogInfo applogging.LoggingInfo
// 		expectErr     string
// 	}{
// 		{req: http.NewRequest("POST", "http://localhost:5000/upvotelog?user=1234", nil), expectLogInfo: applogging.LoggingInfo{EntityType: "user", EntityID: "1234"}},
// 		{req: http.NewRequest("POST", "http://localhost:5000/upvoteloguser=1234", nil), expectErr: "Func Failure"},
// 	}
// 	//Act
// 	logInfo, err := applogging.GetLoggingInfo(req)
// 	//Assert

// 	if logInfo.EntityID != "1234" ||
// 		logInfo.EntityType != "user" {
// 		t.Errorf("Test fail")
// 	}
// }
