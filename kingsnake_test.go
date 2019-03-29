// git push -u origin master
package kingsnake_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	homedir "github.com/mitchellh/go-homedir"

	"github.com/runeimp/kingsnake"
	"github.com/spf13/viper"
)

func TestBasic(t *testing.T) {
	king := kingsnake.New()
	configHome := viper.New()
	configProject := viper.New()
	configRuneImp := viper.New()

	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	configHome.AddConfigPath(home)
	if path, ok := os.LookupEnv("XDG_DATA_HOME"); ok == true {
		xdgDataTK := fmt.Sprintf("%s/kingsnake", path)
		configHome.AddConfigPath(xdgDataTK)
		configHome.SetConfigName("config")
	}
	if path, ok := os.LookupEnv("XDG_CONFIG_HOME"); ok == true {
		xdgConfigTK := fmt.Sprintf("%s/kingsnake", path)
		configHome.AddConfigPath(xdgConfigTK)
		configHome.SetConfigName("config")
	}
	configHome.AutomaticEnv() // read in environment variables that match

	configProject.AddConfigPath("..")
	configProject.AddConfigPath(".")
	configProject.SetConfigName(".kingsnake")

	configRuneImp.SetConfigFile(".runeimp.yaml")

	king.Eat("project", configProject)
	king.Eat("runeimp", configRuneImp)
	king.Eat("home", configHome)

	if err := configProject.ReadInConfig(); err == nil {
		log.Println("configProject.Using config file:", configProject.ConfigFileUsed())
		printViper("configProject", "runeimp", configProject)
		printViper("configProject", "file_owner", configProject)
		printViper("configProject", "home_root", configProject)
		printViper("configProject", "project_owner", configProject)
		printViper("configProject", "xdg_config_home", configProject)
		log.Println()
	} else {
		log.Println(err)
		log.Println()
	}

	if err := configRuneImp.ReadInConfig(); err == nil {
		log.Println("configRuneImp.Using config file:", configRuneImp.ConfigFileUsed())
		printViper("configRuneImp", "runeimp", configRuneImp)
		printViper("configRuneImp", "file_owner", configRuneImp)
		printViper("configRuneImp", "home_root", configRuneImp)
		printViper("configRuneImp", "project_owner", configRuneImp)
		printViper("configRuneImp", "xdg_config_home", configRuneImp)
		log.Println()
	} else {
		log.Println(err)
		log.Println()
	}

	if err := configHome.ReadInConfig(); err == nil {
		log.Println("configHome.Using config file:", configHome.ConfigFileUsed())
		printViper("configHome", "runeimp", configHome)
		printViper("configHome", "file_owner", configHome)
		printViper("configHome", "home_root", configHome)
		printViper("configHome", "project_owner", configHome)
		printViper("configHome", "xdg_config_home", configHome)
		log.Println()
	} else {
		log.Println(err)
		log.Println()
	}

	log.Println("king Using config file:", king.ConfigFileUsed())
	log.Println("king.runeimp:", king.Get("runeimp"))
	log.Println("king.file_owner:", king.Get("file_owner"))
	log.Println("king.home_root:", king.Get("home_root"))
	log.Println("king.project_owner:", king.Get("project_owner"))
	log.Println("king.xdg_config_home:", king.Get("xdg_config_home"))
}

func printViper(prefix, key string, vipr *viper.Viper) {
	log.Printf("%s.%s: %#v (%T)", prefix, key, vipr.Get(key), vipr.Get(key))
}
