package backend

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"mime"
	"net/http"
	"strconv"
)

type NoteServer struct {
	Router    *mux.Router
	NoteStore *NoteStore
}

func CreateServer() NoteServer {
	router := mux.NewRouter()
	router.StrictSlash(true)

	noteStore := createNotStore()
	server := NoteServer{router, noteStore}
	return server
}

type RequestNote struct {
	Content string   `json:"text"`
	Tags    []string `json:"tags"`
}

func (ns *NoteServer) RegisterRoutes() {

	ns.Router.HandleFunc("/note/", ns.createNoteHandler).Methods("POST")
	ns.Router.HandleFunc("/note/{id:[0-9]+}/", ns.getNoteHandler).Methods("GET")
	ns.Router.HandleFunc("/note/{id:[0-9]+}/", ns.updateNoteHandler).Methods("PUT")
}

func (ns *NoteServer) createNoteHandler(w http.ResponseWriter, req *http.Request) {

	type ResponseId struct {
		Id int `json:"id"`
	}

	// Enforce a JSON Content-Type.
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	var rn RequestNote
	if err := dec.Decode(&rn); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := ns.NoteStore.CreateNote(rn.Content, rn.Tags)
	renderJSON(w, ResponseId{Id: id})
}

func (ns *NoteServer) getNoteHandler(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	task, err := ns.NoteStore.GetNote(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	renderJSON(w, task)
}

func (ns *NoteServer) updateNoteHandler(w http.ResponseWriter, req *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	contentType := req.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediaType != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	var rn RequestNote
	if err := dec.Decode(&rn); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ns.NoteStore.UpdateNote(id, rn.Content, rn.Tags)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}

// interface{} means we accept any data type as argument
func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v) // Convert v into json
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
