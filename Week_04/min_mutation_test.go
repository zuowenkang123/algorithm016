package Week_04

import (
	"testing"
)

func TestMinMutation(t *testing.T) {
	minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"})
}
