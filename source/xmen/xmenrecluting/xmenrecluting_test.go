package xmenrecluting_test

import (
	"testing"
	"fmt"
	"xmenrecluting"
)

func TestMutant(t *testing.T) {
	finalResponse := "OK"

	testG := xmenRecluting.isMutant("[]")
	
	if finalResponse == testG {
		t.Log("Los datos esperados son los correctos")
	} else {
		t.Error("Lod datos no coinciden")
		t.Fail()
	}
}