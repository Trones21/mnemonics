package applogging

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mnemonics/mnemonic"
	"net/http"
	"strings"
	"unicode/utf8"
)

func SetupRoutes(apiBasePath string) {
	//is there a way to specify where the parameters should be: {entityType}="{id}"
	// http.Handle(fmt.Sprintf("%s/viewlog/", apiBasePath), (w http.ResponseWriter, r *http.Request) => {
	// logInfo := getLoggingInfo()
	// return

	// })

	http.HandleFunc(fmt.Sprintf("%s/upvotelog", apiBasePath), func(w http.ResponseWriter, r *http.Request) {
		logInfo, err := ParseLoggingInfo(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		switch logInfo.EntityType {
		case "mnemonic":
			mnemonic, err := mnemonic.GetMnemonic(logInfo.EntityID)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			if mnemonic == nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(fmt.Sprintf("ID not found  - mnemonicID: %s", logInfo.EntityID)))
				return
			}
		case "collection":
			// collection, err := mnemonic.GetCollection(logInfo.EntityID)
			// if err != nil {
			// 	w.WriteHeader(http.StatusInternalServerError)
			// 	w.Write([]byte(err.Error()))
			// 	return
			// }
			// if collection == nil {
			// 	w.WriteHeader(http.StatusNotFound)
			// 	w.Write([]byte(fmt.Sprintf("ID not found  - collectionID: %s", logInfo.EntityID)))
			// 	return
			// }
			return
		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`
			JSON body must be in the format {
				actorID: <uuid>
				...
			} `,
			))
			return
		}
	})
}

type LoggingInfo struct {
	EntityType        string //enum ('user', 'mnemonic', 'collection') //
	EntityID          string //Can i use a uuid type or must it just be a string?
	ActorID           string
	IpAddress         string
	IsRegisteredActor bool
}

func ParseLoggingInfo(r *http.Request) (*LoggingInfo, error) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	logInfoB := &LoggingInfo{}
	err = json.Unmarshal(bodyBytes, &logInfoB)
	if err != nil {
		return nil, err
	}
	//Check for required properties?

	return logInfoB, nil

}

func ParsingQueryString(r *http.Request) (*LoggingInfo, error) {
	logInfo := &LoggingInfo{}
	if strings.Contains(r.URL.RawQuery, "=") && utf8.RuneCountInString(r.URL.RawQuery) >= 3 {
		logInfo.EntityType = strings.Split(r.URL.RawQuery, "=")[0]
		logInfo.EntityID = strings.Split(r.URL.RawQuery, "=")[1]
		return logInfo, nil
	} else {
		err := errors.New("raw query missing = or is too short")
		return nil, err
	}
}
