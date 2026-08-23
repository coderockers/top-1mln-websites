//go:build ignore

// NetAPI Top 1 Million Websites - download and explore the list in Go.
//
// Free, no token. CC BY 4.0 - credit NetAPI: https://netapi.com/top-1-mln-websites/
//
// Standard library only. Usage: go run go-example.go
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

const top1mURL = "https://netapi.com/netapi_top1mln.csv"

func main() {
	// 1. Download (about 22 MB) and parse it as a stream.
	resp, err := http.Get(top1mURL)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Fprintf(os.Stderr, "HTTP %d: %s\n", resp.StatusCode, body)
		os.Exit(1)
	}

	rankByDomain := make(map[string]int, 1_000_000)
	tldCounts := make(map[string]int)
	var top10 []string

	r := csv.NewReader(resp.Body)
	r.FieldsPerRecord = 2
	if _, err := r.Read(); err != nil { // header: NetAPI_Rank,Domain
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		rank, _ := strconv.Atoi(rec[0])
		domain := rec[1]

		rankByDomain[domain] = rank
		if len(top10) < 10 {
			top10 = append(top10, fmt.Sprintf("%3d  %s", rank, domain))
		}
		tld := domain[strings.LastIndex(domain, ".")+1:]
		tldCounts[tld]++
	}

	fmt.Printf("%d rows loaded\n", len(rankByDomain))

	// 2. Top 10
	fmt.Println("\nTop 10:")
	fmt.Println(strings.Join(top10, "\n"))

	// 3. Rank lookup
	fmt.Println("\nRank lookup:")
	for _, d := range []string{"wikipedia.org", "github.com", "example.com", "this-domain-does-not-exist.tld"} {
		if rank, ok := rankByDomain[d]; ok {
			fmt.Printf("%-35s %d\n", d, rank)
		} else {
			fmt.Printf("%-35s not in Top 1M\n", d)
		}
	}

	// 4. Top 10 TLDs
	type kv struct {
		tld string
		n   int
	}
	var tlds []kv
	for t, n := range tldCounts {
		tlds = append(tlds, kv{t, n})
	}
	sort.Slice(tlds, func(i, j int) bool { return tlds[i].n > tlds[j].n })
	fmt.Println("\nTop TLDs:")
	for _, e := range tlds[:10] {
		fmt.Printf(".%-10s %9d\n", e.tld, e.n)
	}
}
