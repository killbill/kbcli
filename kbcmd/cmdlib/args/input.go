package args

import (
	"fmt"
	"strings"
)

// InputArg - Represents command line input.
// This is in the format of: KEY=VALUE
type InputArg string

// Split splits input into key value pair
func (ia InputArg) Split() (string, string, error) {
	comps := strings.Split(string(ia), "=")
	if len(comps) != 2 {
		return "", "", fmt.Errorf("Invalid input %s. Expecting PROPERTY=VALUE", string(ia))
	}
	return comps[0], comps[1], nil
}

// InputArgs - list
type InputArgs []InputArg

// ParseArgs parses given raw input args and returns list of Inputs
func ParseArgs(argsList []string) ([]Input, error) {
	var result []Input
	for _, a := range argsList {
		ia := InputArg(a)
		k, v, err := ia.Split()
		if err != nil {
			return nil, err
		}
		result = append(result, Input{Key: k, Value: v})
	}
	return result, nil
}

// Input represents parsed input arg.
type Input struct {
	Key   string
	Value string
}

// KeyLower returns lowercase  key
func (i Input) KeyLower() string {
	return strings.ToLower(i.Key)
}

// Inputs - list
type Inputs []Input
