package parser

import (
	"fmt"
	"github.com/influxdata/flux/ast"
	"github.com/influxdata/flux/internal/token"
	"github.com/influxdata/flux/libflux/go/libflux"
	"io/ioutil"
	"os"
	"path/filepath"
)

func parseFile(f *token.File, src []byte) (*ast.File, error) {
	astFile := libflux.Parse(string(src))
	defer astFile.Free()

	data, err := astFile.MarshalFB()
	if err != nil {
		return nil, err
	}

	pkg := ast.DeserializeFromFlatBuffer(data)
	file := pkg.Files[0]
	file.Name = f.Name()

	// The go parser will not fill in the imports if there are
	// none so we remove them here to retain compatibility.
	if len(file.Imports) == 0 {
		file.Imports = nil
	}
	return file, nil
}

func parseBytesToHandle(fset *token.FileSet, path string) (*libflux.ASTPkg, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	_ = fset.AddFile(filepath.Base(path), int(fi.Size()))
	src, err := ioutil.ReadFile(path)
	fmt.Println("current file src:", string(src))
	if err != nil {
		return nil, err
	}
	astFile := libflux.Parse(string(src))
	fmt.Println("file", astFile)
	return astFile, nil
}
