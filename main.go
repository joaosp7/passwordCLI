package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/joaosp7/passwordCLI/filesystem"
	"github.com/joaosp7/passwordCLI/permutation"
	"golang.design/x/clipboard"
)
type Options struct{
	password string
	path string
	copy bool
	one bool
}


func main() {
		opt := Options{}
		flag.StringVar(&opt.password,"password", "", "Password provided for permutation.")
		flag.StringVar(&opt.path,"path", "", "Path provided to save the passwords in a txt file.")
		flag.BoolVar(&opt.copy,"copy", false, "Option to copy password to clipboard.")
		flag.BoolVar(&opt.one, "one", false, "Option to generete only one random password.")
		flag.Parse()

		start := time.Now()

    perm := permutation.NewPermutator()
		if (opt.one){
			result := permutation.OneRandomPermutation(opt.password)
			fmt.Println("One permutation generated: ", result)
			return 
		}
    results, err := perm.FindPermutations(opt.password)

    if err != nil {
        log.Fatalf("Error generating permutations: %v", err)
    }
		if (opt.copy && len(results) > 0){
			clipboard.Write(clipboard.FmtText, []byte(results[0]))
			fmt.Println("Permutation one coppied to clipboard.")
		}
		elapsed := time.Since(start)
		if (len(opt.path) > 0 ){
			errFile := filesystem.GenerateTxt(opt.path, results)
			if errFile!=nil{
				panic("Error creating the txt file. Please try again!")
			}
		}
    //fmt.Printf("Permutations of %q: %s\n", input, strings.Join(results, " "))
		fmt.Println("Permutations generated: ", len(results))
		fmt.Println("Process took around: ", elapsed)
}