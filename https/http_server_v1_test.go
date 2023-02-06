package https

import (
	"log"
	"net/http"
	"testing"
)

// v1
func TestV1(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("httpserver v1"))
	})
	http.HandleFunc("/bye", sayBye1)
	log.Println("Starting v1 server ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sayBye1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye ,this is v1 httpServer"))
}
