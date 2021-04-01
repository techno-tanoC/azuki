package main

import (
	"testing"
)

func TestBuildFreshName(t *testing.T) {
	lc := NewLockCopy(".", 1000, 1000)

	fresh := lc.buildFreshName("go", "mod")
	if fresh != "go(1).mod" {
		t.Errorf("buildFreshName: expect go(1).mod, actual %v", fresh)
	}

	fresh = lc.buildFreshName("unknown", "mod")
	if fresh != "unknown.mod" {
		t.Errorf("buildFreshName: expect unknown.mod, actual %v", fresh)
	}
}
