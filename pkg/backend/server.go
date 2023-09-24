package backend

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"mime"
	"net/http"
)

type NoteServer struct {
	router    *mux.Router
	NoteStore *NoteStore
}

func CreateServer() NoteServer {
	router := mux.NewRouter()
	router.StrictSlash(true)

	noteStore := createNotStore()
	server := NoteServer{router, noteStore}
	return server
}

func (ns *NoteServer) RegisterRoutes() {

	ns.router.HandleFunc("/note/", ns.createNoteHandler).Methods("POST")
	ns.router.HandleFunc("/task/{id:[0-9]+}/", ns.getNoteHandler).Methods("GET")
	ns.router.HandleFunc("/task/{id:[0-9]+}/", ns.updateNoteHandler).Methods("PUT")
}

func (ns *NoteServer) createNoteHandler(w http.ResponseWriter, req *http.Request) {
	type RequestNote struct {
		Content string   `json:"text"`
		Tags    []string `json:"tags"`
	}

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

}

func (ns *NoteServer) updateNoteHandler(w http.ResponseWriter, req *http.Request) {

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
