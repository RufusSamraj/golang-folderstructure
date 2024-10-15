package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"main.go/auth"
	"main.go/contoller"
	"main.go/middleware"
	"main.go/repository"
	"main.go/service"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	// Uncomment below line to activate auto migrate
	//migration.AutoMigrate()

	repo := repository.NewRepository()
	svc := service.NewService(repo)
	ctrl := contoller.NewController(svc)

	authCtrl := auth.AuthController{}

	mainRouter := router.PathPrefix("").Subrouter()
	mainRouter.HandleFunc("/getStudent/{rollNo}", ctrl.GetStudent).Methods(http.MethodGet)
	mainRouter.HandleFunc("/createStudent", ctrl.CreateStudent).Methods(http.MethodPost)

	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/refreshToken", authCtrl.RefreshToken).Methods(http.MethodGet)

	mainRouter.Use(middleware.JwtAuthMiddleware)
	mainRouter.Use(middleware.PanicRecovery)

	listenAddress := ":8080"
	fmt.Println("Starting server on port " + listenAddress)
	log.Printf("Listening on %s. Go to http://127.0.0.1%s", listenAddress, listenAddress)
	log.Println(http.ListenAndServe(listenAddress, enableCors(router)))
}

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Add("Content-Type", "application/json")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
