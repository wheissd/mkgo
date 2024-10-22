package annotations

type annotation struct {
	name string
}

func (a annotation) Name() string {
	return a.name
}

type EdgeConfig struct {
	readMap   modeMap
	createMap modeMap
	updateMap modeMap
	deleteMap modeMap
}

const EdgeConfigID = "EdgeConfig"

func (*EdgeConfig) Name() string {
	return EdgeConfigID
}

func Edge() *EdgeConfig {
	return &EdgeConfig{}
}

type edgeConfigOption interface {
	IsOption()
}

func (e *EdgeConfig) EnableCreate(opts ...edgeConfigOption) *EdgeConfig {
	setModeMap(true, &e.createMap, wrapOpts(opts...))
	return e
}

func (e *EdgeConfig) DisableCreate(opts ...entityConfigOption) *EdgeConfig {
	setModeMap(false, &e.createMap, opts...)
	return e
}

func (e *EdgeConfig) EnableUpdate(opts ...edgeConfigOption) *EdgeConfig {
	setModeMap(true, &e.updateMap, wrapOpts(opts...))
	return e
}

func (e *EdgeConfig) DisableUpdate(opts ...entityConfigOption) *EdgeConfig {
	setModeMap(false, &e.updateMap, opts...)
	return e
}

func (e *EdgeConfig) EnableDelete(opts ...edgeConfigOption) *EdgeConfig {
	setModeMap(true, &e.deleteMap, wrapOpts(opts...))
	return e
}

func (e *EdgeConfig) DisableDelete(opts ...entityConfigOption) *EdgeConfig {
	setModeMap(false, &e.deleteMap, opts...)
	return e
}

func (e *EdgeConfig) EnableRead(opts ...entityConfigOption) *EdgeConfig {
	setModeMap(true, &e.readMap, opts...)
	return e
}

func (e *EdgeConfig) DisableRead(opts ...entityConfigOption) *EdgeConfig {
	setModeMap(false, &e.readMap, opts...)
	return e
}

func (c *EdgeConfig) GetReadEnabled(mode string) *bool {
	return c.readMap.Get(mode)
}

func (c *EdgeConfig) GetCreateEnabled(mode string) *bool {
	return c.createMap.Get(mode)
}

func (c *EdgeConfig) GetUpdateEnabled(mode string) *bool {
	return c.updateMap.Get(mode)
}

func (c *EdgeConfig) GetDeleteEnabled(mode string) *bool {
	return c.deleteMap.Get(mode)
}
