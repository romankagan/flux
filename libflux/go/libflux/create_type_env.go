package libflux

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -llibstd
// #include "flux.h"
// #include <stdlib.h>
import "C"


// EnvStdlib takes care of creating a flux_buffer_t, passes the buffer to
// the Flatbuffers TypeEnvironment and then takes care of freeing the data
func EnvStdlib() []byte {
	var buf C.struct_flux_buffer_t
	C.flux_get_env_stdlib(&buf)
	defer C.flux_free(buf.data)
	return C.GoBytes(buf.data, C.int(buf.len))
}
