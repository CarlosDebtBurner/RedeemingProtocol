package cycles

import (
	"math"
	"sort"
)

// SimulateSpendingAndTax simulates a citizen's token use based on income, loan behavior, and education.
func SimulateSpendingAndTax(name string) float64 {
	initialTokens := 280000.0
	spending := initialTokens * 0.7

	income := map[string]float64{
		"Carlos":  125000,
		"Maria":   170000,
		"Elena":   145000,
		"Luis":    100000,
		"Zenaida": 88000,
	}[name]

	// Education bonus
	if name == "Carlos" || name == "Elena" {
		initialTokens += 5000
	}

	// Loan behavior
	if name == "Carlos" || name == "Elena" || name == "Luis" {
		spending += 50000 // loan granted
		spending -= 50000 // loan repaid
	}

	// Income tax logic
	tax := income * 0.27
	finalTokens := initialTokens - spending + (income - tax)

	return finalTokens
}

// CalculateGini computes the Gini coefficient from the token wealth map.
func CalculateGini(wealth map[string]float64) float64 {
	values := make([]float64, 0, len(wealth))
	for _, v := range wealth {
		values = append(values, v)
	}
	sort.Float64s(values)

	n := float64(len(values))
	sum := 0.0
	for i, v := range values {
		sum += float64(i+1) * v
	}

	total := 0.0
	for _, v := range values {
		total += v
	}

	if total == 0 {
		return 0
	}

	gini := (2*sum)/(n*total) - (n+1)/n
	return math.Max(0, math.Min(1, gini))
}
