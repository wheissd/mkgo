package annotations

type annotation struct {
	name string
}

func (a annotation) Name() string {
	return a.name
}

type EdgeConfig struct {
	readMap  modeMap
	writeMap modeMap
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

func (e *EdgeConfig) EnableWrite(opts ...edgeConfigOption) *EdgeConfig {
	setModeMap(true, &e.writeMap, wrapOpts(opts...))
	return e
}

func (e *EdgeConfig) DisableWrite(opts ...entityConfigOption) *EdgeConfig {
	setModeMap(false, &e.writeMap, opts...)
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

func (c *EdgeConfig) GetWriteEnabled(mode string) *bool {
	return c.writeMap.Get(mode)
}
