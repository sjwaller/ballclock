package main

import (
	"testing"
)

func TestClockTick(t *testing.T) {

	c := clock{}
	c.init(35)

	c.tick()

	if c.elapsed != 1 {
		t.Errorf("Expected elapsed of 1, but got %v", c.elapsed)
	}

	if len(c.Main) != 34 {
		t.Errorf("Expected Main length of 34, but got %v", len(c.Main))
	}

	if len(c.Min.balls) != 1 {
		t.Errorf("Expected Min length of 1, but got %v", len(c.Min.balls))
	}

	// tick 4 more times (5 mins)
	for i := 1; i <= 4; i++ {
		c.tick()
	}

	if len(c.FiveMin.balls) != 1 {
		t.Errorf("Expected FiveMin length of 1, but got %v", len(c.FiveMin.balls))
	}

	// tick 55 more times (60 mins)
	for i := 1; i <= 55; i++ {
		c.tick()
	}

	if len(c.Hour.balls) != 1 {
		t.Errorf("Expected Hour length of 1, but got %v", len(c.Hour.balls))
	}

}

func TestClockDiscard(t *testing.T) {
	c := clock{}
	c.init(35)

	balls := []ball{ball(36), ball(37), ball(38)}

	c.discard(balls)

	if len(c.Main) != 38 {
		t.Errorf("Expected Main length of 38, but got %v", len(c.Main))
	}
}

func TestClockIsInitial(t *testing.T) {
	c := clock{}
	c.init(35)

	if !c.isInitial() {
		t.Errorf("Expected isInitial to be true, but got %v", c.isInitial())
	}

	c.tick()

	if c.isInitial() {
		t.Errorf("Expected isInitial to be false, but got %v", c.isInitial())
	}

}
