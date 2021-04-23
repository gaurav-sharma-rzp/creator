package lsystem

type Grammar struct {
	initiator []string            `json:"initiator"`
	rules     map[string][]string `json:"rules"`
	alphabet  []string            `json:"alphabet"`
}

func (g *Grammar) Initiator() []string {
	return g.initiator
}

func (g *Grammar) Alphabets() []string {
	return g.alphabet
}

func (g *Grammar) Rules() map[string][]string {
	return g.rules
}
