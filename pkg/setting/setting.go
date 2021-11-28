package setting

import (
	"github.com/spf13/viper"
	"time"
)

// Setting store all configuration of application.
// the value read by viper from a config file or environment variables
type Setting struct {
	DBDriver 			string `mapstructure:"DB_DRIVER"`
	DBSource 			string `mapstructure:"DB_SOURCE"`
	ServerAddress 		string `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey 	string `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`

	LogSavePath			string	`mapstructure:"LOG_SAVE_PATH"`
	LogSaveName			string	`mapstructure:"LOG_FILE_NAME"`
	LogFileExt			string	`mapstructure:"LOG_FILE_EXT"`
	TimeFormat			string	`mapstructure:"TIME_FORMAT"`

	DefaultPrefix		string	`mapstructure:"DEFAULT_PREFIX"`
	DefaultCallerDepth	int64	`mapstructure:"DEFAULT_CALLER_DEPTH"`
	LogPrefix			string	`mapstructure:"LOG_PREFIX"`
}

// LoadSetting read configuration from file or environment variables.
func LoadSetting(path string) (setting Setting, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	// viper.Set("levelFlags", []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"})

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&setting)
	return
}

