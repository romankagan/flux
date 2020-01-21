// +build libflux

package libflux_test

import (
	"fmt"
	"testing"

	"github.com/influxdata/flux/libflux/go/libflux"
)

func TestParse(t *testing.T) {
	text := `
package main

from(bucket: "telegraf")
	|> range(start: -5m)
	|> mean()
`
	ast := libflux.Parse(text)

	jsonBuf, freeFn, err := ast.ToJSON()
	if err != nil {
		panic(err)
	}
	fmt.Printf("json has %v bytes:\n%v\n", len(jsonBuf), string(jsonBuf))
	freeFn()

	fbBuf, offset, freeFn, err := ast.MarshalFB()
	if err != nil {
		panic(err)
	}
	fmt.Printf("flatbuffer has %v bytes, offset %v.\n", len(fbBuf), offset)
	freeFn()
	ast.Free()
}
