package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

type clock struct {
	numberOfBalls int // number of balls
	limit         int // tick limit
	mode          int // mode 1 or 2
	elapsed       int // total minutes elapsed

	Main    []ball // main ball container
	Min     track  // minutes track
	FiveMin track  // five minutes track
	Hour    track  // hours track
}

// init function
//
// accepts 1 or 2 parameters:
//  - init(numberOfBalls int) inits with mode 1
//  - init(numberOfBalls int, limit int) inits with mode 2
func (c *clock) init(params ...int) {

	// check arg length
	switch len(params) {
	case 1:
		c.mode = 1
	case 2:
		c.mode = 2
		c.limit = params[1]
	default:
		fmt.Println("invalid arguements ...")
		os.Exit(1)
	}

	// check ball range
	if (params[0] < 27) || (params[0] > 127) {
		fmt.Println("range must be between 27 - 127")
		os.Exit(1)
	} else {
		c.numberOfBalls = params[0]
	}

	// init clock track
	c.Main = []ball{}
	for i := 1; i <= c.numberOfBalls; i++ {
		c.Main = append(c.Main, ball(i))
	}

	// init other tracks
	c.Min.init(4)
	c.FiveMin.init(11)
	c.Hour.init(11)

	// init elapsed
	c.elapsed = 0
}

// main run loop
func (c *clock) run() {

Loop:
	for {
		c.tick()
		switch c.mode {
		case 1:
			// initial ordering reached
			if c.isInitial() {
				c.printDuration()
				break Loop
			}
			break
		case 2:
			// time limit finished print state
			if c.elapsed >= c.limit {
				c.printJSON()
				break Loop
			}
			break
		}
	}
}

// logic performed every minute
func (c *clock) tick() {
	var discard []ball
	var dropped ball

	// pop the first ball from beginning of Main
	dropped = c.Main[0]

	copy(c.Main, c.Main[1:])
	c.Main = c.Main[:len(c.Main)-1]

	// load Min rail
	discard, dropped = c.Min.load(dropped)
	c.discard(discard)

	if dropped != 0 {
		// load FiveMin rail
		discard, dropped = c.FiveMin.load(dropped)
		c.discard(discard)

		if dropped != 0 {
			// load Hour rail
			discard, dropped = c.Hour.load(dropped)
			c.discard(discard)

			if dropped != 0 {
				discard = []ball{dropped}
				c.discard(discard)
			}
		}
	}

	c.elapsed++
}

// handler for discarded balls
func (c *clock) discard(balls []ball) {
	if balls != nil {
		// append balls to main slice
		c.Main = append(c.Main, balls...)
	}
}

// test if balls reached original order?
func (c clock) isInitial() bool {

	// squash into single slice
	s := append(c.Main, c.Min.balls...)
	s = append(s, c.FiveMin.balls...)
	s = append(s, c.Hour.balls...)

	// compare ball value at index == index + 1
	for i := 0; i < c.numberOfBalls; i = i + 1 {
		if int(s[i]) != i+1 {
			return false
		}
	}

	return true
}

// prints number of elapsed days
func (c clock) printDuration() {
	days := math.Ceil(float64(c.elapsed) * 0.00069444)
	fmt.Println(c.numberOfBalls, "balls cycle after", days, "days.")
}

// prints clock as JSON in desired format
func (c clock) printJSON() {
	b, _ := json.Marshal(&struct {
		Min     []ball
		FiveMin []ball
		Hour    []ball
		Main    []ball
	}{
		Min:     c.Min.balls,
		FiveMin: c.FiveMin.balls,
		Hour:    c.Hour.balls,
		Main:    c.Main,
	})

	fmt.Println(string(b))
}
