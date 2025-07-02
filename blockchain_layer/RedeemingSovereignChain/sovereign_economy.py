import matplotlib.pyplot as plt
import pandas as pd
import numpy as np

# === Citizen Data ===
citizens = [
    {"name": "Carlos", "education": 20000, "luxury": 10000},
    {"name": "Maria", "education": 15000, "luxury": 20000},
    {"name": "Luis", "education": 10000, "luxury": 5000},
    {"name": "Zenaida", "education": 25000, "luxury": 8000},
    {"name": "Jose", "education": 18000, "luxury": 15000},
]

# === Constants ===
GRANT = 280000
SAVINGS_RATE = 0.10
LUXURY_TAX_RATE = 0.75
LUXURY_TAX_THRESHOLD = 12000
EDUCATION_REWARD_RATE = 0.5
MIN_REMAINING = 50000  # Minimum allowed after redistribution

# === Gini Calculation ===
def gini(values):
    sorted_vals = np.sort(values)
    n = len(values)
    index = np.arange(1, n + 1)
    return (2 * np.sum(index * sorted_vals)) / (n * np.sum(sorted_vals)) - (n + 1) / n

# === Redistribution Function ===
def redistribute_to_bottom(df, min_balance=MIN_REMAINING):
    poor = df[df["Remaining Spendable"] < min_balance]
    total_needed = (min_balance - poor["Remaining Spendable"]).sum()
    donors = df[df["Remaining Spendable"] > min_balance]

    if not donors.empty:
        tax_per_donor = total_needed / len(donors)
        df.loc[df["Remaining Spendable"] < min_balance, "Remaining Spendable"] = min_balance
        df.loc[df["Remaining Spendable"] > min_balance, "Remaining Spendable"] -= tax_per_donor

    return df

# === Main Simulation ===
results = []

for c in citizens:
    auto_saved = GRANT * SAVINGS_RATE
    spendable = GRANT - auto_saved

    luxury_excess = max(0, c["luxury"] - LUXURY_TAX_THRESHOLD)
    luxury_taxed = luxury_excess * LUXURY_TAX_RATE
    debt_burned = luxury_taxed

    remaining = spendable - c["education"] - c["luxury"]
    education_reward = c["education"] * EDUCATION_REWARD_RATE
    remaining += education_reward

    results.append({
        "Name": c["name"],
        "Total Grant": GRANT,
        "Auto-Saved (10%)": auto_saved,
        "Spendable (90%)": spendable,
        "Education Spent": c["education"],
        "Luxury Spent (taxed)": luxury_taxed,
        "Debt Burned": debt_burned,
        "Education Reward": education_reward,
        "Remaining Spendable": remaining
    })

df = pd.DataFrame(results)

# === Gini Before Redistribution ===
gini_before = gini(df["Remaining Spendable"])
print(df.to_string(index=False))
print("\nGini Coefficient BEFORE Redistribution:", round(gini_before, 3))

# === Redistribution and Gini After ===
df = redistribute_to_bottom(df)
gini_after = gini(df["Remaining Spendable"])
print("\nGini Coefficient AFTER Redistribution :", round(gini_after, 3))

# === Chart Visualization ===
labels = df["Name"]
auto_saved = df["Auto-Saved (10%)"]
education = df["Education Spent"]
luxury = df["Luxury Spent (taxed)"]
remaining = df["Remaining Spendable"]

fig, ax = plt.subplots(figsize=(10, 6))
bar1 = ax.bar(labels, auto_saved, label='Auto-Saved (10%)')
bar2 = ax.bar(labels, education, bottom=auto_saved, label='Education')
bar3 = ax.bar(labels, luxury, bottom=auto_saved + education, label='Luxury (Taxed)')
bar4 = ax.bar(labels, remaining, bottom=auto_saved + education + luxury, label='Remaining')

ax.set_title("Sovereign Token Allocation per Citizen (With Redistribution)")
ax.set_ylabel("Token Amount")
ax.legend()
plt.tight_layout()
plt.show()
