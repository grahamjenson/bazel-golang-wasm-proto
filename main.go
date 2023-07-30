package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/grahamjenson/bazel-golang-wasm-proto/protos/api"
	"github.com/grahamjenson/bazel-golang-wasm-proto/server"
	"github.com/maxence-charriere/go-app/v6/pkg/app"
)

var (
	bootstrapLoc = flag.String("bootstrap-css-path", "", "path to the bootstrap.css file")
	wasmLoc      = flag.String("wasm-path", "", "path to the web app wasm file")
)

func main() {
	flag.Parse()

	// Since locations of these files may vary over bazel versions,
	// this is one way to ensure they keep working.
	if *bootstrapLoc == "" {
		log.Fatalf("The flag --bootstrap-css-path is required.")
	}
	if *wasmLoc == "" {
		log.Fatalf("The flag --bootstrap-css-path is required.")
	}
	mux := http.NewServeMux()

	app := &app.Handler{
		Title:  "EC2Instances",
		Author: "Graham Jenson",
		Styles: []string{"bootstrap.css"},
	}

	mux.HandleFunc("/app.wasm", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, *wasmLoc)
	})

	mux.HandleFunc("/bootstrap.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, *bootstrapLoc)
	})

	// Handle API
	api.RegisterApiHTTPMux(mux, &server.Server{})

	// Handle go-app
	mux.Handle("/", app)

	fmt.Println("starting local server on http://localhost:7000")
	log.Fatal(http.ListenAndServe(":7000", mux))
}
