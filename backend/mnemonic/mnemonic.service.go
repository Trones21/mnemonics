package mnemonic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mnemonics/cors"
	"net/http"
	"strings"
)

func SetupRoutes(apiBasePath string) {
	handleMnemonics := http.HandlerFunc(MnemonicListHandler)
	handleMnemonic := http.HandlerFunc(MnemonicItemHandler)
	http.Handle(fmt.Sprintf("%s/mnemonicsList", apiBasePath), handleMnemonics)
	http.Handle(fmt.Sprintf("%s/mnemonic/", apiBasePath), cors.Middleware(handleMnemonic))
}

func MnemonicItemHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "mnemonic/")
	mnemonicID := urlPathSegments[len(urlPathSegments)-1]

	switch r.Method {
	case http.MethodGet:
		fmt.Println("Get Item")
		mnemonic, err := GetMnemonic(mnemonicID)
		if mnemonic == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not found"))
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			//Follow this guide to send specific info in debug, while responding with
			//vague info for prod (to prevent consumers from understanding internals)
			//http://speakmy.name/2014/07/29/http-request-debugging-in-go/
			w.Write([]byte("err"))
		}

		mnemonicJSON, err := json.Marshal(mnemonic)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error parsing JSON in service file"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(mnemonicJSON)
	case http.MethodPut:
		var updatedMnemonic Mnemonic
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(bodyBytes, &updatedMnemonic)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if updatedMnemonic.MnemonicID != mnemonicID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		UpdateMnenomic(updatedMnemonic)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Successfully updated mnemonic"))
		return
	case http.MethodPost:
		var newMnemonic Mnemonic
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(bodyBytes, &newMnemonic)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		rowsCreated, err := AddMnenomic(newMnemonic)
		if rowsCreated != 1 {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Successfully created mnemonic"))
		return
	case http.MethodDelete:
		DeleteMnemonic(mnemonicID)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}

func MnemonicListHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		mnemonicsList, err := GetMnemonicList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("%s", err)))
			return
		}
		mnemonicsJSON, err := json.Marshal(mnemonicsList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error parsing JSON"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(mnemonicsJSON)
		return
	case http.MethodOptions:
		return
	}
}

func validateMnemonic() {

}
