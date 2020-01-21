package semantic

import (
	"github.com/influxdata/flux/libflux/go/libflux"
)

// AnalyzeSource parses and analyzes the given Flux source,
// using libflux.
func AnalyzeSource(fluxSrc string) (*Package, error) {
	ast := libflux.Parse(fluxSrc)
	return AnalyzePackage(ast)
}

func AnalyzePackage(astPkg *libflux.ASTPkg) (*Package, error) {
	defer astPkg.Free()
	sem, err := libflux.Analyze(astPkg)
	if err != nil {
		return nil, err
	}
	defer sem.Free()
	bs, err := sem.MarshalFB()
	if err != nil {
		return nil, err
	}
	return DeserializeFromFlatBuffer(bs)
}


/*
1. flux_get_env_stdlib func def is written in flux.h
1. flux_get_env_stdlib takes in the flux_buffer_t (which is the TypeEnv) and copies it over
1. EnvStdlib called flux_get_env_stdlib which produces a flux_buffer_t
1. EnvStdlib is called in lookup.go and produces a TypeEnvironment
*/

/*
1. flux_merge_ast_pkgs func def is written in flux.h
1. flux_merge_ast_pkgs merges the two files and produces an output file
*/