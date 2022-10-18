package config

type Config interface {
	// method for retrieving configuration settings
	GetString(name string) (configValue string, found bool)
	GetInt(name string) (configValue int, found bool)
	GetBool(name string) (configValue bool, found bool)
	GetFloat(name string) (configValue float64, found bool)

	// method for retrieving configuration settings with Default Value
	GetStringDefaultValue(name, defaultValue string) (configValue string)
	GetIntDefaultValue(name string, defaultValue int) (configValue int)
	GetBoolDefaultValue(name string, defaultValue bool) (configValue bool)
	GetFloatDefaultValue(name string, defaultValue float64) (configValue float64)

	// method for Nested Configurations
	GetNested(nestedName string) (nested Config, found bool)
}
