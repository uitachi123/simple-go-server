package api

import (
	"go-server/pkg/db"
	"io"
	"net/http"
	"strings"
)

func Users(w http.ResponseWriter, r *http.Request) {
	args := strings.Split(r.URL.String(), "/")
	io.WriteString(w, args[len(args)-1])
}

func list(handler *db.DBHandler) []db.User {
	var res []db.User
	return res
}
