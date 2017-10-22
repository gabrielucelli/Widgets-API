package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
    _ "github.com/mattn/go-sqlite3"
)

type App struct {
	Router 	*mux.Router
	DB		*sql.DB
}

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
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
    a.Router.HandleFunc("/user", a.getUsers).Methods("GET")
    a.Router.HandleFunc("/user/{id:[0-9]+}", a.getUser).Methods("GET")
    a.Router.HandleFunc("/widget", a.getWidgets).Methods("GET")
    a.Router.HandleFunc("/widget", a.createWidget).Methods("POST")
    a.Router.HandleFunc("/widget/{id:[0-9]+}", a.getWidget).Methods("GET")
    a.Router.HandleFunc("/widget/{id:[0-9]+}", a.updateWidget).Methods("PUT")
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

func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}