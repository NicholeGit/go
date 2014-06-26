// sqrt_test.go
package mymath

import "testing"

func TestAdd(t *testing.T) {
	v := Add(2, 5)
	if v != 7 {
		t.Errorf("Sqrt(16) failed. Got %v, expected 4.", v)
	}
}
