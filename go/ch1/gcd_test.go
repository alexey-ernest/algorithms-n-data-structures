package ch1

import "testing"

func TestGreaterCommondDivisor(t *testing.T) {
    got := gcd(12, 8)
    if got != 4 {
        t.Errorf("gcd(12, 4) = %d; want 4", got)
    }
}