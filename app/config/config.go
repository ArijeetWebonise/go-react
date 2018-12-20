package config

import (
	"github.com/ArijeetBaruah/MyBlog/pkg/logger"
	"github.com/spf13/viper"
)

//Config defines behaviour for constructing app configuration
type Config interface {
	ConstructAppConfig() *AppConfig
}

//AppConfig defines properties required by application
type AppConfig struct {
	DB             DbConfig
	Logger         logger.ILogger
	Port           string
	CSRFAuthkey    string
	SessionAuthkey string
	GraphQL        GraphQLConfig
}

//GraphQLConfig config properties for GraphQl
type GraphQLConfig struct {
	Pretty     bool
	GraphiQL   bool
	Playground bool
}

// DbConfig wrapper for DB config
type DbConfig struct {
	DbUserName   string
	DbPassword   string
	DbHost       string
	DbName       string
	DbDriverName string
	DbDataSource string
}

//ConstructAppConfig prepares <AppConfig> from environment variables
func (appConfig *AppConfig) ConstructAppConfig() *AppConfig {
	viper.SetEnvPrefix("GR")
	viper.AutomaticEnv()

	appConfig.Port = appConfig.validateEnvVar("PORT")
	appConfig.CSRFAuthkey = appConfig.validateEnvVar("CSRF_AUTH_KEY")
	appConfig.SessionAuthkey = appConfig.validateEnvVar("SESSION_AUTH_KEY")

	appConfig.DB.DbUserName = appConfig.validateEnvVar("DB_USERNAME")
	appConfig.DB.DbPassword = appConfig.validateEnvVar("DB_PASSWORD")
	appConfig.DB.DbHost = appConfig.validateEnvVar("DB_HOST")
	appConfig.DB.DbName = appConfig.validateEnvVar("DB_NAME")
	appConfig.DB.DbDriverName = appConfig.validateEnvVar("DB_DRIVER_NAME")
	appConfig.DB.DbDataSource = appConfig.validateEnvVar("DB_DATA_SOURCE")
	appConfig.GraphQL.GraphiQL = viper.GetBool("GRAPHIQL")
	appConfig.GraphQL.Playground = viper.GetBool("GRAPHQL_PLAYGROUND")
	appConfig.GraphQL.Pretty = viper.GetBool("GRAPHQL_PRETTY")
	return appConfig
}

//validateEnvVar fetches environment variable value for a given <key> if present
//else panics
func (appConfig *AppConfig) validateEnvVar(key string) string {
	value := viper.GetString(key)
	if value == "" {
		appConfig.Logger.Panicf("%s is not set", key)
	}
	return value
}
