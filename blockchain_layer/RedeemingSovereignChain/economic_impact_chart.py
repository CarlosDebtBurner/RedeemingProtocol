import matplotlib.pyplot as plt
import numpy as np

# Economic indicators and values
indicators = ['Wealth Gini', 'Debt-to-GDP', 'Luxury Demand']
current_system = [0.85, 150, 25]
redeeming_model = [0.40, 0, 6]

x = np.arange(len(indicators))  # the label locations
width = 0.35  # the width of the bars

# Create the bar chart
fig, ax = plt.subplots()
bars1 = ax.bar(x - width/2, current_system, width, label='Current System')
bars2 = ax.bar(x + width/2, redeeming_model, width, label='Redeeming Model')

# Add labels, title, and formatting
ax.set_ylabel('Value')
ax.set_title('Projected Economic Impacts (21-Year Horizon)')
ax.set_xticks(x)
ax.set_xticklabels(indicators)
ax.legend()

# Annotate bars with their heights
for bar in bars1 + bars2:
    height = bar.get_height()
    ax.annotate(f'{height}',
                xy=(bar.get_x() + bar.get_width() / 2, height),
                xytext=(0, 3),
                textcoords="offset points",
                ha='center', va='bottom')

plt.tight_layout()
plt.show()
