package main


import (
	"github.com/nelsam/goweb_angular_demo/controllers"
	"github.com/stretchr/goweb"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

// main initializes and starts the web server.
func main() {
	controllers.MapControllers()
	MapStaticFiles()

	log.Print("Mapped URLs, creating server...")
	address := ":" + os.Getenv("PORT")
	server := &http.Server{
		Addr:           address,
		Handler:        goweb.DefaultHttpHandler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Print("Created server, listening for requests.")
	listener, listenErr := net.Listen("tcp", address)
	if listenErr != nil {
		log.Printf("Could not listen for tcp connections on address %s: %s", address, listenErr.Error())
		panic(listenErr)
	}
	server.Serve(listener)
}
