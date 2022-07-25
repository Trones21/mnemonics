package applogging_test

import (
	"mnemonics/applogging"
	"net/http"
	"testing"
)

//Code paths to test
//Should Pass - g?user=1234
//Should Fail - g?user1234 (Missing =)
//Should Fail? - g    (no RawQuery)
//ShouldFail - g?usser=1234  (Not a valid entityType)

func TestGetLoggingInfo(t *testing.T) {
	//Arrange
	req, err := http.NewRequest("POST", "http://localhost:5000/upvotelog?user1234", nil)
	if err != nil {
		t.Errorf("Error Creating Request")
	}
	//Act
	logInfo, err := applogging.GetLoggingInfo(req)
	if err != nil {
		//Returning an error is actually the desired behavior for improperly formatted data (url params or header vals)
		t.Errorf(err.Error())
	}
	//Assert
	if logInfo.EntityID != "1234" ||
		logInfo.EntityType != "user" {
		t.Errorf(err.Error())
	}
}

func TestSandbox(t *testing.T) {
	str := "123"
	if "123" != str {
		t.FailNow()
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
