package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	gotHour, gotMinute, gotSecond := ExtractTimeUnits(3600)
	assert.Equal(t, gotHour, 1)
	assert.Equal(t, gotMinute, 0)
	assert.Equal(t, gotSecond, 0)
	got := ConvertToDigitalFormat(2, 20, 2)
	assert.Equal(t, "02:20:02", got)
}

// func BenchmarkSample(b *testing.B) {
// 	for range b.N {
// 		rand.Int()
// 	}
// }

//t.Error: Logs the error but allows the test to continue running.
//t.Fatal: Logs the error and stops the test immediately.
//go test -run TestAdd
