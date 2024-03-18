// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"
	"github.com/google/uuid"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// DeleteHumanParams is parameters of deleteHuman operation.
type DeleteHumanParams struct {
	// ID of the Human.
	ID uuid.UUID
}

func unpackDeleteHumanParams(packed middleware.Parameters) (params DeleteHumanParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(uuid.UUID)
	}
	return params
}

func decodeDeleteHumanParams(args [1]string, argsEscaped bool, r *http.Request) (params DeleteHumanParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUUID(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ListHumanParams is parameters of listHuman operation.
type ListHumanParams struct {
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

func unpackListHumanParams(packed middleware.Parameters) (params ListHumanParams) {
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Page = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "itemsPerPage",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.ItemsPerPage = v.(OptInt)
		}
	}
	return params
}

func decodeListHumanParams(args [0]string, argsEscaped bool, r *http.Request) (params ListHumanParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.Page.Get(); ok {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           1,
							MaxSet:        false,
							Max:           0,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: itemsPerPage.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "itemsPerPage",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotItemsPerPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotItemsPerPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ItemsPerPage.SetTo(paramsDotItemsPerPageVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.ItemsPerPage.Get(); ok {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           1,
							MaxSet:        true,
							Max:           255,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "itemsPerPage",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// ReadHumanParams is parameters of readHuman operation.
type ReadHumanParams struct {
	// ID of the Human.
	ID uuid.UUID
}

func unpackReadHumanParams(packed middleware.Parameters) (params ReadHumanParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(uuid.UUID)
	}
	return params
}

func decodeReadHumanParams(args [1]string, argsEscaped bool, r *http.Request) (params ReadHumanParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUUID(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UpdateHumanParams is parameters of updateHuman operation.
type UpdateHumanParams struct {
	// ID of the Human.
	ID uuid.UUID
}

func unpackUpdateHumanParams(packed middleware.Parameters) (params UpdateHumanParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(uuid.UUID)
	}
	return params
}

func decodeUpdateHumanParams(args [1]string, argsEscaped bool, r *http.Request) (params UpdateHumanParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToUUID(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}
