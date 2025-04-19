package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/joaosp7/passwordCLI/permutation"
	"golang.design/x/clipboard"
)
type Options struct{
	password string
	copy bool
}


func main() {
		opt := Options{}
		flag.StringVar(&opt.password,"password", "", "Password provided for permutation.")
		flag.BoolVar(&opt.copy,"copy", false, "Option to copy password to clipboard.")
		flag.Parse()

		start := time.Now()

    perm := permutation.NewPermutator()
    results, err := perm.FindPermutations(opt.password)

    if err != nil {
        log.Fatalf("Error generating permutations: %v", err)
    }
		if (opt.copy && len(results) > 0){
			clipboard.Write(clipboard.FmtText, []byte(results[0]))
			fmt.Println("Permutation one coppied to clipboard.")
		}
		elapsed := time.Since(start)
    //fmt.Printf("Permutations of %q: %s\n", input, strings.Join(results, " "))
		fmt.Println("Permutations generated: ", len(results))
		fmt.Println("Process took around: ", elapsed)
}