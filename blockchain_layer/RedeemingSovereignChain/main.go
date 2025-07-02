package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type Wallet struct {
	Name            string
	SpendingTokens  float64
	SavingsTokens   float64
	TotalTokens     float64
	DebtBurned      bool
	EducationUnlock bool
}

type Transaction struct {
	Name     string
	Item     string
	Amount   float64
	Category string
}

var taxTable = map[string]float64{
	"essential":   0.0,
	"luxury":      0.75,
	"beer":        0.25,
	"cleaning":    0.05,
	"maintenance": 0.10,
}

var behaviorApproved = map[string]bool{
	"Carlos":  true,
	"Maria":   false,
	"Elena":   true,
	"Luis":    true,
	"Zenaida": false,
}

var educationList = map[string]string{
	"Carlos": "Tuition",
	"Elena":  "Seminar",
}

var incomeRegistry = map[string]float64{
	"Carlos":  125000,
	"Maria":   170000,
	"Elena":   145000,
	"Luis":    100000,
	"Zenaida": 88000,
}

var wallets []Wallet
var totalBurned float64

func createWallet(name string) Wallet {
	w := Wallet{
		Name:            name,
		SpendingTokens:  252000,
		SavingsTokens:   28000,
		TotalTokens:     280000,
		DebtBurned:      false,
		EducationUnlock: false,
	}
	fmt.Printf("Wallet Created: %s — Spending: %.2f | Savings: %.2f\n\n", w.Name, w.SpendingTokens, w.SavingsTokens)
	return w
}

func processPurchase(w *Wallet, item string, amount float64, category string) {
	taxRate := taxTable[category]
	tax := amount * taxRate
	total := amount + tax

	if w.SpendingTokens >= total {
		w.SpendingTokens -= total
		fmt.Printf("Purchase: %s bought %s for %.2f (%s)\nTax: %.2f | Total: %.2f | Remaining Spending: %.2f\n\n",
			w.Name, item, amount, category, tax, total, w.SpendingTokens)
	} else {
		fmt.Printf("❌ Not enough tokens for %s to buy %s. Needed: %.2f, Available: %.2f\n\n", w.Name, item, total, w.SpendingTokens)
	}
}

func approveLoan(w *Wallet, amount float64) {
	if behaviorApproved[w.Name] {
		w.SpendingTokens += amount
		fmt.Printf("Loan approved: %.2f tokens added to spending only.\nSpending: %.2f | Total Tokens: %.2f\n\n", amount, w.SpendingTokens, w.TotalTokens)
	} else {
		fmt.Printf("❌ Loan denied: behavior not eligible.\n\n")
	}
}

func repayLoan(w *Wallet, amount float64) {
	if w.SpendingTokens >= amount {
		w.SpendingTokens -= amount
		totalBurned += amount
		w.DebtBurned = true
		fmt.Printf("Loan repaid. Remaining Spending: %.2f\n\n", w.SpendingTokens)
	} else {
		fmt.Println("❌ No loan to repay.")
	}
}

func simulateIncomeTax(w *Wallet, income float64) {
	var tax float64
	switch {
	case income <= 100000:
		tax = income * 0.20
	case income <= 150000:
		tax = income * 0.27
	default:
		tax = income * 0.34
	}
	if w.SpendingTokens >= tax {
		w.SpendingTokens -= tax
		fmt.Printf("Income: %.2f | Income Tax Owed: %.2f | Remaining: %.2f\n\n", income, tax, w.SpendingTokens)
	} else {
		fmt.Printf("Simulated Income: %.2f\nTotal Income Tax Owed: %.2f\nNot enough tokens to deduct income tax.\n\n", income, tax)
	}
}

func unlockEducation(w *Wallet) {
	if _, ok := educationList[w.Name]; ok && w.SavingsTokens >= 5000 {
		w.SavingsTokens -= 5000
		w.EducationUnlock = true
		fmt.Printf("%s used savings for %s. Education unlocked.\n", w.Name, educationList[w.Name])
		w.SpendingTokens += 5000 // Education reward added to spendable tokens
		fmt.Println("🎓 Education reward added: +5000 tokens\n")
	}
}

