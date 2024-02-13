package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	min, max, total float64
	n               int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage:   go run naive.go $OUTUT")
		fmt.Fprintln(os.Stderr, "Example: go run naive.go data/2m.csv")
		os.Exit(1)
	}

	inf, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("can't read %v: %v", os.Args[1], err)
	}

	points := map[string]*point{}

	sc := bufio.NewScanner(inf)
	linenr := 0
	for sc.Scan() {
		linenr++
		l := sc.Text()
		if l == "" {
			continue
		}
		parts := strings.Split(l, ";")
		if len(parts) != 2 {
			log.Fatalf("%s:%d %q: need two parts, separated by ';'", os.Args[1], linenr, l)
		}
		t, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			log.Fatalf("%s:%d %q: cannot parse number following ';': %v", os.Args[1], linenr, l, err)
		}
		p, ok := points[parts[0]]
		if !ok {
			points[parts[0]] = &point{min: t, max: t, total: t, n: 1}
		} else {
			p.min = min(p.min, t)
			p.max = max(p.max, t)
			p.total += t
			p.n++
		}
	}

	locs := make([]string, len(points))
	i := 0
	for l := range points {
		locs[i] = l
		i++
	}
	sort.Strings(locs)
	for _, l := range locs {
		fmt.Printf("%-30s %.2f %.2f %.2f\n", l, points[l].min, points[l].total/float64(points[l].n), points[l].max)
	}
}
