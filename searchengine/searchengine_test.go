package searchengine_test

import (
	"reflect"
	"testing"

	se "github.com/kgosse/testforgc/searchengine"
)

func TestNewList(t *testing.T) {
	l := se.NewList()
	want := "searchengine.List"
	if reflect.TypeOf(l).String() != want {
		t.Errorf("want %s, get %s\n", want, reflect.TypeOf(l))
	}
}

func TestAdd(t *testing.T) {
}

func TestGetUniqueElementsCount(t *testing.T) {
}

func TestGetTopElements(t *testing.T) {
}
