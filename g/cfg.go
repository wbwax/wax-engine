package g

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/wbwax/wax-engine/utils"

	"gopkg.in/yaml.v2"
)

var (
	// Config is a global config variable
	Config *ServerCfg
)

// ServerCfg defines the server config
type ServerCfg struct {
	Log  logCfg     `yaml:"log"`  // log config
	HTTP HTTPConfig `yaml:"http"` // http config
}

// logCfg defines the logger config
type logCfg struct {
	MaxSize    int    `yaml:"max_size"`    // unit: MB
	MaxAge     int    `yaml:"max_age"`     // unit: day
	MaxBackups int    `yaml:"max_backups"` // unit: short
	Level      string `yaml:"level"`       // log level
	Path       string `yaml:"path"`        // path to hold log file
	Encoding   string `yaml:"encoding"`    // json or console
}

// HTTPConfig defines http config
type HTTPConfig struct {
	Enable bool `yaml:"enable"` // start http server when it is true
	Port   int  `yaml:"port"`   // port to listen
}

// LoadServerConfig loads server config file
// @filename: config filename
func LoadServerConfig(filename string) error {
	if filename == "" {
		return errors.New("use -c to specify configuration file")
	}

	if !utils.IsExist(filename) {
		return fmt.Errorf("config file '%s' is not existent", filename)
	}

	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file '%s', err: %s", filename, err.Error())
	}

	var config ServerCfg
	if err := yaml.Unmarshal(fileContent, &config); err != nil {
		return fmt.Errorf("failed to unmarshal yaml file, err: %s", err.Error())
	}

	// succeed
	Config = &config
	return nil
}
