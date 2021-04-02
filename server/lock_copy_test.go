package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBuildFreshName(t *testing.T) {
	lc := NewLockCopy(".", 1000, 1000)

	fresh := lc.buildFreshName("go", "mod")
	diff := cmp.Diff(fresh, "go(1).mod")
	if diff != "" {
		t.Errorf("TestBuildFreshName: (-got +want)\n%s", diff)
	}

	fresh = lc.buildFreshName("unknown", "mod")
	diff = cmp.Diff(fresh, "unknown.mod")
	if diff != "" {
		t.Errorf("TestBuildFreshName2: (-got +want)\n%s", diff)
	}
}
