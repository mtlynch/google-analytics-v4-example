package main

import (
	"flag"
	"log"
	"net/url"

	ga "github.com/mtlynch/google-analytics-v4-example/google_analytics"
)

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
	keyFile := flag.String("keyFile", nil, "Path to service account private key JSON")
	viewID := flag.String("viewID", nil, "Google Analytics view ID")
	flag.Parse()

	mf, err := ga.New(*keyFile, *viewID)
	if err != nil {
		panic("Error while creating Google Analytics Reporting Service")
	}

	pc, err := mf.PageViewsByPath("2019-12-01", "today")
	if err != nil {
		panic("getting page counts failed")
	}

	for _, pvc := range coalescePageViews(pc) {
		log.Printf("%03d: %s", pvc.Views, pvc.Path)
	}
}
