package main
import (
	"log"
	"github.com/lcslima45/prolog/internal/server"
	)
func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
