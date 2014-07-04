package main

import (
	"flag"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/golang/glog"
)

var (
	Templates *template.Template
)

func main() {
	listenAddr := flag.String("listen", "127.0.0.1:5000", "[address]:port to listen on")
	mongoURL := flag.String("mongo", "mongodb://localhost:27017/eatem", "URL for the MongoDB database")
	templatesDir := flag.String("templates", "./templates", "path to directory containing HTML templates")
	staticDir := flag.String("static", "./static", "path to directory containing static files")

	var err error

	glog.Infoln("connecting to database")

	glog.Infoln("loading templates")
	glob := filepath.Join(*templatesDir, "*.html")
	Templates, err = template.ParseGlob(filepath.Join(*templatesDir, "*.html"))
	if err != nil {
		glog.Fatalln(err)
	}

	http.Handle("/static/", http.FileServer(http.StripPrefix("/static", http.Dir(*staticDir))))
	http.HandleFunc("/index.html", handleIndex)

	glog.Infoln("listening on", *listenAddr)
	if err := http.ListenAndServe(*listenAddr); err != nil {
		glog.Fatalln(err)
	}
	glog.Infoln("terminated gracefully")
}
