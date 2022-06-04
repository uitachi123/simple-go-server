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

func list() ([]db.User, error) {
	d, err := db.Init()
	if err != nil {
		return []db.User{}, err
	}
	txn := d.Txn(false)
	defer txn.Abort()
	iter, err := txn.Get("user", "id")
	if err != nil {
		return []db.User{}, err
	}
	var res []db.User
	for {
		elem := iter.Next()
		if elem == nil {
			break
		}
		res = append(res, *elem.(*db.User))
	}
	return res, nil
}
