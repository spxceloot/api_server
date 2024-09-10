package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github/luqxus/spxce/service"
	"github/luqxus/spxce/types"
	"io"
	"log"
	"net/http"
	"time"
)

type APIServerConfig struct {
	Port    int
	Host    string
	Service service.Service
}

type APIFunc func(w http.ResponseWriter, r *http.Request) error

type DefaultAPIConfig APIServerConfig

type APIServer struct {
	APIServerConfig
	router *http.ServeMux
}

func New(config APIServerConfig) *APIServer {
	return &APIServer{
		APIServerConfig: config,
		router:          http.NewServeMux(),
	}
}

func (api *APIServer) Run() error {

	api.router.HandleFunc("POST /register", api.handler(api.register))
	api.router.HandleFunc("POST /login", api.handler(api.login))

	log.Printf("Server listening on port:%d\n", api.Port)
	return http.ListenAndServe(fmt.Sprintf("%s:%d", api.Host, api.Port), api.router)
}

func (api *APIServer) handler(fn APIFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
		defer cancel()

		err := fn(w, r.WithContext(ctx))

		if err != nil {
			log.Println(err)
			http.Error(w, "unknown error occured while processing response", http.StatusInternalServerError)
		}
	}
}

func (api *APIServer) root(w http.ResponseWriter, r *http.Request) error {
	// TODO: get location from request if provided
	// THEN: return parking spaces near that location
	var location *types.GeoLocation

	_ = readBody(r.Body, location)
	// ELSE: when location is not provided
	// THEN: return top 5 popular parking spaces

	return nil
}

func (api *APIServer) register(w http.ResponseWriter, r *http.Request) error {

	// request data | CreateUserRequest
	reqData := new(types.CreateUserRequest)

	// parse request body into CreateUserRequest
	if err := readBody(r.Body, reqData); err != nil {
		// on parse failure respond with error
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return nil
	}

	// create new user service, get token or error on failure
	token, err := api.Service.CreateUser(r.Context(), reqData)
	if err != nil {
		// on create user error respond with error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	// set Authorization header to jwt token
	w.Header().Set("Authorization", token)

	// set status code to 200
	w.WriteHeader(http.StatusOK)

	// response with success
	return writeJSON(w, map[string]string{"message": "user created successfully"})

}

func (api *APIServer) login(w http.ResponseWriter, r *http.Request) error {
	// request data storer | LoginRequest
	reqData := new(types.LoginRequest)

	// parse request request data into LoginRequest
	if err := readBody(r.Body, reqData); err != nil {
		// on parse error respond with error
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return nil
	}

	// login user service | get token or error on failure
	token, err := api.Service.Login(r.Context(), reqData)
	if err != nil {
		// on login failure respond with error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	// set Authorization error to jwt token
	w.Header().Set("Authorization", token)

	// set status code to 200
	w.WriteHeader(http.StatusOK)

	// respond with success
	return writeJSON(w, map[string]string{"message": "successfully login"})
}

func readBody(r io.Reader, v any) error {
	return json.NewDecoder(r).Decode(v)
}

func writeJSON(w io.Writer, v any) error {
	return json.NewEncoder(w).Encode(v)
}
