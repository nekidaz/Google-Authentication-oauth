package test

import (
	"GoogleAuthv2.0/controllers"
	"testing"
)

func TestExtractTokenFromHeader(t *testing.T) {
	got := controllers.ExtractTokenFromHeader("Bearer 456784564867")
	want := "456784564867"

	if got != want {
		t.Error("Test failed")
	}
}
