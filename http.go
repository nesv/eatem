package main

import (
	"net/http"

	"github.com/golang/glog"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if err := Templates.ExecuteTemplate(w, "index", nil); err != nil {
		glog.Errorln(err)
	}
}
