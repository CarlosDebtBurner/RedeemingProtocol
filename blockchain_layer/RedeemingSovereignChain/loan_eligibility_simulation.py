# loan_eligibility_simulation.py

import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns

# Define sample citizens
citizens = [
    {"name": "Carlos", "education": 20000, "luxury": 10000},
    {"name": "Maria", "education": 25000, "luxury": 5000},
    {"name": "Luis", "education": 15000, "luxury": 20000},
    {"name": "Zenaida", "education": 18000, "luxury": 10000},
    {"name": "Jose", "education": 30000, "luxury": 2000},
]

TOTAL_GRANT = 280000
AUTO_SAVED_RATIO = 0.10
LUXURY_TAX_RATE = 0.75

# Eligibility check
def check_loan_eligibility(citizen):
    total_spent = citizen["education"] + citizen["luxury"]
    luxury_ratio = citizen["luxury"] / total_spent if total_spent else 0
    savings_rate = AUTO_SAVED_RATIO
    return luxury_ratio <= 0.25 and savings_rate >= 0.10

# Create DataFrame
df = pd.DataFrame(citizens)
df["Total Grant"] = TOTAL_GRANT
df["Auto-Saved"] = TOTAL_GRANT * AUTO_SAVED_RATIO
df["Spendable"] = TOTAL_GRANT - df["Auto-Saved"]
df["Total Spent"] = df["education"] + df["luxury"]
df["Luxury Ratio"] = df["luxury"] / df["Total Spent"]
df["Loan Eligible"] = df.apply(check_loan_eligibility, axis=1)

# Plot
fig, ax = plt.subplots(figsize=(10, 6))
plt.bar(df["name"], df["Spendable"], color="lightgray", label="Spendable")
plt.bar(df["name"], df["education"], color="orange", label="Education")
plt.bar(df["name"], df["luxury"] * LUXURY_TAX_RATE, bottom=df["education"], color="red", label="Luxury (75% tax)")

for i, eligible in enumerate(df["Loan Eligible"]):
    mark = "\u2713" if eligible else "\u2717"
    color = "green" if eligible else "darkred"
    ax.text(i, TOTAL_GRANT * 0.95, mark, ha="center", fontsize=16, color=color)

plt.title("Loan Eligibility Based on Spending Behavior")
plt.ylabel("Tokens")
plt.legend()
plt.xticks(rotation=45)
plt.tight_layout()
plt.show()
