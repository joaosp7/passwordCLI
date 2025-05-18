package permutation

import (
	"math/rand"
	"regexp"
	"sort"
	"strings"

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
	regx := regexp.MustCompile(`\s+`)
	input_trim := regx.ReplaceAllString(input, "")
	chars := []rune(input_trim)

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

func OneRandomPermutation(masterPassword string) string {
	masterPasswordsSlice := strings.Split(masterPassword, "")
	

	var builder strings.Builder
	for range masterPasswordsSlice{

		position := rand.Intn(len(masterPasswordsSlice))
		character := masterPasswordsSlice[position]
		builder.WriteString(character)
		lefString := masterPasswordsSlice[:position]
		rightString := masterPasswordsSlice[position:]
		

		masterPasswordsSlice = append(lefString, rightString...)

	}
	result := builder.String()

	return result


}
