package main

import "testing"

func TestGreeting(t *testing.T) {
	res := greeting("Code.education Rocks!")
	if res != "<b>Code.education Rocks!</b>" {
		t.Error("Erro! correto é <b>Code.education Rocks!</b>", res)
	}
}
