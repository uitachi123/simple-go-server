package api

import (
	"go-server/pkg/db"
	"reflect"
	"testing"
)

func Test_List(t *testing.T) {
	handler, err := db.Init()
	if err != nil {
		t.Errorf("Init db failure %v", err)
	}
	users := list(handler)
	if !reflect.DeepEqual(len(users), 2) {
		t.Errorf("Wrong user counts - expected 2, actual %d", len(users))
	}
}
