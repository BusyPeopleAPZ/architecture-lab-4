package main

import (
	"fmt"
	"testing"
	"strings"

	"github.com/BusyPeopleAPZ/architecture-lab-4/engine"
)

var command = "print"
var cntRes engine.Command

func BenchmarkCount(b *testing.B) {
	baseLen := 200000
	for i := 0; i < 13; i++ {
		baseLen = 2 * baseLen
		inputValue := command
		inputValue += strings.Repeat("BBB", baseLen)

		b.Run(fmt.Sprintf("len=%d", baseLen), func(b *testing.B) {
			cntRes = parse(inputValue)
		})
	}
}
