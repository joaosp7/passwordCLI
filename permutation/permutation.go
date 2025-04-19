package permutation

import (
	"sort"

	customerrors "github.com/joaosp7/passwordCLI/errors"
)

type Permutator struct {
	results map[string]struct{} //Set definition

}

func NewPermutator () *Permutator {
	return &Permutator{
		results: make(map[string]struct{}),
	}
}

func (p *Permutator) FindPermutations(input string) ([]string, error) {
	if (len(input) == 0) {
		return nil, customerrors.ErrEmptyInput
	}

	chars := []rune(input)

	sortedChars := make([]rune, len(chars))

	copy(sortedChars, chars)
	sort.Slice(sortedChars, func(i, j int) bool {
		return sortedChars[i] < sortedChars[j]
	})

	p.results = make(map[string]struct{})


	p.recurPermute(chars, 0)

	return p.getSortedResults(), nil

}

func (p *Permutator) recurPermute(chars []rune, index int) {
	if index == len(chars){
		p.results[string(chars)]  = struct{}{}
		return
	}
	
	for i:= index; i < len(chars); i++ {
		chars[index], chars[i] = chars[i], chars[index]

		p.recurPermute(chars, index + 1)

		chars[index], chars[i] = chars[i], chars[index]

	}

}


func (p *Permutator) getSortedResults() []string {
	
	results := make([]string, 0 , len(p.results))
	for result := range p.results {
		results = append(results, result)
	}

	sort.Strings(results)

	return results

}

