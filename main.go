package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

type InputRow struct {
	Reviewer string `csv:"Email Address"`
	Reviewee string `csv:"Whose code are you reviewing?"`
}

type Reviewee map[string]struct{}
type Reviewer map[string]Reviewee

type OutputRow struct {
	Student string `json:"student"`
	Score   int    `json:"score"`
	Comment string `json:"comment"`
}

func main() {
	// Constants to use for future code review assignments
	inFileName := "project06-responses.csv"
	outFileName := "project06.json"
	points := 26

	// Read in the CSV file
	f, err := os.Open(inFileName)
	if err != nil {
		log.Fatal("os.Open: ", err)
	}
	defer f.Close()
	var inrows []*InputRow
	err = gocsv.Unmarshal(f, &inrows)
	if err != nil {
		log.Fatalln("Unmarshal: ", err)
	}

	// Consolidate the CSV file into a map of reviewer-to-reviewees
	// Hey that's an inverted index!
	reviewers := make(Reviewer)
	for _, inrow := range inrows {
		login, found := strings.CutSuffix((*inrow).Reviewer, "@dons.usfca.edu")
		if !found {
			log.Fatalln("not a dons email: ", (*inrow).Reviewer)
		}
		if _, exists := reviewers[login]; !exists {
			reviewers[login] = make(Reviewee)
		}
		reviewers[login][inrow.Reviewee] = struct{}{}
	}

	// Serialize the inverted index into a slice suitable for JSON output
	outrows := []OutputRow{}
	for reviewer, reviewees := range reviewers {
		outpoints := 0
		outcomments := "reviewed: "
		for reviewee := range reviewees {
			outpoints += points
			outcomments += reviewee + ", "
		}
		outrows = append(outrows, OutputRow{
			reviewer, outpoints, outcomments,
		})
	}
	j, err := json.MarshalIndent(outrows, "", "  ")
	if err != nil {
		log.Fatalln("Marshal out: ", err)
	}

	// Output JSON suitable for "grade upload"
	of, err := os.Create(outFileName)
	if err != nil {
		log.Fatalln("Open outfile: ", err)
	}
	defer of.Close()
	n, err := of.Write(j)
	if err != nil || n != len(j) {
		log.Fatalln("Write failed: ", err)
	}
}
