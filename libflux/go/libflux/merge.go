package libflux

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -llibstd
// #include "flux.h"
// #include <stdlib.h>
import "C"

// Merge packages merges the files of a given input package into a given
// output package.
func MergePackages(inPkg *ASTPkg, outPkg *ASTPkg) *ASTPkg {
	if inPkg == nil {
		return outPkg
	}
	// This modifies outPkg in place prior to return
	C.flux_merge_ast_pkg_files(inPkg.ptr, outPkg.ptr)
	return outPkg
}
