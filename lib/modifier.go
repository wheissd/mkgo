package lib

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
)

type PathInfo struct {
	Method string
	Op     *openapi3.Operation
	Path   string
}

func AllPaths(spec *openapi3.T, f func(PathInfo)) {
	for key := range spec.Paths {
		if spec.Paths[key].Get != nil {
			f(PathInfo{Op: spec.Paths[key].Get, Method: http.MethodGet, Path: key})
		}
		if spec.Paths[key].Post != nil {
			f(PathInfo{Op: spec.Paths[key].Post, Method: http.MethodPost, Path: key})
		}
		if spec.Paths[key].Put != nil {
			f(PathInfo{Op: spec.Paths[key].Put, Method: http.MethodPut, Path: key})
		}
		if spec.Paths[key].Options != nil {
			f(PathInfo{Op: spec.Paths[key].Options, Method: http.MethodOptions, Path: key})
		}
		if spec.Paths[key].Delete != nil {
			f(PathInfo{Op: spec.Paths[key].Delete, Method: http.MethodDelete, Path: key})
		}
		if spec.Paths[key].Patch != nil {
			f(PathInfo{Op: spec.Paths[key].Patch, Method: http.MethodPatch, Path: key})
		}
	}
}
