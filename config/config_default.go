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

// func (d *DefaultConfig) GetNested(name string) (nested Config, exist bool) {
// 	result, exist := d.getValue(name)
// 	if exist {
// 		if nestedData, ok := result.(map[string]interface{}); ok {
// 			nested = &DefaultConfig{configData: nestedData}
// 		}
// 	}
// 	return
// }

// func (d *DefaultConfig) GetString(name string) (result string, exist bool) {
// 	result, exist := d.getValue(name)
// 	if exist {
// 		result = result.(string)
// 	}
// 	return
// }
