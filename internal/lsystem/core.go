package lsystem

import (
	"github.com/private/creator/internal/logger"
	"github.com/private/creator/pkg/lsystem"
)

func Expand(grammar Grammar, steps int) []string {
	if steps > 15 {
		logger.Info("steps cannot be more than 15")
		steps = 15
	}
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
