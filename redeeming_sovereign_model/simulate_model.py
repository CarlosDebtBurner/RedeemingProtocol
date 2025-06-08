def simulate_poverty_death(st_supply, demurrage_rate):
    wealth_gini = 0.85  # Current global inequality
    for year in range(1, 11):
        wealth_gini -= (st_supply * demurrage_rate) / 1000
        if wealth_gini < 0.25:
            return f"Poverty collapsed by Year {year}"
    return f"Poverty still present after 10 years. Final Gini: {wealth_gini:.2f}"

# Example run
print(simulate_poverty_death(1e12, 0.5))
