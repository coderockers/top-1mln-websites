"""NetAPI Top 1 Million Websites - download and explore the list in Python.

Free, no token. CC BY 4.0 - credit NetAPI: https://netapi.com/top-1-mln-websites/

Requires: requests (pip install requests)
"""

import csv
import io
from collections import Counter

import requests

TOP1M_URL = "https://netapi.com/netapi_top1mln.csv"

# 1. Download (about 22 MB). The list is rebuilt daily - fetch it at most once a day.
resp = requests.get(TOP1M_URL, timeout=120)
resp.raise_for_status()

rows = list(csv.DictReader(io.StringIO(resp.text)))
print(f"{len(rows):,} rows loaded")

# 2. Top 10
print("\nTop 10:")
for row in rows[:10]:
    print(f'{row["NetAPI_Rank"]:>3}  {row["Domain"]}')

# 3. Rank lookup
rank_by_domain = {row["Domain"]: int(row["NetAPI_Rank"]) for row in rows}
print("\nRank lookup:")
for domain in ("wikipedia.org", "github.com", "example.com", "this-domain-does-not-exist.tld"):
    print(f"{domain:35} {rank_by_domain.get(domain, 'not in Top 1M')}")

# 4. Top 10 TLDs by number of domains in the list
tld_counts = Counter(row["Domain"].rsplit(".", 1)[-1] for row in rows)
print("\nTop TLDs:")
for tld, n in tld_counts.most_common(10):
    print(f".{tld:10} {n:>8,}")

# 5. Save a dated copy for reproducibility
# from datetime import date
# open(f"netapi_top1mln_{date.today():%Y-%m-%d}.csv", "w", encoding="utf-8").write(resp.text)
