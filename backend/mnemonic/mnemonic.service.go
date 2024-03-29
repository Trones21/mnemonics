package mnemonic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func SetupRoutes(apiBasePath string) {
	handleMnemonics := http.HandlerFunc(MnemonicListHandler)
	handleMnemonic := http.HandlerFunc(MnemonicItemHandler)
	http.Handle(fmt.Sprintf("%s/mnemonicsList", apiBasePath), handleMnemonics)
	http.Handle(fmt.Sprintf("%s/mnemonic/", apiBasePath), handleMnemonic) //cors.Middleware(handleMnemonic))
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
		//THis is a candidate for Monad design (or just factor out into it's own function?)
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
		fmt.Println("Post Mnemonic")
		var newMnemonic Mnemonic
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("Error Reading Body"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(bodyBytes, &newMnemonic)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		rowsCreated, err := AddMnenomic(newMnemonic)
		if err != nil {
			res, _ := json.Marshal(newMnemonic)
			w.Write([]byte(err.Error()))
			w.Write([]byte("\n Verify property names in request object match unmarshhalled object  \n"))
			w.Write([]byte("\n Request Bytes: \n"))
			w.Write(bodyBytes)
			w.Write([]byte("\n Unmarshalled Object: \n"))
			w.Write([]byte(res))

			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Println(rowsCreated)
		//This is actually 0
		// if rowsCreated != 1 {
		// 	fmt.Println(err)
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	return
		// }

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Successfully created mnemonic"))
		return

	case http.MethodDelete:
		err := DeleteMnemonic(mnemonicID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
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
