package main

import (
	"testing"
)

func TestTrackInit(t *testing.T) {
	tr := track{}

	tr.init(4)

	if tr.capacity != 4 {
		t.Errorf("Expected capacity to be 4, but got %v", tr.capacity)
	}

	if len(tr.balls) != 0 {
		t.Errorf("Expected capacity to be 0, but got %v", len(tr.balls))
	}

}

func TestTrackLoad(t *testing.T) {
	tr := track{}
	tr.init(1)

	tr.load(ball(1))

	if len(tr.balls) != 1 {
		t.Errorf("Expected capacity to be 1, but got %v", len(tr.balls))
	}

	tr.load(ball(1))

	if len(tr.balls) != 0 {
		t.Errorf("Expected capacity to be 0, but got %v", len(tr.balls))
	}
}
