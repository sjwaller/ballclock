package main

type track struct {
	capacity int
	balls    []ball
}

// init function
func (t *track) init(capacity int) {
	t.capacity = capacity
	t.balls = []ball{}
}

// load new ball to track
func (t *track) load(b ball) ([]ball, ball) {
	var drop ball
	var discard []ball

	t.balls = append(t.balls, b)

	// should track tilt?
	if len(t.balls) > t.capacity {

		// balls to discard
		discard = t.balls[:len(t.balls)-1]

		// reverse discarded balls
		for i, j := 0, len(discard)-1; i < j; i, j = i+1, j-1 {
			discard[i], discard[j] = discard[j], discard[i]
		}

		// balls to drop
		drop = t.balls[len(t.balls)-1]

		// reset balls
		t.balls = []ball{}
	}

	return discard, drop
}
