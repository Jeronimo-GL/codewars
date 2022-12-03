package bouncing

import (
	"testing"
)

/*
 Para ejecutar los tests:
 go test codewars/bouncing
*/
func Test_1(t *testing.T) {
	if BouncingBall(3, 0.66, 1.5) != 3 {
		t.Error("3, 0.66, 1.5 no es 3")
	}
}

func Test_2(t *testing.T) {
	if BouncingBall(40, 0.4, 10) != 3 {
		t.Error("40, 0.4, 10 no es 3")
	}
}

func Test_window_1(t *testing.T) {
	if BouncingBall(5, 0.5, -1.5) != -1 {
		t.Error("Window too low. Must be in (0, h]")
	}
}

func Test_window_2(t *testing.T) {
	if BouncingBall(5, 0.5, 0) != -1 {
		t.Error("Window too low. Must be in (0, h]")
	}
}

func Test_window_3(t *testing.T) {
	if BouncingBall(5, 0.5, 5.1) != -1 {
		t.Error("Window too high. Must be in (0, h]")
	}
}

func Test_bounce_1(t *testing.T) {
	if BouncingBall(5, -1, 1.5) != -1 {
		t.Error("bounce must be between 0 and 1")
	}
}

func Test_bounce_2(t *testing.T) {
	if BouncingBall(5, 0, 1.5) != -1 {
		t.Error("bounce must be in (0, 1)")
	}
}

func Test_bounce_3(t *testing.T) {
	if BouncingBall(5, 1, 1.5) != -1 {
		t.Error("bounce must be in (0, 1)")
	}
}

func Test_bounce_4(t *testing.T) {
	if BouncingBall(5, 2, 1.5) != -1 {
		t.Error("bounce must be in (0, 1)")
	}
}

func Test_h_positive(t *testing.T) {
	if BouncingBall(-5, 10, 1.5) != -1 {
		t.Error("h must be positive")
	}
}
