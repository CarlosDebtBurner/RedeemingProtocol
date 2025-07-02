package main

import (
	"fmt"
	"rsm/cycles"
	"sort"
	"strings"
)

func main() {
	citizens := []string{"Carlos", "Maria", "Elena", "Luis", "Zenaida"}
	finalWealth := make(map[string]float64)

	for _, name := range citizens {
		wealth := cycles.SimulateSpendingAndTax(name)
		finalWealth[name] = wealth
		fmt.Printf("%s has %.2f tokens remaining.\n", name, wealth)
	}

	sort.Strings(citizens)
	fmt.Println("\n--- Token Holdings Distribution ---")
	for _, name := range citizens {
		value := finalWealth[name]
		bar := bar(value)
		fmt.Printf("%-8s ┤%s  (%.0f tokens)\n", name+":", bar, value)
	}

	gini := cycles.CalculateGini(finalWealth)
	equality := int((1 - gini) * 100)
	giniBar := strings.Repeat("█", equality/5) + strings.Repeat("-", 20-equality/5)
	fmt.Printf("\nGini Coefficient (equality): %.5f\n", gini)
	fmt.Printf("Equality Level: [%s] %d%% Equality\n", giniBar, equality)
}

func bar(value float64) string {
	barLen := int(value / 10000)
	if barLen > 25 {
		barLen = 25
	}
	return strings.Repeat("█", barLen)
}
