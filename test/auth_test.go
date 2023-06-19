package test

import (
	"github.com/nekidaz/Google-Authentication-oauth/controllers"
	"testing"
)

func TestExtractTokenFromHeader(t *testing.T) {
	got := controllers.ExtractTokenFromHeader("Bearer 456784564867")
	want := "456784564867"

	if got != want {
		t.Error("Test failed")
	}
}
