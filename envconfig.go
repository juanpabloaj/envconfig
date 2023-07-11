// Package envconfig gets values from environment variables
package envconfig

import (
	"log"
	"os"
	"strconv"
)

// Config contains the configurations variables
type Config struct {
	Ints                map[string]int
	Strings             map[string]string
	usingDefaultValueFn func(string, interface{})
}

// usingDefaultValueFunc is a function that is called when a variable is not set
func defaultValueShowsMessage(envName string, value interface{}) {
	log.Printf("<%s> is not set, using default value <%v>\n", envName, value)
}

// UsingDefaultValueFunc sets the function that is called when a variable is not set
// or to disable the warning message with a function that does nothing
// func(name string, value interface{}) {}
func (c *Config) UsingDefaultValueFunc(f func(string, interface{})) {
	c.usingDefaultValueFn = f
}

// LoadInt gets an int from the environment variable or it uses the default value
func (c *Config) LoadInt(name string, defaultValue int) error {
	value := os.Getenv(name)
	if value == "" {
		c.usingDefaultValueFn(name, defaultValue)
		c.Ints[name] = defaultValue
		return nil
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return err
	}

	c.Ints[name] = intValue
	return nil
}

// LoadString gets a string from the environment variable
func (c *Config) LoadString(name string, defaultValue string) {
	value := os.Getenv(name)
	if value == "" {
		c.usingDefaultValueFn(name, defaultValue)
		c.Strings[name] = defaultValue
		return
	}
	c.Strings[name] = value
}

// New creates a new Config
func New() *Config {
	return &Config{
		Ints:                make(map[string]int),
		Strings:             make(map[string]string),
		usingDefaultValueFn: defaultValueShowsMessage,
	}
}
