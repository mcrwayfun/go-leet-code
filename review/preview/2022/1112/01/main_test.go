package main

import "testing"
import "github.com/stretchr/testify/assert"

func Test_myPow(t *testing.T) {
	assert.Equal(t, float64(1024), myPowTimeOut(2.000000, 10))
	assert.Equal(t, 9.261000000000001, myPowTimeOut(2.1, 3))
	assert.Equal(t, 0.25, myPowTimeOut(2.00000, -2))
}
