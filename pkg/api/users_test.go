package api

import (
	"reflect"
	"testing"
)

func Test_List(t *testing.T) {
	users, _ := list()
	if !reflect.DeepEqual(len(users), 2) {
		t.Errorf("Wrong user counts - expected 2, actual %d", len(users))
	}
	if !reflect.DeepEqual("alice@test.com", users[0].Email) {
		t.Errorf("Data mismatch, expected:  %s got: %s", "alice@test.com", users[0].Email)
	}
	if !reflect.DeepEqual("Alice", users[0].Name) {
		t.Errorf("Data mismatch, expected:  %s got: %s", "Alice", users[0].Name)
	}
	if !reflect.DeepEqual("bob@test.com", users[1].Email) {
		t.Errorf("Data mismatch, expected:  %s got: %s", "bob@test.com", users[1].Email)
	}
	if !reflect.DeepEqual("Bob", users[1].Name) {
		t.Errorf("Data mismatch, expected:  %s got: %s", "Bob", users[1].Name)
	}
}
