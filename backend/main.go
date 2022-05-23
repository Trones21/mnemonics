package main

import (
	"fmt"
	"mnemonics/database"
	"mnemonics/mnemonic"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const apiBasePath = "/api"

func main() {
	fmt.Println("Starting Server")
	database.SetupDatabase()
	//Routes in .service.go files
	mnemonic.SetupRoutes(apiBasePath)
	//user.SetupRoutes(apiBasePath)

	//Routes With MiddleWare
	// mnemonicItemHandlerWrapped := http.HandlerFunc(MnemonicItemHandler)
	// mnemonicListHandlerWrapped := http.HandlerFunc(MnemonicListHandler)
	// http.Handle("/mnemonicsList", middlewareHandler(mnemonicListHandlerWrapped))
	// http.Handle("/mnemonic/", middlewareHandler(mnemonicItemHandlerWrapped))

	//Without Middleware
	// http.HandleFunc("/mnemonicsList", MnemonicListHandler)
	// http.HandleFunc("/mnemonic/", MnemonicItemHandler)

	http.ListenAndServe("localhost:5000", nil)

}

//Example Middleware
func middlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before handler")
		handler.ServeHTTP(w, r)
	})
}
