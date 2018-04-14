package cmdlib

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/killbill/kbcli/kbclient"
)

// Options for command line
type Options struct {
	Host       string
	Username   string
	Password   string
	CreatedBy  string
	APIKey     string
	APISecret  string
	PrintDebug bool
	Args       []string
	client     *kbclient.KillBill
	Log        Logger
	out        io.Writer
	FO         *FormatOptions
}

// Client returns killbill client
func (o *Options) Client() *kbclient.KillBill {
	return o.client
}

// Output writes the output
func (o *Options) Output(format string, args ...interface{}) {
	o.out.Write([]byte(fmt.Sprintf(format, args...)))
}

// Outputln writes the output and prints new line at the end
func (o *Options) Outputln(format string, args ...interface{}) {
	o.out.Write([]byte(fmt.Sprintf(format, args...) + "\n"))
}

// Print writes formatted output of given resource
func (o *Options) Print(v interface{}) {
	rows, err := getFormattedOutput(o.Log, v, *o.FO, getDefaultFormatter(o.Log, v))
	if err != nil {
		o.out.Write([]byte(fmt.Sprintf("%v\n", err)))
		return
	}
	o.out.Write([]byte(fmt.Sprintf("%s\n", strings.Join(rows, "\n"))))
}

// PrintCustom - print with custom formatter
func (o *Options) PrintCustom(v interface{}, f Formatter) {
	rows, err := getFormattedOutput(o.Log, v, *o.FO, f)
	if err != nil {
		o.out.Write([]byte(fmt.Sprintf("%v\n", err)))
		return
	}
	o.out.Write([]byte(fmt.Sprintf("%s\n", strings.Join(rows, "\n"))))
}

// Format formats given resource and returns as list of lines
func (o *Options) Format(v interface{}, f Formatter) []string {
	rows, err := getFormattedOutput(o.Log, v, *o.FO, f)
	if err != nil {
		panic(err)
	}
	return rows
}

// HandlerFn - signature of handler
type HandlerFn func(ctx context.Context, o *Options) error

// ErrorInvalidArgs - error returned when args are invalid
var ErrorInvalidArgs = fmt.Errorf("invalid args. try -h for syntax")

// Logger - Logging interface.
type Logger interface {
	// Infof - print info.
	Infof(format string, args ...interface{})

	// Warningf - print warning.
	Warningf(format string, args ...interface{})

	// Errorf - print error.
	Errorf(format string, args ...interface{})
}

// MarshalJSON serializes given type to json.
func MarshalJSON(v interface{}) string {
	contents, _ := json.MarshalIndent(v, "", "  ")
	return string(contents)
}
