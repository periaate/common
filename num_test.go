package common

import (
	"testing"
)

type scExpect struct {
	inp int
	max int
	exp int
}

func TestSmartClamp(t *testing.T) {
	cases := []scExpect{
		{inp: 5, max: 10, exp: 5},
		{inp: 50, max: 5, exp: 5},
		{inp: 5, max: 5, exp: 5},
		{inp: 5, max: 4, exp: 4},
		{inp: 5, max: 0, exp: 0},
		{inp: 0, max: 5, exp: 0},
		{inp: 0, max: 0, exp: 0},
		{inp: -1, max: 10, exp: 9},
		{inp: -1, max: 5, exp: 4},
		{inp: -1, max: 0, exp: 0},
		{inp: -1, max: -10, exp: -1},
		{inp: -1, max: -5, exp: -1},
		{inp: -1, max: -1, exp: -1},
		{inp: 1, max: -10, exp: -9},
		{inp: 1, max: -5, exp: -4},
		{inp: -201, max: 10, exp: 9},
		{inp: 201, max: -10, exp: -9},
	}

	for _, c := range cases {
		got := SmartClamp(c.inp, c.max)
		if got != c.exp {
			t.Errorf("SmartClamp(%d, %d) == %d, want %d", c.inp, c.max, got, c.exp)
		}
	}
}

type absExpect struct {
	inp int
	exp int
}

func TestAbs(t *testing.T) {
	cases := []absExpect{
		{inp: 5, exp: 5},
		{inp: -5, exp: 5},
		{inp: 0, exp: 0},
	}

	for _, c := range cases {
		got := Abs(c.inp)
		if got != c.exp {
			t.Errorf("Abs(%d) == %d, want %d", c.inp, got, c.exp)
		}
	}
}
