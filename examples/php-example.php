<?php
/**
 * NetAPI Top 1 Million Websites - download and explore the list in PHP.
 *
 * Free, no token. CC BY 4.0 - credit NetAPI: https://netapi.com/top-1-mln-websites/
 *
 * Requires the cURL extension. Run from the command line: php php-example.php
 * Memory: about 150 MB for the full list in arrays; raise memory_limit if needed.
 */

ini_set('memory_limit', '512M');

$top1mUrl = 'https://netapi.com/netapi_top1mln.csv';

// 1. Download (about 22 MB). The list is rebuilt daily - fetch it at most once a day.
$ch = curl_init($top1mUrl);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_FOLLOWLOCATION, true);
$csv = curl_exec($ch);
$httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
curl_close($ch);

if ($httpCode !== 200 || $csv === false) {
    fwrite(STDERR, "HTTP $httpCode: download failed\n");
    exit(1);
}

$rankByDomain = [];
$tldCounts = [];
$top10 = [];

$lines = explode("\n", $csv);
array_shift($lines); // header: NetAPI_Rank,Domain

foreach ($lines as $line) {
    if ($line === '') {
        continue;
    }
    [$rank, $domain] = explode(',', $line, 2);
    $rank = (int)$rank;

    $rankByDomain[$domain] = $rank;
    if (count($top10) < 10) {
        $top10[] = sprintf('%3d  %s', $rank, $domain);
    }
    $tld = substr($domain, strrpos($domain, '.') + 1);
    $tldCounts[$tld] = ($tldCounts[$tld] ?? 0) + 1;
}

printf("%s rows loaded\n", number_format(count($rankByDomain)));

// 2. Top 10
echo "\nTop 10:\n" . implode("\n", $top10) . "\n";

// 3. Rank lookup
echo "\nRank lookup:\n";
foreach (['wikipedia.org', 'github.com', 'example.com', 'this-domain-does-not-exist.tld'] as $d) {
    printf("%-35s %s\n", $d, $rankByDomain[$d] ?? 'not in Top 1M');
}

// 4. Top 10 TLDs
arsort($tldCounts);
echo "\nTop TLDs:\n";
foreach (array_slice($tldCounts, 0, 10, true) as $tld => $n) {
    printf(".%-10s %9s\n", $tld, number_format($n));
}
