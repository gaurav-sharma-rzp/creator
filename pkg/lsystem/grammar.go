package lsystem

// Grammar interface has functions needed to work with l-system.
// Any model that needs to use l-system for generating and expanding symbols needs to implement this interface
type Grammar interface {
	// Initiator returns the starting symbol to be used for generating the l-system
	GetInitiator() []string

	// Alphabets returns all the valid alphabets that this grammar supports
	GetAlphabet() []string

	// Rules returns the set of rules to follow while expanding the symbol
	GetRules() map[string][]string
}
