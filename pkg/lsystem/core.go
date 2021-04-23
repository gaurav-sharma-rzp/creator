package lsystem

// ExpandRecursively expands the supplied grammar's initiator recursively, steps number of times using the rules specified by the grammar
// If rule is not defined for a alphabet than it is not changed in the expand process
// It returns the final symbol after recursively expanding and any error encountered
func ExpandRecursively(grammar Grammar, steps int) []string {
	symbol := grammar.GetInitiator()
	if steps == 0 || len(symbol) == 0 {
		return symbol
	}
	rules := grammar.GetRules()
	var nextSymbol []string
	for i := 0; i < steps; i++ {
		for _, alphabet := range symbol {
			if val, ok := rules[alphabet]; ok {
				nextSymbol = append(nextSymbol, val...)
			} else {
				nextSymbol = append(nextSymbol, alphabet)
			}
		}
		symbol = nextSymbol
		nextSymbol = nil
	}
	return symbol
}
