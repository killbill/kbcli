package cmdlib

// FormatType represents type of format
type FormatType int

const (
	// FormatTypeDefault - Default format. Uses tabular for single items, and short for collections.
	FormatTypeDefault FormatType = iota

	// FormatTypeTabular - tabular
	FormatTypeTabular FormatType = iota

	// FormatTypeShort - Tabular but sub items are not printed
	FormatTypeShort FormatType = iota

	// FormatTypeList - List of key value pairs (same items as table)
	FormatTypeList FormatType = iota

	// FormatTypeFullJSON - Full JSON
	FormatTypeFullJSON
)

// Scan from string
func (t *FormatType) Scan(str string) {
	switch str {
	case "default":
		*t = FormatTypeDefault
	case "table":
		*t = FormatTypeTabular
	case "short":
		*t = FormatTypeShort
	case "list":
		*t = FormatTypeList
	case "json":
		*t = FormatTypeFullJSON
	}
}

// FormatOptions - options for formatting
type FormatOptions struct {
	Type FormatType

	// NoHeader - skip printing header for tabular, csv format
	NoHeader bool
}

// CustomFormatter function
type CustomFormatter func(v interface{}, fo FormatOptions) Output

// Formatter represents settings for formatting a type.
type Formatter struct {
	// List of columns
	Columns []Column

	// CustomFn - custom format fn. Used if specified.
	CustomFn CustomFormatter

	// SubItems of the resource
	SubItems []SubItem
}

// Column represents column of a resource.
type Column struct {
	// Name of the column
	Name string

	// JsonPath to get the column
	Path string

	// Getter - Custom getter function
	Getter func(interface{}) interface{}
}

// SubItem represents sub collection
type SubItem struct {
	// Name of the collection
	Name string

	// FieldName of the struct that has the sub collection
	FieldName string

	// Formatter overrides the default formatter for this item
	Formatter *Formatter
}
