// +build libflux

package libflux

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -llibstd
// #include "flux.h"
// #include <stdlib.h>
import "C"

import (
	"reflect"
	"unsafe"
)

func unwrapBuffer(buf C.struct_flux_buffer_t) ([]byte, func()) {
	sh := new(reflect.SliceHeader)
	sh.Data = uintptr(unsafe.Pointer(buf.data))
	sh.Len = int(C.uint(buf.len))
	sh.Cap = sh.Len
	bs := *(*[]byte)(unsafe.Pointer(sh))
	freeFn := func() {
		C.flux_free(buf.data)
	}
	return bs, freeFn
}
