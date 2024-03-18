package annotations

const (
	FieldConfigID = "FieldConfig"
)

type FieldConfig struct {
	// If nil - look at global default
	publicMap modeMap
	// If nil - look at global default
	enableFilter modeMap
}

func (c *FieldConfig) Name() string {
	return FieldConfigID
}

type fieldConfigOption interface {
	IsOption()
}

func (c *FieldConfig) SetPublic(opts ...fieldConfigOption) *FieldConfig {
	setModeMap(true, &c.publicMap, opts...)
	return c
}

func (c *FieldConfig) SetPrivate(opts ...fieldConfigOption) *FieldConfig {
	setModeMap(false, &c.publicMap, opts...)
	return c
}

func Field() *FieldConfig {
	return &FieldConfig{}
}

func (c *FieldConfig) GetPublic(mode string) *bool {
	return c.publicMap.Get(mode)
}

func (c *FieldConfig) EnableFilter(opts ...fieldConfigOption) *FieldConfig {
	setModeMap(true, &c.publicMap, opts...)
	return c
}

func (c *FieldConfig) GetEnableFilter(mode string) *bool {
	return c.publicMap.Get(mode)
}
