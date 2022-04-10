/* GOODJOB modified from the standard hello-app */

// [START gke_goobjob_app]
// [START container_goodjob_app]
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// register hello function to handle all requests
	// We will need to pass file size here.
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	// use PORT environment variable, or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

// hello responds to the request with a plain-text "Hello, world" message.
func hello(w http.ResponseWriter, r *http.Request) {
	parm_list, ok := r.URL.Query()["fichero"]

        if !ok || len(parm_list[0]) < 1 {
		log.Println("¡Necesitamos fichero como argumento de la llamada!")
		fmt.Fprintf(w, "¡Necesitamos fichero como argumento de la llamada!\n")
		return
	}

	fichero := parm_list[0]

	httpClient := &http.Client{}
	resp, err := httpClient.Head(fichero)

	if err != nil {
		log.Fatalf("error on HEAD request: %s", err.Error())
	}

	fmt.Printf("Content-Length: %d \n", resp.ContentLength)

	log.Printf("Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()
	fmt.Fprintf(w, "Hola, Goodjob dice que el valor que necesitas es: !\n")
	fmt.Fprintf(w, "FICHERO: %s!\n", fichero)
	fmt.Fprintf(w, "FICHERO content-Length: %d \n", resp.ContentLength)
	fmt.Fprintf(w, "Version: 6.6.6\n")
	fmt.Fprintf(w, "Hostname: %s\n", host)
}

// [END container_goodjob_app]
// [END gke_goodjob_app]
