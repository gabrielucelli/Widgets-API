package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "strconv"
    "os"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
    "github.com/gorilla/schema"
    "time"
    "github.com/dgrijalva/jwt-go"
    _ "github.com/mattn/go-sqlite3"
    "fmt"
    "strings"
)

type App struct {
	Router 	*mux.Router
	DB		*sql.DB
}

var publicKey = []byte("secret")
var decoder = schema.NewDecoder()

func (a *App) Initialize(dbname string) {

    var err error

    a.DB, err = sql.Open("sqlite3", dbname)

    if err != nil {
        log.Fatal(err)
    }

    a.Router = mux.NewRouter()

    a.initializeRoutes()
}

func (a *App) Run(addr string) { 
	http.ListenAndServe(addr, handlers.CORS()(handlers.LoggingHandler(os.Stdout, a.Router)))
}

func (a *App) initializeRoutes() {
    a.Router.Handle("/user", a.auth(a.getUsers)).Methods("GET")
    a.Router.Handle("/user/{id:[0-9]+}", a.auth(a.getUser)).Methods("GET")
    a.Router.Handle("/widget", a.auth(a.getWidgets)).Methods("GET")
    a.Router.Handle("/widget", a.auth(a.createWidget)).Methods("POST")
    a.Router.Handle("/widget/{id:[0-9]+}", a.auth(a.getWidget)).Methods("GET")
    a.Router.Handle("/widget/{id:[0-9]+}", a.auth(a.updateWidget)).Methods("PUT")
    a.Router.HandleFunc("/token", a.createToken).Methods("POST")
}

func (a *App) getUser(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])

    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid user ID")
        return
    }

    u := user{ID: id}

    if err := u.getUser(a.DB); err != nil {
        switch err {
        	case sql.ErrNoRows:
            	respondWithError(w, http.StatusNotFound, "User not found")
        	default:
            	respondWithError(w, http.StatusInternalServerError, err.Error())
        }
        return
    }

    respondWithJSON(w, http.StatusOK, u)
}

func (a *App) getUsers(w http.ResponseWriter, r *http.Request) {

    users, err := getUsers(a.DB)

    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, users)
}

func (a *App) getWidget(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])

    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid widget ID")
        return
    }

    wid := widget{ID: id}

    if err := wid.getWidget(a.DB); err != nil {
        switch err {
        case sql.ErrNoRows:
            respondWithError(w, http.StatusNotFound, "Widget not found")
        default:
            respondWithError(w, http.StatusInternalServerError, err.Error())
        }
        return
    }

    respondWithJSON(w, http.StatusOK, wid)
}

func (a *App) getWidgets(w http.ResponseWriter, r *http.Request) {

    widgets, err := getWidgets(a.DB)

    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, widgets)
}

func (a *App) createWidget(w http.ResponseWriter, r *http.Request) {

    var wid widget

    decoder := json.NewDecoder(r.Body)

    if err := decoder.Decode(&wid); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    defer r.Body.Close()

    if err := wid.createWidget(a.DB); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusCreated, wid)
}

func (a *App) updateWidget(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)

    id, err := strconv.Atoi(vars["id"])

    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid widget ID")
        return
    }

    var wid widget

    decoder := json.NewDecoder(r.Body)

    if err := decoder.Decode(&wid); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
        return
    }

    defer r.Body.Close()

    wid.ID = id

    if err := wid.updateWidget(a.DB); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, wid)
}

func (a *App) createToken(w http.ResponseWriter, r *http.Request) {

    err := r.ParseForm(); if err != nil {
	        respondWithError(w, http.StatusBadRequest, "ParseForm error")
	      	return
	    }

	var u user

	err = decoder.Decode(&u, r.Form); if err != nil {
    	respondWithError(w, http.StatusBadRequest, "Invalid request payload")
    	return
	}

    err = u.getUser(a.DB)

    if err != nil {
    	respondWithError(w, http.StatusBadRequest, "Invalid user ID")
    	return 
    }

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)

	claims["user"] = u
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

    token.Claims = claims
    
    tokenString, err := token.SignedString(publicKey)

 	if err != nil {
 		respondWithError(w, http.StatusInternalServerError, "Failed to sign token")
 		return
    }

    respondWithJSON(w, http.StatusOK, map[string]string{"token": tokenString})
}

func (a *App) auth(handler func(http.ResponseWriter, *http.Request)) http.Handler {

    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        var token string

        tokens, ok := r.Header["Authorization"]

        if ok && len(tokens) >= 1 {
            token = tokens[0]
            token = strings.TrimPrefix(token, "Bearer ")
        }

        if token == "" {
            respondWithError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
            return
        }

        parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

            if _, ok := token.Method.(*jwt.SigningMethodHMAC); 

            !ok {
                msg := fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
                return nil, msg
            }

            return publicKey, nil

        })

        if err != nil {
        	fmt.Print(err)
            respondWithError(w, http.StatusUnauthorized, "Error parsing token")
            return
        }

        if parsedToken != nil && parsedToken.Valid {
            http.HandlerFunc(handler).ServeHTTP(w, r)
            return
        }

        // Token is invalid
        respondWithError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))

        return
    })
}

func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}