package main

// - Implementation -
//
// Valid numbers are in the range 27 to 127.
//
// Clocks must support two modes of computation.
//
// The first mode takes a single parameter specifying the number of balls and reports the number of balls given in the input and the number of days (24-hour periods) which elapse before the clock returns to its initial ordering.
//
//   Sample Input
//   30
//   45
//
//   Output for the Sample Input
//   30 balls cycle after 15 days.
//   45 balls cycle after 378 days.
//
// The second mode takes two parameters, the number of balls and the number of minutes to run for.  If the number of minutes is specified, the clock must run to the number of minutes and report the state of the tracks at that point in a JSON format.
//
//   Sample Input
//   30 325
//
//   Output for the Sample Input
//   {"Min":[],"FiveMin":[22,13,25,3,7],"Hour":[6,12,17,4,15],"Main"
//   [11,5,26,18,2,30,19,8,24,10,29,20,16,21,28,1,23,14,27,9]}

func main() {
	// new clock
	c := clock{}

	// run in mode 1
	c.init(30)
	c.run()

	c.init(45)
	c.run()

	// run in mode 2
	c.init(30, 325)
	c.run()
}
