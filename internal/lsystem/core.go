package lsystem

import "github.com/private/creator/pkg/lsystem"

func Expand(grammar Grammar, steps int) []string {
	// todo need better name to distinguish between pkg and internal l-system
	return lsystem.ExpandRecursively(&grammar, steps)
}

type Grammar struct {
	Initiator []string
	Rules     map[string][]string
	Alphabet  []string
}

func (g *Grammar) GetInitiator() []string {
	return g.Initiator
}

func (g *Grammar) GetAlphabet() []string {
	return g.Alphabet
}

func (g *Grammar) GetRules() map[string][]string {
	return g.Rules
}
