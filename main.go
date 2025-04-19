package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/joaosp7/passwordCLI/permutation"
)

func main() {
    perm := permutation.NewPermutator()
    input := "test"
    results, err := perm.FindPermutations(input)
    if err != nil {
        log.Fatalf("Error generating permutations: %v", err)
    }

    fmt.Printf("Permutations of %q: %s\n", input, strings.Join(results, " "))
}