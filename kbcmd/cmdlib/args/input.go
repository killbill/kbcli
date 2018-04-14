package args

import (
	"fmt"
	"strings"
)

// Property represents a property
type Property struct {
	Name     string
	Required bool
	Default  string
}

// NameLower returns lowercase  name
func (p Property) NameLower() string {
	return strings.ToLower(p.Name)
}

// Properties - list
type Properties []Property

// ToMap returns map.
func (pl *Properties) ToMap() map[string]Property {
	m := map[string]Property{}
	for _, p := range *pl {
		m[p.NameLower()] = p
	}
	return m
}

// Remove removes the given property from the list
func (pl *Properties) Remove(propName string) {
	propList := []Property(*pl)
	for i, p := range propList {
		if p.NameLower() == strings.ToLower(propName) {
			*pl = append(propList[0:i], propList[i+1:]...)
			return
		}
	}
	panic(fmt.Errorf("property %s not found", propName))
}

// Get returns the property by name.
func (pl *Properties) Get(propName string) *Property {
	propList := (*[]Property)(pl)
	for i := range *propList {
		p := &(*propList)[i]
		if p.NameLower() == strings.ToLower(propName) {
			return p
		}
	}
	panic(fmt.Sprintf("property %s not found", propName))
}

// InputArg - Represents input argument
type InputArg string

// Split splits input into key value pair
func (ia InputArg) Split() (string, string, error) {
	comps := strings.Split(string(ia), "=")
	if len(comps) != 2 {
		return "", "", fmt.Errorf("Invalid input. Expecting PROPERTY=VALUE")
	}
	return comps[0], comps[1], nil
}

// InputArgs - list
type InputArgs []InputArg

// ParseArgs parses given raw input args and returns list of Inputs
func (ial InputArgs) ParseArgs() ([]Input, error) {
	var result []Input
	for _, ia := range ial {
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