func giniCoefficient(wallets []Wallet) float64 {
	// Extract real wealths
	values := []float64{}
	for _, w := range wallets {
		v := w.SpendingTokens + w.SavingsTokens
		values = append(values, v)
	}

	// Sort values
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

func redistribute(wallets []Wallet) {
	var total float64
	for _, w := range wallets {
		total += w.SpendingTokens + w.SavingsTokens
	}
	avg := total / float64(len(wallets))
	for i, w := range wallets {
		totalWealth := w.SpendingTokens + w.SavingsTokens
		if totalWealth < avg*0.8 {
			diff := avg*0.8 - totalWealth
			wallets[i].SpendingTokens += diff
		} else if totalWealth > avg*1.2 {
			diff := totalWealth - avg*1.2
			wallets[i].SpendingTokens -= diff
		}
	}
}

func saveWallets(wallets []Wallet) {
	file, _ := os.Create("wallets.json")
	defer file.Close()
	json.NewEncoder(file).Encode(wallets)
	fmt.Println("Wallet data saved to wallets.json")
}

func printGraph(wallets []Wallet) {
	fmt.Println("\n--- Final RSM Report (Visual Summary) ---")
	gini := giniCoefficient(wallets)
	fmt.Printf("Gini Coefficient (equality): %.5f (raw: %.10f)\n", gini, gini)
	equality := int((1 - gini) * 100)
	bar := strings.Repeat("█", equality/4) + strings.Repeat("-", 25-(equality/4))
	fmt.Printf("Equality Level: [%s] %d%% Equality\n\n", bar, equality)

	dist := make(map[int]int)
	for _, w := range wallets {
		total := int(w.SpendingTokens + w.SavingsTokens)
		dist[total]++
	}
	fmt.Println("Token Holdings Distribution:")
	for k, v := range dist {
		fmt.Printf("%d ┤%s  (%d citizens)\n", k, strings.Repeat("■", v), v)
	}

	fmt.Println("\nCitizen Wealths:")
	for _, w := range wallets {
		total := w.SpendingTokens + w.SavingsTokens
		fmt.Printf("%s: %.2f\n", w.Name, total)
	}

	burned, notBurned := 0, 0
	unlocked := 0
	fmt.Println("\nDebt Burn Status:")
	for _, w := range wallets {
		if w.DebtBurned {
			burned++
		} else {
			notBurned++
		}
		if w.EducationUnlock {
			unlocked++
		}
	}
	fmt.Printf("Debt Burned    : %s %d\n", strings.Repeat("█", burned/2), burned)
	fmt.Printf("Debt Not Burned: %s %d\n\n", strings.Repeat("█", notBurned/2), notBurned)
	fmt.Println("Education Unlocks:")
	count := 0
	for _, w := range wallets {
		if count%5 == 0 {
			fmt.Println()
		}
		status := "✘"
		if w.EducationUnlock {
			status = "✔"
		}
		fmt.Printf("%s %-10s ", status, w.Name)
		count++
	}
	fmt.Printf("\nTotal Unlocked: %d / %d\n\n", unlocked, len(wallets))
}

func main() {
	citizenNames := []string{"Carlos", "Maria", "Elena", "Luis", "Zenaida"}
	wallets = nil

	for _, name := range citizenNames {
		w := createWallet(name)
		processPurchase(&w, "Groceries", 1000, "essential")
		processPurchase(&w, "Beer pack", 500, "beer")
		processPurchase(&w, "Office printer", 3000, "essential")
		processPurchase(&w, "Cleaning service", 1500, "cleaning")
		processPurchase(&w, "HVAC Repair", 2000, "maintenance")
		processPurchase(&w, "Designer Shoes", 20000, "luxury")

		approveLoan(&w, 50000)
		repayLoan(&w, 50000)
		unlockEducation(&w)
		simulateIncomeTax(&w, incomeRegistry[w.Name])

		wallets = append(wallets, w)
	}

	redistribute(wallets)
	saveWallets(wallets)
	printGraph(wallets)
}
