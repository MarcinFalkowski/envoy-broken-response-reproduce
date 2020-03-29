package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	/**
	With transfer-encoding: chunked and upgrade: h2 downstream receives invalid response through Envoy
	 */
	http.HandleFunc("/upgrade-h2-chunked", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("upgrade", "h2")
		w.Header().Set("connection", "upgrade")
		w.Header().Set("transfer-encoding", "chunked")
		fmt.Fprintf(w, "Service response /upgrade-h2-chunked\n")
		// ERROR REPRODUCED
	})

	/**
	Without transfer-encoding: chunked, everything works correctly
	 */
	http.HandleFunc("/upgrade-h2", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("upgrade", "h2")
		w.Header().Set("connection", "upgrade")
		fmt.Fprintf(w, "Service response /upgrade-h2\n")
		// working correctly through Envoy
	})

	/**
	Without upgrade, everything works correctly
	 */
	http.HandleFunc("/chunked", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("transfer-encoding", "chunked")
		fmt.Fprintf(w, "Service response /chunked\n")

	})

	log.Printf("Service listening for HTTPS requests...")
	log.Fatal(http.ListenAndServeTLS(":443", "cert.crt", "key.prv", nil))
}
