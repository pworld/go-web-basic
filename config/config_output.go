package config

import "strings"

type DefaultConfig struct {
	configData map[string]interface{}
}

func (d *DefaultConfig) getValue(name string) (result interface{}, exist bool) {
	data := d.configData

	// loops to separate config name values
	for _, key := range strings.Split(name, ":") {
		result, exist = data[key]
		if newNested, ok := result.(map[string]interface{}); ok && exist {
			data = newNested
		} else {
			return
		}
	}
	return
}

func (d *DefaultConfig) GetNested(name string) (nested Config, exist bool) {
	val, exist := d.getValue(name)
	if exist {
		if nestedData, ok := val.(map[string]interface{}); ok {
			nested = &DefaultConfig{configData: nestedData}
		}
	}
	return
}

func (d *DefaultConfig) GetString(name string) (result string, exist bool) {
	val, exist := d.getValue(name)
	if exist {
		result = val.(string)
	}
	return
}

func (d *DefaultConfig) GetStringDefaultValue(name, defaultVal string) (result string) {
	result, exist := d.GetString(name)
	if !exist {
		result = defaultVal
	}
	return
}

func (d *DefaultConfig) GetInt(name string) (result int, exist bool) {
	val, exist := d.getValue(name)
	if exist {
		result = int(val.(float64))
	}
	return
}

func (d *DefaultConfig) GetIntDefaultValue(name string, defaultVal int) (result int) {
	result, exist := d.GetInt(name)
	if !exist {
		result = defaultVal
	}
	return
}

func (d *DefaultConfig) GetBool(name string) (result bool, exist bool) {
	val, exist := d.getValue(name)
	if exist {
		result = val.(bool)
	}
	return
}

func (d *DefaultConfig) GetBoolDefaultValue(name string, defaultVal bool) (result bool) {
	result, ok := d.GetBool(name)
	if !ok {
		result = defaultVal
	}
	return
}

func (d *DefaultConfig) GetFloat(name string) (result float64, exist bool) {
	val, exist := d.getValue(name)
	if exist {
		result = val.(float64)
	}
	return
}

func (d *DefaultConfig) GetFloatDefaultValue(name string, defaultVal float64) (result float64) {
	result, ok := d.GetFloat(name)
	if !ok {
		result = defaultVal
	}
	return
}
