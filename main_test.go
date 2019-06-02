package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestInitNoArgs(t *testing.T) {

	if os.Getenv("CL_INIT") == "1" {
		c := clock{}
		c.init()
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestInitNoArgs")
	cmd.Env = append(os.Environ(), "CL_INIT=1")
	err := cmd.Run()

	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestInitInvalidArgs(t *testing.T) {

	if os.Getenv("CL_INIT") == "1" {
		c := clock{}
		c.init(1)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestInitInvalidArgs")
	cmd.Env = append(os.Environ(), "CL_INIT=1")
	err := cmd.Run()

	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestInitMaximumArgs(t *testing.T) {

	if os.Getenv("CL_INIT") == "1" {
		c := clock{}
		c.init(1, 2, 3)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestInitInvalidArgs")
	cmd.Env = append(os.Environ(), "CL_INIT=1")
	err := cmd.Run()

	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestInitMode1(t *testing.T) {

	c := clock{}

	c.init(35)

	if c.mode != 1 {
		t.Errorf("Expected mode of 1, but got %v", c.mode)
	}

	if c.numberOfBalls != 35 {
		t.Errorf("Expected numberOfBalls of 35, but got %v", c.numberOfBalls)
	}
}

func TestInitMode2(t *testing.T) {

	c := clock{}

	c.init(35, 200)

	if c.mode != 2 {
		t.Errorf("Expected mode of 2, but got %v", c.mode)
	}

	if c.numberOfBalls != 35 {
		t.Errorf("Expected numberOfBalls of 35, but got %v", c.numberOfBalls)
	}

	if c.limit != 200 {
		t.Errorf("Expected limit of 200, but got %v", c.limit)
	}
}
