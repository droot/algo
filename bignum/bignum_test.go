package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type bigNumTest struct {
	a, b string
	ab   string
}

func TestMultiplyByDigit(t *testing.T) {

	tests := []bigNumTest{
		{
			a:  "1",
			b:  "2",
			ab: "2",
		},
		{
			a:  "1",
			b:  "0",
			ab: "0",
		},
		{
			a:  "5678",
			b:  "1234",
			ab: "7006652",
		},
		{
			a:  "0",
			b:  "1234",
			ab: "0",
		},
	}

	for _, test := range tests {
		a, err := BigNumFrmString(test.a)
		assert.NoError(t, err)
		assert.Equal(t, a.String(), test.a)

		b, err := BigNumFrmString(test.b)
		assert.NoError(t, err)
		assert.Equal(t, b.String(), test.b)

		ab := a.Multiply(b)
		assert.Equal(t, ab.String(), test.ab)
	}
}
