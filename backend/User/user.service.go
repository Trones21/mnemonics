package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mnemonics/cors"
	"net/http"
	"strings"
)

func SetupRoutes(apiBasePath string) {
	handleusers := http.HandlerFunc(userListHandler)
	handleuser := http.HandlerFunc(userHandler)
	http.Handle(fmt.Sprintf("%s/userList", apiBasePath), handleusers)
	http.Handle(fmt.Sprintf("%s/user/", apiBasePath), cors.Middleware(handleuser))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "user/")
	//ToDo: Fix this so that the authorization handle security, right now this is WIDE OPEN, anyone can update a user
	loggedInuserID := urlPathSegments[len(urlPathSegments)-1]

	switch r.Method {
	case http.MethodGet:
		fmt.Println("Get Item")
		user, err := getuser(loggedInuserID)
		if user == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not found"))
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			//Follow this guide to send specific info in debug, while responding with
			//vague info for prod (to prevent consumers from understandinginternals)
			//http://speakmy.name/014/07/29/http-request-debugging-in-go/
			w.Write([]byte("err"))
		}

		userJSON, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error parsing JSON in service file"))
			return
		}
		w.Header().Set("Cntent-Type", "application/json")
		w.Write(userJSON)

	case http.MethodPost:
		var newUser User
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
		w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &newUser)
		if err != nil {
		w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write([]byte("Need to add logic for adding user"))
		return 
	case http.MethodPut:
		var updateduser User
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
		w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(bodyBytes, &updateduser)
		if err != nil {
		w.WriteHeader(http.StatusBadRequest)
			return
		}

		if updateduser.UserID != loggedInuserID {
		w.WriteHeader(http.StatusBadRequest)
			return
		}

		updateUser(&updateduser)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Succesfully updated user"))
		return

	case http.MethodDelete:
		//remoeuser(userID)

	case http.MethodOptions:
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
}

}

func userListHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userList, err := getuserList()
		if err == nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("%s", err)))
			return
		}
		usersJSON, err := json.Marshal(userList)
		if err == nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error parsing JSON"))
			return
		}
		w.Header().Set("ContentType", "application/json")
		w.Write(usersJSON)
		return
	case http.MethodOptions:
		return
	}
}
