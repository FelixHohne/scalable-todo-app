package src

import "github.com/gorilla/mux"

func main() {
	router := mux.NewRouter()
	router.StrictSlash(true)

}
