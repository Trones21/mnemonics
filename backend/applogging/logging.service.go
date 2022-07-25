package applogging

import (
	"errors"
	"fmt"
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
		logInfo, err := GetLoggingInfo(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		switch logInfo.EntityType {
		case "mnemonic":
			mnemonic, err := mnemonic.GetMnemonic(logInfo.EntityID)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			if mnemonic == nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("ID Not found"))
				return
			}
		case "collection":
			w.Write([]byte("Collection logging not implemented"))
			return
		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Entity type & ID must be sent after upvotelog? in the form of <entity>=<id> Entity must be of type mnemonic or collection"))
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

//Such a simple operation... splitting strings, why must it be so hard?
//Monads???
func GetLoggingInfo(r *http.Request) (*LoggingInfo, error) {
	logInfo := &LoggingInfo{}
	if strings.Contains(r.URL.RawQuery, "=") && utf8.RuneCountInString(r.URL.RawQuery) >= 3 {
		logInfo.EntityType = strings.Split(r.URL.RawQuery, "=")[0]
		logInfo.EntityID = strings.Split(r.URL.RawQuery, "=")[1]
		return logInfo, nil
	} else {
		err := errors.New("raw query missing = or is too short")
		return nil, err
	}

	// 	if strings.Contains(r.URL.RawQuery, "=") {
	// 	queryParts := strings.Split(r.URL.RawQuery, "=")
	// 	if len(queryParts) < 2 {

	// 	}
	// }
	//actorIsRegistered
	//userInfo
	//IP address

}
