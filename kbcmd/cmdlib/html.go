package cmdlib

import (
	"bytes"
	"encoding"
	"errors"
	"fmt"
	"github.com/killbill/kbcli/v2/kbcommon"
	"io"
	"reflect"

	"github.com/go-openapi/runtime"
)

// HTMLConsumer creates a new HTML consumer
func HTMLConsumer() runtime.Consumer {
	return runtime.ConsumerFunc(func(reader io.Reader, data interface{}) error {
		if reader == nil {
			return errors.New("TextConsumer requires a reader") // early exit
		}

		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(reader)
		if err != nil {
			return err
		}
		b := buf.Bytes()

		// If the buffer is empty, no need to unmarshal it, which causes a panic.
		if len(b) == 0 {
			data = ""
			return nil
		}

		if tu, ok := data.(encoding.TextUnmarshaler); ok {
			err := tu.UnmarshalText(b)
			if err != nil {
				return fmt.Errorf("text consumer: %v", err)
			}

			return nil
		}

		t := reflect.TypeOf(data)
		if data != nil && t.Kind() == reflect.Ptr {
			v := reflect.Indirect(reflect.ValueOf(data))
			if t.Elem().Kind() == reflect.String {
				v.SetString(string(b))
				return nil
			}
		}

		// For unhandled response codes, go-swagger passes kbcommon.NewKillbillError as data
		// See https://github.com/go-swagger/go-swagger/issues/1929
		if kbErr, ok := data.(**kbcommon.KillbillError); ok {
			(*kbErr).Message = string(b)
			return nil
		}

		return fmt.Errorf("%v (%T) is not supported by the TextConsumer, %s",
			data, data, "can be resolved by supporting TextUnmarshaler interface")
	})
}
