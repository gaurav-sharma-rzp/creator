package flora

type Grammar struct {
	initiator []string            `json:"initiator"`
	rules     map[string][]string `json:"rules"`
	alphabet  []string            `json:"alphabet"`
}

func (g *Grammar) GetInitiator() []string {
	return g.initiator
}

func (g *Grammar) GetAlphabet() []string {
	return g.alphabet
}

func (g *Grammar) GetRules() map[string][]string {
	return g.rules
}
