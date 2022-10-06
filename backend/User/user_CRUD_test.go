package user_crud_test

import (
	"net/http"
	"testing"
)

//create user
//Read user (see table)
//Update User (see table)
//Delete - Not sure of business logic, should we ever delete a user?
//----//(Definitely don't want to cascade delete UGC (user gen content), rather transfer ownership to anonymous)

//var url //figure out how to apply this to all tests later

func Test_createUser(t *Testing.T) {
	body := 
	req, err := http.NewRequest("POST", "http://localhost:5000/user", body)
	if err != nil {
		t.Errorf("Error Creating Request")
	}
}

//Table test to simulate many scenarios
func Test_getUser(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}
