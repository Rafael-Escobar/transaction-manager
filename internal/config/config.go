package config

import (
	"fmt"
	"log"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/spf13/viper"
)

const (
	DefaultConfigPath = "./internal/config/config.yml"
)

type DaemonOptions struct {
	ConfigPath string
	Verbose    bool
}

type RelationalDB struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	Name     string `mapstructure:"DB_NAME"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
}

func (db RelationalDB) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable connect_timeout=10",
		db.Host,
		db.Port,
		db.User,
		db.Name,
		db.Password)
}

type RelationalDBConnection struct {
	MaxIdleConns int           `mapstructure:"DB_MAX_IDLE_CONNS"`
	MaxOpenConns int           `mapstructure:"DB_MAX_OPEN_CONNS"`
	MaxLifeTime  time.Duration `mapstructure:"DB_MAX_LIFE_TIME"`
	MaxIdleTime  time.Duration `mapstructure:"DB_MAX_IDLE_TIME"`
}

type Config struct {
	Environment            Environment
	AppName                string                 `mapstructure:"APP_NAME"`
	Port                   string                 `mapstructure:"PORT"`
	LogLevel               string                 `mapstructure:"LOG_LEVEL"`
	RelationalDB           RelationalDB           `mapstructure:",squash"`
	RelationalDBConnection RelationalDBConnection `mapstructure:",squash"`
	Cors                   Cors                   `mapstructure:",squash"`
}

type Project struct {
	ID                        string `mapstructure:"PROJECT_ID"`
	IdentityPlatformProjectID string `mapstructure:"IDP_PROJECT_ID"`
}

// Cors contains the configuration for the CORS (Cross-Origin Resource Sharing) middleware.
type Cors struct {
	AllowHeaders     []string `mapstructure:"CORS_ALLOW_HEADERS"`
	AllowCredentials bool     `mapstructure:"CORS_ALLOW_CREDENTIALS"`
	AllowOrigins     []string `mapstructure:"CORS_ALLOW_ORIGINS"`
}

// ToGinConfiguration converts Cors to cors.Config.
//
// If AllowOrigins is nil or empty, AllowAllOrigins is enabled.
func (c Cors) ToGinConfiguration() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AddAllowHeaders(c.AllowHeaders...)
	corsConfig.AllowCredentials = c.AllowCredentials

	if len(c.AllowOrigins) == 0 {
		corsConfig.AllowAllOrigins = true
	}

	corsConfig.AllowOrigins = c.AllowOrigins

	return corsConfig
}

func Load(configPath string) (*Config, error) {
	initializeViper(viper.GetViper())
	err := loadViperConfig(viper.GetViper(), getConfigPath(configPath))
	if err != nil {
		return nil, err
	}

	viperEnv := initializeViperEnv(viper.GetString("Env"))

	config, err := buildConfig(viperEnv)
	if err != nil {
		return nil, err
	}

	config.log()

	return config, nil
}

func buildConfig(viperEnv *viper.Viper) (*Config, error) {
	var config *Config
	if err := viperEnv.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	env, err := Parse(viper.GetString("Env"))
	if err != nil {
		return nil, err
	}
	config.Environment = env

	return config, nil
}

func getConfigPath(configPath string) string {
	if configPath == "" {
		return DefaultConfigPath
	}
	return configPath
}

func bindEnv(v *viper.Viper, key ...string) {
	err := v.BindEnv(key...)
	if err != nil {
		log.Panic(err)
	}
}

func initializeViper(env *viper.Viper) {
	env.SetDefault("Env", "development")
	bindEnv(env, "ENV")
}

func initializeViperEnv(e string) *viper.Viper {
	viperEnv := viper.Sub(e)
	if viperEnv == nil {
		log.Panicf("config for environment '%v' doesn't exist", e)
	}

	viperEnv.SetDefault("PORT", "8080")
	viperEnv.SetDefault("DB_TYPE", "postgres")
	viperEnv.SetDefault("LOG_LEVEL", "INFO")

	bindEnvs(viperEnv)

	return viperEnv
}

func bindEnvs(viperEnv *viper.Viper) {
	bindEnv(viperEnv, "ENV")
	bindEnv(viperEnv, "DB_HOST")
	bindEnv(viperEnv, "DB_PROJECT_ID")
}

func loadViperConfig(env *viper.Viper, configPath string) error {
	extension := filepath.Ext(configPath)
	filename := filepath.Base(configPath)
	configName := strings.TrimSuffix(filename, extension)
	env.SetConfigName(configName)

	if len(extension) > 1 {
		configType := extension[1:]
		env.SetConfigType(configType)
	}

	configDir := filepath.Dir(configPath)
	env.AddConfigPath(configDir)

	err := env.ReadInConfig()
	if err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	return nil
}

func (c Config) log() {
	logStructure(c)
}

func logStructure(s interface{}) {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)
	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)
		fieldValue := v.Field(i)

		// Get the field tag value
		tag := field.Tag.Get("mapstructure")

		if tag == ",squash" {
			logStructure(fieldValue.Interface())
			continue
		}

		var value interface{}
		tagName := strings.ToLower(tag)
		isSecret := strings.Contains(tagName, "secret") || strings.Contains(tagName, "key") || strings.Contains(tagName, "password") || strings.Contains(tagName, "token") || strings.Contains(tagName, "tls")
		if isSecret {
			value = "***Secret***"
		} else {
			value = fieldValue.Interface()
		}

		if tag != "" {
			log.Printf("%v: '%v'", tag, value)
		}
	}
}
