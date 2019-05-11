package main

import (
	"fmt"
	"strings"
)

func main() {
	dna := "GATACCA"
	rna := DNAtoRNA(dna)
	fmt.Println(rna)
}

// DNAtoRNA take a dna strengt and return its rna correspondent
func DNAtoRNA(dna string) string {
	translate := func(r rune) rune {
		switch {
		case r == 'T':
			return 'U'
		case r == 'A':
			return 'T'

		case r == 'C':
			return 'G'

		case r == 'G':
			return 'C'
		}
		return 'X'
	}
	rna := strings.Map(translate, dna)
	return rna
}
