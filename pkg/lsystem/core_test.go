package lsystem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type grammar struct {
	initiator []string
	rules     map[string][]string
	alphabet  []string
}

func (tg *grammar) GetInitiator() []string {
	return tg.initiator
}

func (tg *grammar) GetAlphabet() []string {
	return tg.alphabet
}

func (tg *grammar) GetRules() map[string][]string {
	return tg.rules
}

func TestExpandRecursivelyTest(t *testing.T) {
	simpleGrammar := &grammar{
		initiator: []string{
			"A",
		},
		rules: map[string][]string{
			"A": {"A", "B", "B", "C"},
			"B": {"B", "C"},
			"C": {"C", "A"},
		},
		alphabet: []string{
			"A", "B", "C",
		},
	}
	var result []string
	// testing for 0 steps
	result = ExpandRecursively(simpleGrammar, 0)
	assert.Equal(t, []string{"A"}, result)
	// testing for 1 steps
	result = ExpandRecursively(simpleGrammar, 1)
	assert.Equal(t, []string{"A", "B", "B", "C"}, result)
	// testing for 3 steps
	result = ExpandRecursively(simpleGrammar, 3)
	assert.Equal(t, []string{"A", "B", "B", "C", "B", "C", "B", "C", "C", "A", "B", "C", "C", "A", "B", "C", "C", "A", "C", "A", "A", "B", "B", "C"}, result)
}
