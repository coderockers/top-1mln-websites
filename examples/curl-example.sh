#!/usr/bin/env sh
# NetAPI Top 1 Million Websites - download and explore the list with cURL
# and standard Unix tools. Free, no token. CC BY 4.0 - credit NetAPI:
# https://netapi.com/top-1-mln-websites/
#
# Usage: sh curl-example.sh

set -e

FILE=netapi_top1mln.csv

# 1. Download today's list (about 22 MB). -z: skip if the local copy is newer.
curl --fail --silent --show-error -z "$FILE" -o "$FILE" https://netapi.com/netapi_top1mln.csv

# 2. Row count (excluding the header)
printf 'Rows: '
tail -n +2 "$FILE" | wc -l

# 3. Top 10
echo 'Top 10:'
head -n 11 "$FILE" | tail -n +2

# 4. Rank of a domain (exact match on the Domain column)
for d in wikipedia.org github.com example.com; do
  rank=$(grep -m1 ",$d\$" "$FILE" | cut -d, -f1)
  printf '%s -> %s
' "$d" "${rank:-not in Top 1M}"
done

# 5. Top 10 TLDs by number of domains in the list
echo 'Top TLDs:'
tail -n +2 "$FILE" | cut -d, -f2 | awk -F. '{print $NF}' | sort | uniq -c | sort -rn | head -n 10
