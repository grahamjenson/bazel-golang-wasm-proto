package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/grahamjenson/bazel-golang-wasm-proto/protos/api"
	"github.com/grahamjenson/bazel-golang-wasm-proto/server"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

var (
	bootstrapLoc = flag.String("bootstrap-css-path", "", "path to the bootstrap.css file")
	wasmLoc      = flag.String("wasm-path", "", "path to the web app wasm file")
	iconLoc      = flag.String("icon-path", "", "path to the icon")
	port         = flag.Int("port", 7000, "default port to use")
)

func main() {
	app.RouteFunc("/", func() app.Composer {
		// Note that app.Route can not be used to initialize `manager` correctly.
		// This is confusing, but seems to be go-app design choice.
		// See: https://github.com/maxence-charriere/go-app/issues/853
		manager := Manager{
			SearchBar:     &SearchBar{},
			InstanceTable: &InstanceTable{},
		}
		manager.SearchBar.SetManager(&manager)
		manager.InstanceTable.SetManager(&manager)
		return &manager
	})
	app.RunWhenOnBrowser()

	// This is the server part.
	// Unclear why the below has to be in this file, but
	// otherwise it does not work.
	flag.Parse()

	// Since locations of these files may vary over bazel versions,
	// this is one way to ensure they keep working.
	if *bootstrapLoc == "" {
		log.Fatalf("The flag --bootstrap-css-path is required.")
	}
	if *wasmLoc == "" {
		log.Fatalf("The flag --bootstrap-css-path is required.")
	}

	app := &app.Handler{
		Title:  "EC2Instances",
		Author: "Graham Jenson",
		Styles: []string{"/web/bootstrap.css"},
		Icon: app.Icon{
			// Not setting the icon defaults it to go-app icon which is not
			// fetchable due to CORS blocking.
			Default: "/web/icon.png",
		},
	}

	mux := http.NewServeMux()

	// In go-app v9, the static resources *must* be in `/web/...`.
	mux.HandleFunc("/web/app.wasm", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("handling %v\n", *wasmLoc)
		http.ServeFile(w, r, *wasmLoc)
	})

	mux.HandleFunc("/web/icon.png", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("handling %v\n", *iconLoc)
		http.ServeFile(w, r, *iconLoc)
	})

	mux.HandleFunc("/web/bootstrap.css", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("handling %v\n", *bootstrapLoc)
		// go-app v9 requires setting content type on CSS.
		r.Header.Add("Content-Type", "text/css")
		http.ServeFile(w, r, *bootstrapLoc)
	})

	// Handle API
	api.RegisterApiHTTPMux(mux, &server.Server{})

	// Handle go-app
	mux.Handle("/", app)

	log.Printf("starting local server on http://localhost:%v\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *port), mux))
}
