package main

import (
	"net/http"

	"github.com/ilgooz/httpres"
)

var (
	githash    string
	buildstamp string
)

type VersionResponse struct {
	Hash string `json:"hash"`
	Time string `json:"time"`
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	rp := VersionResponse{
		Hash: githash,
		Time: buildstamp,
	}

	httpres.Json(w, http.StatusOK, rp)
}
