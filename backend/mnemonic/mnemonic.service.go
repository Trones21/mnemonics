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
	http.Handle(fmt.Sprintf("%s/mnemonics/", apiBasePath), cors.Middleware(handleMnemonic))
}

func MnemonicItemHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "mnemonics/")
	mnemonicID := urlPathSegments[len(urlPathSegments)-1]

	switch r.Method {
	case http.MethodGet:
		fmt.Println("Get Item")
		mnemonic, err := getMnemonic(mnemonicID)
		if mnemonic == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not found"))
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			//Follow this guide to send specific info in debug, while responding with
			//vague info for prod (to prevent consumers from understaning internal)
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

	case http.MethodPost:
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

		addOrUpdateMnenomic(&updatedMnemonic)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Successfully updated mnemonic"))
		return

	case http.MethodDelete:
		//removeMnemonic(mnemonicID)

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
		mnemonicsList, err := getMnemonicList()
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
		// case http.MethodPost:
		//add a new mnemonic
		var newMnemonic Mnemonic
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("body", err.Error())
			return
		}
		err = json.Unmarshal(bodyBytes, &newMnemonic)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("body", err.Error())
			return
		}
		//existingMnem := getMnemonic(newMnemonic.MnemonicID)

		//_, err = addOrUpdateMnenomic(newMnemonic)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		return
	case http.MethodOptions:
		return
	}
}

func validateMnemonic() {

}
