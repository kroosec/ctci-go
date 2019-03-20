package main

import (
	"reflect"
	"testing"
)

func AssertInputs(t *testing.T, input, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got '%+v' want '%+v', given '%+v'", got, want, input)
	}
}
