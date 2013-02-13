package future

import (
	"testing"
	"time"
)

func add(n int, m int) int {
	time.Sleep(1 * time.Second)

	return n + m
}

func TestNew(t *testing.T) {
	var (
		a Future
		b Future
		f Futurized
	)

	f = New(add)
	start := time.Now()

	a = f(1, 2)
	b = f(3, 4)

	v_a := <-a
	v_b := <-b

	if time.Since(start).Seconds() > 1.5 {
		t.Fatalf("too late")
	}

	if v_a != 3 || v_b != 7 {
		t.Fatalf("invalid value: %d %d", v_a, v_b)
	}
}
