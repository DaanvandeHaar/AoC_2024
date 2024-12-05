package main

func Part1() int {
	rules, updates := ReadFile()

	rulesMap := GetRulesMap(rules)

	middles := GetAllowedUpdatesMiddle(rulesMap, updates)

	var total int
	for _, middle := range middles {
		total += middle
	}

	return total
}
