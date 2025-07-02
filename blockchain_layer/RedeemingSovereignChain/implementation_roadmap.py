import matplotlib.pyplot as plt
import pandas as pd
import matplotlib.dates as mdates
from datetime import datetime

# Define project tasks and timelines
tasks = {
    "Sovereign Grant": ("2025-01-01", "2025-03-31"),
    "Tax Oracle": ("2025-04-01", "2025-05-31"),
    "Debt Incinerator": ("2025-06-01", "2025-08-31"),
    "Nation Adoption": ("2026-01-01", "2026-12-31"),
}

# Convert to DataFrame
df = pd.DataFrame([
    {"Task": task, "Start": datetime.strptime(start, "%Y-%m-%d"), "End": datetime.strptime(end, "%Y-%m-%d")}
    for task, (start, end) in tasks.items()
])

# Plot setup
fig, ax = plt.subplots(figsize=(10, 5))
for idx, row in df.iterrows():
    ax.barh(row["Task"], (row["End"] - row["Start"]).days, left=row["Start"], color="steelblue")

# Format x-axis
ax.xaxis.set_major_locator(mdates.MonthLocator(interval=2))
ax.xaxis.set_major_formatter(mdates.DateFormatter("%b %Y"))
plt.xticks(rotation=45)
plt.title("Redeeming Sovereign Model – Deployment Roadmap")
plt.xlabel("Timeline")
plt.ylabel("Project Phases")
plt.tight_layout()
plt.show()
