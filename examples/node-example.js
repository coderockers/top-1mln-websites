// NetAPI Top 1 Million Websites - download and explore the list in Node.js.
//
// Free, no token. CC BY 4.0 - credit NetAPI: https://netapi.com/top-1-mln-websites/
//
// Requires Node.js 18+ (built-in fetch). No dependencies.
// Usage: node node-example.js

const { Readable } = require("node:stream");
const readline = require("node:readline");

const TOP1M_URL = "https://netapi.com/netapi_top1mln.csv";

async function main() {
  // 1. Download (about 22 MB) and read it line by line without buffering the whole file.
  const resp = await fetch(TOP1M_URL);
  if (!resp.ok) {
    throw new Error(`HTTP ${resp.status}: ${await resp.text()}`);
  }

  const rankByDomain = new Map();
  const tldCounts = new Map();
  const top10 = [];

  const lines = readline.createInterface({
    input: Readable.fromWeb(resp.body),
    crlfDelay: Infinity,
  });

  let isHeader = true;
  for await (const line of lines) {
    if (isHeader) { isHeader = false; continue; } // "NetAPI_Rank,Domain"
    if (!line) continue;

    const comma = line.indexOf(",");
    const rank = Number(line.slice(0, comma));
    const domain = line.slice(comma + 1);

    rankByDomain.set(domain, rank);
    if (top10.length < 10) top10.push(`${rank}  ${domain}`);

    const tld = domain.slice(domain.lastIndexOf(".") + 1);
    tldCounts.set(tld, (tldCounts.get(tld) || 0) + 1);
  }

  console.log(`${rankByDomain.size.toLocaleString("en-US")} rows loaded`);

  // 2. Top 10
  console.log("\nTop 10:");
  console.log(top10.join("\n"));

  // 3. Rank lookup
  console.log("\nRank lookup:");
  for (const d of ["wikipedia.org", "github.com", "example.com", "this-domain-does-not-exist.tld"]) {
    console.log(d.padEnd(35), rankByDomain.get(d) ?? "not in Top 1M");
  }

  // 4. Top 10 TLDs
  console.log("\nTop TLDs:");
  [...tldCounts.entries()]
    .sort((a, b) => b[1] - a[1])
    .slice(0, 10)
    .forEach(([tld, n]) => console.log(`.${tld.padEnd(10)} ${n.toLocaleString("en-US").padStart(9)}`));
}

main().catch((err) => {
  console.error(err.message);
  process.exitCode = 1;
});
