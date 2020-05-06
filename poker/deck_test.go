package main

import "testing"

func TTestNewDeck(t *testing.T)  {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}
}