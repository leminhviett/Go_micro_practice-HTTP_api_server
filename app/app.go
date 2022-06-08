package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/leminhviett/Go_micro_practice-utils_lib/logger"
)

func Start(accHandler AccountHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/accounts", accHandler.GetAllAccounts).
		Methods(http.MethodGet).
		Name("get_all_accounts")
	
	router.HandleFunc("/account/{id:[0-9]+}", accHandler.GetAccountById).
		Methods(http.MethodGet).
		Name("get_specific_account")

	router.HandleFunc("/account", accHandler.RegisterAccount).
		Methods(http.MethodPost).
		Name("create_new_account")


	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	fmt.Println(host, port)
	address := fmt.Sprintf("%s:%s", host, port)
	logger.Info("Running API-server on " + address)
	log.Fatal(http.ListenAndServe(address, router))
}

func UnimplementedHandler(w http.ResponseWriter, r *http.Request)  {
	writeResponse(w, 200, "Good route")
}