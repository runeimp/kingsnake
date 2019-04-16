Kingsnake v0.3.0
================

King snakes eat other snakes. This kingsnake eats viper configs!

Library to allow the use of multiple [Viper][] instances and checking for values in each consumed viper until a match is found or the viper's den has been exhausted.

Handy when you want your project config values to add to or overwrite your home directory config values and those test config values to have presidence over all when specified. Or something along those lines.

Kingsnake currently supports Viper method propogation for ConfigFileUsed, Get, GetBool, GetFloat64 GetInt, GetString, and IsSet.


Example
-------

```go
import (
	"github.com/runeimp/kingsnake"
)

// rootCmd: Yes Cobra rocks!
var rootCmd = &cobra.Command{
	Use:     "queenie",
	Short:   "QueenSnake"
	Long:    `Everything is QueenSnake!
Everything is cool when your eating other snakes.`,
	Version: "0.1.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	configFile    string
	configCustom  *viper.Viper
	configHome    *viper.Viper
	configProject *viper.Viper
	kingSnake     kingsnake.Kingsnake
)

func init() {
	// log.Println("cmd/root.init()")
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "Custom config file. Takes presidence over ./.queensnake.yaml (or ../.queensnake.yaml), and $HOME/.queensnake.yaml)")
}

func initConfig() {
	configCustom = viper.New()
	configHome = viper.New()
	configProject = viper.New()
	if configFile != "" {
		// Use config file from the flag.
		configCustom.SetConfigFile(configFile)
	}

	configProject.AddConfigPath("..")
	configProject.AddConfigPath(".")
	configProject.SetConfigName(".queensnake")

	// Search config in home directory with name ".queensnake" (without extension).
	if path, ok := os.LookupEnv("HOME"); ok == true {
		configHome.AddConfigPath(path)
		configHome.SetConfigName(".queensnake")
	}
	if path, ok := os.LookupEnv("XDG_DATA_HOME"); ok == true {
		configHome.AddConfigPath(path + "/queensnake")
		configHome.SetConfigName("config")
	}
	if path, ok := os.LookupEnv("XDG_CONFIG_HOME"); ok == true {
		configHome.AddConfigPath(path + "/queensnake")
		configHome.SetConfigName("config")
	}
	configHome.SetConfigName(".queensnake")

	if err := configCustom.ReadInConfig(); err != nil {
		log.Println("configCustom", err)
	}
	if err := configProject.ReadInConfig(); err != nil {
		log.Println("configProject", err)
	}
	if err := configHome.ReadInConfig(); err != nil {
		log.Println("configconfigHomeCustom", err)
	}

	kingSnake = kingsnake.New()
	kingSnake.Eat("custom", configCustom)
	kingSnake.Eat("project", configProject)
	kingSnake.Eat("home", configHome)

	fmt.Println(kingSnake.GetString("project.sharable.data"))
	// kingSnake will check each config for the key
	// "project.sharable.data" in the order eaten
	// Thus the project config will have higher precedence than the home config.
	// And the custom config (if used) will have the highest precedence of all.
}
```



[Viper]: https://github.com/spf13/viper



