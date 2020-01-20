package main

import (
	"flag"
	"fmt"
	"net/url"

	ga "github.com/mtlynch/google-analytics-v4-example/google_analytics"
)

// coalescePageViews sums together page URLs that have the same path but might
// have different query parameters. For example:
//
// Input:
// {
//   {"/foo?a=b", 2},
//   {"/foo?a=c", 3},
// }
//
// Output:
// {
//   {"/foo", 5},
// }
func coalescePageViews(pvcs []ga.PageViewCount) []ga.PageViewCount {
	totals := map[string]uint64{}
	coalesced := []ga.PageViewCount{}
	for _, pvc := range pvcs {
		u, err := url.Parse(pvc.Path)
		if err != nil {
			panic(err)
		}
		if _, ok := totals[u.EscapedPath()]; ok {
			totals[u.EscapedPath()] += pvc.Views
		} else {
			totals[u.EscapedPath()] = pvc.Views
		}
	}
	for p, c := range totals {
		coalesced = append(coalesced, ga.PageViewCount{p, c})
	}
	return coalesced
}

func main() {
	keyFile := flag.String("keyFile", "", "Path to service account private key JSON")
	viewID := flag.String("viewID", "", "Google Analytics view ID")
	flag.Parse()

	if *keyFile == "" {
		panic("Specify a service account private key JSON file with --keyFile")
	}
	if *viewID == "" {
		panic("Specify a Google Analytics View ID with --viewID")
	}

	mf, err := ga.New(*keyFile, *viewID)
	if err != nil {
		panic("Error while creating Google Analytics Reporting Service")
	}

	pc, err := mf.PageViewsByPath("2019-12-01", "today")
	if err != nil {
		panic(err)
	}

	for _, pvc := range coalescePageViews(pc) {
		fmt.Printf("%s: %d\n", pvc.Path, pvc.Views)
	}
}
