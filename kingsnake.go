package kingsnake

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	pkgName    = "Kingsnake"
	pkgVersion = "0.3.0"
)

// Kingsnake eats vipers
type Kingsnake struct {
	configs map[string]*viper.Viper
	names   []string
	files   []string
}

// func (ks Kingsnake) Add(key string, cfg *viper.Viper) {
// 	// ks.configs[key] = cfg
// 	ks.Eat(key, cfg)
// }

// Eat allows kingsnake to consume vipers
func (ks *Kingsnake) Eat(name string, cfg *viper.Viper) {
	ks.configs[name] = cfg
	ks.names = append(ks.names, name)
}

func (ks *Kingsnake) configFileNormalize(cfg *viper.Viper) *viper.Viper {
	var err error
	file := cfg.ConfigFileUsed()
	if len(file) > 0 {
		if filepath.IsAbs(file) == false {
			if file, err = filepath.Abs(file); err != nil {
				log.Panicln(err)
			} else {
				cfg.SetConfigFile(file)
			}
		}
	}

	return cfg
}

// ConfigFileUsed returns a slice of config files referenced by eaten vipers
func (ks *Kingsnake) ConfigFileUsed() []string {
	if len(ks.files) == 0 {
		for name, cfg := range ks.configs {
			ks.configs[name] = ks.configFileNormalize(cfg)
			ks.files = append(ks.files, ks.configs[name].ConfigFileUsed())
		}
	}
	return ks.files
}

// Get returns the value of a key
func (ks *Kingsnake) Get(key string) interface{} {
	for _, name := range ks.names {
		cfg := ks.configs[name]
		if cfg.IsSet(key) {
			return cfg.GetString(key)
		}
	}
	return nil
}

// GetBool returns the value of a key as a system integer
func (ks *Kingsnake) GetBool(key string) bool {
	return ks.Get(key).(bool)
}

// GetFloat64 returns the value of a key as a system integer
func (ks *Kingsnake) GetFloat64(key string) float64 {
	return ks.Get(key).(float64)
}

// GetInt returns the value of a key as a system integer
func (ks *Kingsnake) GetInt(key string) int {
	return ks.Get(key).(int)
}

// GetString returns the value of a key as a string
func (ks *Kingsnake) GetString(key string) string {
	return ks.Get(key).(string)
}

// IsSet returns true if the key is set in any config
func (ks *Kingsnake) IsSet(key string) bool {
	var result = false

	for _, cfg := range ks.configs {
		if cfg.IsSet(key) {
			result = true
			break
		}
	}

	return result
}

// New creates a kingsnake for your enjoyment
func (ks *Kingsnake) New() Kingsnake {
	return Kingsnake{}
}

func (ks *Kingsnake) String() string {
	return fmt.Sprintf("%v", ks.names)
}

// New creates a kingsnake for your enjoyment
func New() Kingsnake {
	return Kingsnake{
		configs: make(map[string]*viper.Viper),
	}
}
