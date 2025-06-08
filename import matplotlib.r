import matplotlib.pyplot as plt

def simulate_poverty_death(st_supply, demurrage_rate):
    wealth_gini = 0.85  # Current global inequality
    gini_progress = [wealth_gini]

    for year in range(1, 11):
        wealth_gini -= (st_supply * demurrage_rate) / 1000
        if wealth_gini < 0.25:
            print(f"Poverty collapsed by Year {year}")
            gini_progress.append(max(wealth_gini, 0))
            break
        gini_progress.append(max(wealth_gini, 0))

    else:
        print(f"Poverty still present after 10 years. Final Gini: {wealth_gini:.2f}")

    # Plot and save
    years = list(range(0, len(gini_progress)))
    plt.plot(years, gini_progress, marker='o', color='green')
    plt.xlabel("Year")
    plt.ylabel("Gini Coefficient")
    plt.title("Gini Inequality Simulation")
    plt.grid(True)
    plt.savefig("gini_simulation.png")
    print("Chart saved as gini_simulation.png")

simulate_poverty_death(1e12, 0.5)
