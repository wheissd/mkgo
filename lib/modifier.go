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

	for key := range spec.Paths.Map() {
		if spec.Paths.Find(key).Get != nil {
			f(PathInfo{Op: spec.Paths.Find(key).Get, Method: http.MethodGet, Path: key})
		}
		if spec.Paths.Find(key).Post != nil {
			f(PathInfo{Op: spec.Paths.Find(key).Post, Method: http.MethodPost, Path: key})
		}
		if spec.Paths.Find(key).Put != nil {
			f(PathInfo{Op: spec.Paths.Find(key).Put, Method: http.MethodPut, Path: key})
		}
		if spec.Paths.Find(key).Options != nil {
			f(PathInfo{Op: spec.Paths.Find(key).Options, Method: http.MethodOptions, Path: key})
		}
		if spec.Paths.Find(key).Delete != nil {
			f(PathInfo{Op: spec.Paths.Find(key).Delete, Method: http.MethodDelete, Path: key})
		}
		if spec.Paths.Find(key).Patch != nil {
			f(PathInfo{Op: spec.Paths.Find(key).Patch, Method: http.MethodPatch, Path: key})
		}
	}
}
