package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dgageot/demoit/files"
)

// Code returns the tree folder.
func Tree(w http.ResponseWriter, r *http.Request) {
	folder := strings.TrimPrefix(r.URL.Path, "/tree/")

	if !files.Exists(folder) {
		http.NotFound(w, r)
		return
	}

	tree, err := files.Tree(folder)
	if err != nil {
		http.Error(w, "Unable to get tree for folder " + folder, 500)
		return
	}

	b, err := json.Marshal(tree)
	if err != nil {
		http.Error(w, "Unable to convert object to json", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}