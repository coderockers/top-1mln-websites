# NetAPI Top 1 Million Websites

**A free, daily-updated CSV ranking of the one million most popular domains on the Internet.**
Built as a drop-in replacement for the discontinued Alexa Top 1M list. Licensed under [CC BY 4.0](#license) — attribution to NetAPI is required.

| | |
|---|---|
| **Download** | [`https://netapi.com/netapi_top1mln.csv`](https://netapi.com/netapi_top1mln.csv) (≈ 22 MB, plain CSV, no token or registration) |
| **Official page** | [netapi.com/top-1-mln-websites/](https://netapi.com/top-1-mln-websites/) — methodology, last-updated timestamp, translations |
| **Update cadence** | Rebuilt every day; dated snapshots in [Releases](https://github.com/coderockers/top-1mln-websites/releases) |
| **Rows** | 1,000,000 + header |
| **License** | [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/) — free for any use, including commercial, **with attribution** |
| **Publisher** | [NetAPI](https://netapi.com/) |

This repository holds the documentation, ready-to-run download/parse examples (cURL, Python, Node.js, Go, PHP) and daily dated snapshots of the list as GitHub Releases. The current list is always served from the permanent URL above; it is not stored in the git tree.

---

## Why this list exists

Rankings of popular websites are used everywhere: as seeds for web crawls, as allow-lists in security tooling, as sampling frames in research, for benchmarking DNS and TLS deployments, or simply to answer "how popular is this domain?". After Alexa shut down its Top 1M file in 2022, many teams were left without a simple, regularly updated list in the familiar `rank,domain` format. The NetAPI Top 1M fills that gap — same shape, same size, rebuilt daily, free.

## Quick start

```bash
# Download today's list
curl -s https://netapi.com/netapi_top1mln.csv -o netapi_top1mln.csv

# Top 10
head -n 11 netapi_top1mln.csv

# Rank of a domain
grep -m1 ',wikipedia.org$' netapi_top1mln.csv
```

## File format

Plain CSV, UTF-8, one header row, rows sorted by ascending rank:

```
NetAPI_Rank,Domain
1,google.com
2,microsoft.com
3,apple.com
```

| Column | Type | Description |
|---|---|---|
| `NetAPI_Rank` | integer | Rank from 1 (most popular) to 1,000,000. |
| `Domain` | string | Registrable domain without subdomain — `google.com`, not `www.google.com`. Internationalised domains are in punycode. |

Every rank appears exactly once; every domain appears exactly once.

## How the ranking is built

The list combines open popularity inputs that capture both **authority / SEO-style** evidence (link graph) and **traffic-style** evidence (DNS query popularity and real-user Chrome experience data). Ranks from each source are merged with [reciprocal rank fusion](https://en.wikipedia.org/wiki/Reciprocal_rank_fusion) into a single ordered list, so no single noisy source dominates the final order. Hostnames are reduced to registrable domains using NetAPI's TLD map.

The result is a synthetic popularity approximation — not a panel of measured page views — designed for research, security checks, large-scale crawls and Alexa-style Top 1M workflows. The full description lives on the [official page](https://netapi.com/top-1-mln-websites/).

## Examples

Each example downloads the list and shows how to work with it: look up the rank of a domain, print the top N, or count domains per TLD. All run as-is, no API token needed.

| File | Language | Notes |
|---|---|---|
| [examples/curl-example.sh](examples/curl-example.sh) | cURL + shell | download, top 10, rank lookup, top TLDs with standard Unix tools |
| [examples/python-example.py](examples/python-example.py) | Python 3 | `requests` + `csv`; rank lookup and TLD stats |
| [examples/node-example.js](examples/node-example.js) | Node.js 18+ | built-in `fetch`, no dependencies; streams the file line by line |
| [examples/go-example.go](examples/go-example.go) | Go | standard library only (`go run examples/go-example.go`) |
| [examples/php-example.php](examples/php-example.php) | PHP 7.4+ | cURL extension; rank lookup and TLD stats |

Tip: the file is regenerated once a day — cache it locally and fetch it at most once per day.

## Citation and attribution

Attribution is a condition of the license. If you use the NetAPI Top 1M in research, products, security tooling, blog posts or derived datasets, credit **NetAPI** and link to the official page:

> Source: [NetAPI Top 1 Million Websites](https://netapi.com/top-1-mln-websites/), CC BY 4.0.

For academic work:

```bibtex
@misc{netapi_top1m,
  title        = {NetAPI Top 1 Million Websites},
  author       = {{NetAPI}},
  howpublished = {\url{https://netapi.com/top-1-mln-websites/}},
  note         = {Daily ranking, CC BY 4.0. Retrieved YYYY-MM-DD},
  year         = {2026}
}
```

A machine-readable citation is in [CITATION.cff](CITATION.cff) (GitHub shows it under "Cite this repository"). Because the list changes every day, state the retrieval date and keep a dated copy when your work needs to be reproducible.

## FAQ

**Is registration or an API key required?**
No. The file is public; fetch it directly from `https://netapi.com/netapi_top1mln.csv`.

**How often is it updated?**
Daily. The "Last updated" timestamp is shown on the [official page](https://netapi.com/top-1-mln-websites/).

**Are subdomains included?**
No. Every entry is a registrable domain (`example.co.uk`, `google.com`); `www.`, `m.` and other subdomains are folded into their parent domain.

**Can I use it commercially?**
Yes. CC BY 4.0 allows any use, including commercial and derived works, as long as NetAPI is credited with a link.

**Is there a historical archive?**
Yes. A GitHub Action publishes a dated snapshot of the list every day under [Releases](https://github.com/coderockers/top-1mln-websites/releases) (gzip-compressed CSV + SHA-256 checksums), so you can reference or download the exact list from a given day. The live file at `https://netapi.com/netapi_top1mln.csv` always remains the source of truth.

**Where do I report a problem with the list?**
Open an issue here or use [netapi.com/contact-us/](https://netapi.com/contact-us/).

## Related NetAPI data

Free datasets (CC BY 4.0):

- [Compromised domains](https://netapi.com/compromised-urls/) and [compromised IP addresses](https://netapi.com/compromised-ips/) — daily threat-intelligence feeds, current and all-time lists
- [TLD abuse report](https://netapi.com/research/tld-abuse/) and [registrar abuse report](https://netapi.com/research/registrar-abuse/) — which zones and registrars host the most compromised domains; hub at [netapi.com/research/](https://netapi.com/research/)

Domain lists and datasets (subscription):

- [All domain lists](https://netapi.com/lists-all-domains/) — every registered domain in 1,500+ TLDs, one file per zone, plus daily new and deleted domains; [ccTLD lists](https://netapi.com/lists-cctld-domains/)
- [All domain datasets](https://netapi.com/datasets-all-domains/) — domain lists enriched with nameservers, hosting IP and country, emails, phone numbers and Majestic rank; [ccTLD datasets](https://netapi.com/datasets-cctld-domains/)
- Per-zone pages, e.g. [.com domains](https://netapi.com/com-domains/), [.net domains](https://netapi.com/net-domains/), [.org domains](https://netapi.com/org-domains/)
- [Domains by DNS provider](https://netapi.com/dns-providers/) — e.g. [domains on Cloudflare](https://netapi.com/dns-cloudflare/)
- [Domains by registrar (WHOIS)](https://netapi.com/whois-providers/) — with registration and expiration dates
- [Domain lookup](https://netapi.com/lookup-domain/) and [IP lookup](https://netapi.com/lookup-ip/)

Documentation: [API reference](https://netapi.com/help/api/), [FAQ](https://netapi.com/help/faq/), [plans](https://netapi.com/plans/), companion repo [coderockers/netapi](https://github.com/coderockers/netapi).

## License

The **NetAPI Top 1 Million Websites** list is © NetAPI / CodeRockers LTD and published under the [Creative Commons Attribution 4.0 International (CC BY 4.0)](https://creativecommons.org/licenses/by/4.0/) license. You may copy, redistribute, adapt and build upon it for any purpose, including commercially, provided you give appropriate credit to NetAPI with a link to `https://netapi.com/top-1-mln-websites/` and indicate if changes were made. See [LICENSE](LICENSE).

The example code in this repository is released under the MIT License (see [LICENSE](LICENSE)).
