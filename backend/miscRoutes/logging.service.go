package logging

func SetupRoutes(apiBasePath string) {
	//is there a way to specify where the parameters should be: {entityType}="{id}"
	http.Handle(fmt.Sprintf("%s/viewlog/", apiBasePath), (w http.ResponseWriter, r *http.Request) => {
			
	urlPathSegments := strings.Split(r.URL.Path, "viewlog/")
	entityType := urlPathSemgents['n'] //ToDo: figure out 
	switch entityType{

	}

	})

	http.Handle(fmt.Sprintf("%s/upvotelog", apiBasePath) (w http.ResponseWriter, r *http.Request) => {
	logInfo = getLoggingInfo(r)	
	switch entityType {
	case mnemonic:
		mnemonic, err := GetMnemonic(ID)
		if mnemonic == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("ID Not found"))
			return
		}
	case collection:
			w.Write([]byte("Collection logging not implemented"))
			return
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Entity type & ID must be sent after upvotelog? in the form of <entity>=<id> Entity must be of type mnemonic or collection"))
		return
	}
	})
}

struct LoggingInfo{
	entityType enum ('user', 'mnemonic', 'collection') //how to use enums
	entityID   string//Can i use a uuid type or must it just be a string?
	actorIDorIP string
}

//need to go over structs again... can I use anonymous types?
func getLoggingInfo(r *http.Request) &LoggingInfo{
	logInfo := &LoggingInfo 
	entityInfo := strings.Split(r.URL.Path, "?")[1]
	loggingInfo.entityType := strings.Split(entityInfo, "=")[0]
	entityID := strings.Split(entityInfo, "=")[1] 
	//getuserInfo or IP address
}
