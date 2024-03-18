package annotations

type entityConfigOption interface {
	IsOption()
}

type EntityConfig struct {
	createMap   modeMap
	updateMap   modeMap
	deleteMap   modeMap
	readOneMap  modeMap
	readManyMap modeMap
}

func Entity() *EntityConfig {
	return &EntityConfig{}
}

func (c *EntityConfig) EnableCreate(opts ...entityConfigOption) *EntityConfig {
	setModeMap(true, &c.createMap, opts...)
	return c
}

func (c *EntityConfig) DisableCreate(opts ...entityConfigOption) *EntityConfig {
	setModeMap(false, &c.createMap, opts...)
	return c
}

func (c *EntityConfig) EnableUpdate(opts ...entityConfigOption) *EntityConfig {
	setModeMap(true, &c.updateMap, opts...)
	return c
}

func (c *EntityConfig) DisableUpdate(opts ...entityConfigOption) *EntityConfig {
	setModeMap(false, &c.updateMap, opts...)
	return c
}

func (c *EntityConfig) EnableDelete(opts ...entityConfigOption) *EntityConfig {
	setModeMap(true, &c.deleteMap, opts...)
	return c
}

func (c *EntityConfig) DisableDelete(opts ...entityConfigOption) *EntityConfig {
	setModeMap(false, &c.deleteMap, opts...)
	return c
}

func (c *EntityConfig) EnableReadOne(opts ...entityConfigOption) *EntityConfig {
	setModeMap(true, &c.readOneMap, opts...)
	return c
}

func (c *EntityConfig) DisableReadOne(opts ...entityConfigOption) *EntityConfig {
	setModeMap(false, &c.readOneMap, opts...)
	return c
}

func (c *EntityConfig) EnableReadMany(opts ...entityConfigOption) *EntityConfig {
	setModeMap(true, &c.readManyMap, opts...)
	return c
}

func (c *EntityConfig) DisableReadMany(opts ...entityConfigOption) *EntityConfig {
	setModeMap(false, &c.readManyMap, opts...)
	return c
}

const EntityConfigID = "EntityConfig"

func (c *EntityConfig) Name() string {
	return EntityConfigID
}

func (c *EntityConfig) GetCreateOpEnabled(mode string) *bool {
	return c.createMap.Get(mode)
}

func (c *EntityConfig) GetUpdateOpEnabled(mode string) *bool {
	return c.updateMap.Get(mode)
}

func (c *EntityConfig) GetDeleteOpEnabled(mode string) *bool {
	return c.deleteMap.Get(mode)
}

func (c *EntityConfig) GetReadOneOpEnabled(mode string) *bool {
	return c.readOneMap.Get(mode)
}

func (c *EntityConfig) GetReadManyOpEnabled(mode string) *bool {
	return c.readManyMap.Get(mode)
}
