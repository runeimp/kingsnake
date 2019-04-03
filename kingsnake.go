package kingsnake

import (
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	pkgName    = "Kingsnake"
	pkgVersion = "0.2.0"
)

// Kingsnake eats vipers
type Kingsnake struct {
	configs map[string]*viper.Viper
	files   []string
}

// func (ks Kingsnake) Add(key string, cfg *viper.Viper) {
// 	// ks.configs[key] = cfg
// 	ks.Eat(key, cfg)
// }

// Eat allows kingsnake to consume vipers
func (ks Kingsnake) Eat(key string, cfg *viper.Viper) {
	// if err := cfg.ReadInConfig(); err == nil {
	// 	log.Println("Kingsnake.Eat() | cfg file:", cfg.ConfigFileUsed())
	// 	// log.Println()
	// }
	ks.configs[key] = cfg
}

func (ks Kingsnake) configFileNormalize(cfg *viper.Viper) *viper.Viper {
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
func (ks Kingsnake) ConfigFileUsed() []string {
	if len(ks.files) == 0 {
		for key, cfg := range ks.configs {
			ks.configs[key] = ks.configFileNormalize(cfg)
			ks.files = append(ks.files, ks.configs[key].ConfigFileUsed())
		}
	}
	return ks.files
}

// Get returns the value of a key
func (ks Kingsnake) Get(key string) interface{} {
	var result interface{}

	for _, cfg := range ks.configs {
		if cfg.IsSet(key) {
			result = cfg.Get(key)
			break
		}
	}

	return result
}

// GetInt returns the value of a key as a system integer
func (ks Kingsnake) GetInt(key string) int {
	var result int

	for _, cfg := range ks.configs {
		if cfg.IsSet(key) {
			result = cfg.GetInt(key)
			break
		}
	}

	return result
}

// GetString returns the value of a key as a string
func (ks Kingsnake) GetString(key string) string {
	var result = ""

	for _, cfg := range ks.configs {
		if cfg.IsSet(key) {
			result = cfg.GetString(key)
			break
		}
	}

	return result
}

// IsSet returns true if the key is set in any config
func (ks Kingsnake) IsSet(key string) bool {
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
func (ks Kingsnake) New() Kingsnake {
	return Kingsnake{}
}

// New creates a kingsnake for your enjoyment
func New() Kingsnake {
	return Kingsnake{
		configs: make(map[string]*viper.Viper),
	}
}
