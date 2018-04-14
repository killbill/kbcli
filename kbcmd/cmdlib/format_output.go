package cmdlib

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/oliveagle/jsonpath"
)

// formatRegistry stores the formatter functions
var formatRegistry = map[reflect.Type]Formatter{}

// AddFormatter adds format function to the registry
func AddFormatter(t reflect.Type, f Formatter) {
	formatRegistry[t] = f
}

// OutputRow - A single output row
type OutputRow struct {
	// Values for this output
	Values []string

	// Child items
	Children []Output
}

// Output for a command
type Output struct {
	// Title of the output if any.
	Title string

	// Column names of this output
	Columns []string

	// Rows of this output
	Rows []OutputRow
}

// NewOutput creates new output
func NewOutput(f Formatter) Output {
	o := Output{}
	for _, c := range f.Columns {
		o.Columns = append(o.Columns, c.Name)
	}
	return o
}

// applyFormatter applies the given formatter for the resource.
func (out *Output) applyFormatter(log Logger, v interface{}, fo FormatOptions, f Formatter) error {
	// For custom formatters, just run and return.
	if f.CustomFn != nil {
		*out = f.CustomFn(v, fo)
		return nil
	}

	jsonData, err := toGenericObject(v)
	if err != nil {
		return err
	}

	var row OutputRow

	for _, c := range f.Columns {
		res, _ := jsonpath.JsonPathLookup(jsonData, c.Path)
		row.Values = append(row.Values, fmt.Sprintf("%v", res))
	}

	// Apply formatter for sub items
	for _, subItem := range f.SubItems {
		si, err := getFieldValue(v, subItem.FieldName)
		if err != nil {
			return err
		}
		if si == nil || reflect.ValueOf(si).IsNil() {
			continue
		}
		var sf Formatter
		if subItem.Formatter != nil {
			sf = *subItem.Formatter
		} else {
			sf = getDefaultFormatter(log, si)
		}

		subOut := NewOutput(sf)
		subOut.Title = subItem.Name
		err = subOut.process(log, si, fo, sf)
		if err != nil {
			return err
		}
		row.Children = append(row.Children, subOut)
	}

	out.Rows = append(out.Rows, row)

	return nil
}

// process applies formatters recursively and populates the Output structure.
func (out *Output) process(log Logger, v interface{}, fo FormatOptions, f Formatter) error {
	if reflect.TypeOf(v).Kind() == reflect.Slice {
		s := reflect.ValueOf(v)
		for i := 0; i < s.Len(); i++ {
			err := out.process(log, s.Index(i).Interface(), fo, f)
			if err != nil {
				return err
			}
		}
		return nil
	}

	err := out.applyFormatter(log, v, fo, f)
	if err != nil {
		log.Warningf("unable to format output. %v\n", err)
		return fmt.Errorf("unable to format output. %v", err)
	}

	return nil
}

// getFormattedOutput returns finally formatted output as list of lines.
func getFormattedOutput(log Logger, v interface{}, fo FormatOptions, f Formatter) ([]string, error) {
	out := NewOutput(f)
	if fo.Type == FormatTypeFullJSON {
		return strings.Split(MarshalJSON(v), "\n"), nil
	}

	if fo.Type == FormatTypeDefault {
		if reflect.TypeOf(v).Kind() == reflect.Slice {
			fo.Type = FormatTypeShort
		} else {
			fo.Type = FormatTypeTabular
		}
	}

	err := out.process(log, v, fo, f)
	if err != nil {
		return nil, err
	}

	return printColumns(out, fo, "")
}
