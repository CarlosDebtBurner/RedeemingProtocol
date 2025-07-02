import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns

data = {
    'Module': ['SovereignGrant', 'TaxOracle', 'DebtIncinerator'],
    'Transactions Processed': [1000, 800, 500],
    'Success Rate (%)': [99.5, 97.2, 95.0]
}

df = pd.DataFrame(data)

plt.figure(figsize=(10, 5))
sns.barplot(x='Module', y='Transactions Processed', data=df)
plt.title('Smart Contract Modules - Transactions Processed')
plt.ylabel('Number of Transactions')
plt.xlabel('Smart Contract Module')
plt.tight_layout()
plt.show()

plt.figure(figsize=(10, 5))
sns.barplot(x='Module', y='Success Rate (%)', data=df)
plt.title('Smart Contract Modules - Success Rate')
plt.ylabel('Success Rate (%)')
plt.xlabel('Smart Contract Module')
plt.ylim(90, 100)
plt.tight_layout()
plt.show()
