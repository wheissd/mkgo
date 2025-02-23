package annotations

import (
	"github.com/samber/lo"
	"github.com/wheissd/mkgo/internal/config"
)

type Modes []string

func (m Modes) IsGlobalOption() {}
func (m Modes) IsOption()       {}

// If value of map is nil - look at global default
type modeMap struct {
	// setting value in any mode
	global      bool
	globalIsSet bool
	// mode map is set if needed to enable feature in one modes, and disable in others
	mm map[string]*bool
}

func (mm modeMap) Set(mode string, v bool) { mm.mm[mode] = lo.ToPtr(v) }

func (mm modeMap) Get(mode string) *bool {
	if mm.globalIsSet {
		return &mm.global
	}
	return mm.mm[mode]
}

type globalOption interface {
	IsGlobalOption()
}

func setModeMap[T any](v bool, mm *modeMap, opts ...T) {
	globalOpts := filterGlobalOptions(wrapOpts(opts...)...)
	//logger.Get().Info("EnableReadMany", slog.Any("opts", globalOpts))

	var preventDefault bool
	if mm.mm == nil {
		mm.mm = make(map[string]*bool)
	}
	anyMode := true
	for _, opt := range globalOpts {
		if modes, ok := opt.(Modes); ok {
			anyMode = false
			for _, mode := range modes {
				mm.Set(mode, v)
			}
			preventDefault = true
		}
	}
	if anyMode {
		mm.globalIsSet = true
		mm.global = v
		return
	}
	if !preventDefault {
		mm.Set(config.DefaultMode, v)
	}
}

func wrapOpts[T any](opts ...T) []any {
	res := make([]any, 0, len(opts))
	for i := range opts {
		res = append(res, opts[i])
	}
	return res
}

func filterGlobalOptions(opts ...any) []globalOption {
	res := make([]globalOption, 0, len(opts))
	for _, o := range opts {
		if goi, ok := o.(globalOption); ok {
			res = append(res, goi)
		}
	}
	return res
}
