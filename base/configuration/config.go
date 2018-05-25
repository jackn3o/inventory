package configuration

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Default config keys for json format config
const (
	AppName         = "app.name"
	AppListenHost   = "app.host"
	AppListenPort   = "app.port"
	AppAPIBase      = "app.apiBase"
	AppReadTimeout  = "app.timeout.read"
	AppWriteTimeout = "app.timeout.write"

	DatabaseHost     = "database.connection.host"
	DatabaseUsername = "database.connection.username"
	DatabasePassword = "database.connection.password"
	DatabaseTimeout  = "database.connection.timeout"

	MainDatabaseName = "database.main.name"
	DemoDatabaseName = "database.demo.name"

	LogDir            = "log.dir"
	LogLevel          = "log.level"
	LogEnableRotation = "log.enableRotation"

	MasterDatabaseName = "database.master.name"

	SecretKey = "secretKey"
)

// Config is duplicate of viper interface
type Config interface {
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	GetDecodedString(key string) []byte
	InConfig(key string) bool
	IsSet(key string) bool
	Set(key string, value interface{})
	SetDefault(key string, value interface{})
	SetEnvPrefix(in string)
	Unmarshal(s interface{}) error
	Dump(w io.Writer)
	FlagBool(name string, value bool, usage string)
	FlagInt(name string, value int, usage string)
	FlagString(name string, value string, usage string)
	Update()
}

type appConfig struct {
	v   *viper.Viper
	cmd *cobra.Command
}

func (cfg *appConfig) Get(key string) interface{} {
	return cfg.v.Get(key)
}

func (cfg *appConfig) GetBool(key string) bool {
	return cfg.v.GetBool(key)
}

func (cfg *appConfig) GetDuration(key string) time.Duration {
	return cfg.v.GetDuration(key)
}

func (cfg *appConfig) GetFloat64(key string) float64 {
	return cfg.v.GetFloat64(key)
}

func (cfg *appConfig) GetInt(key string) int {
	return cfg.v.GetInt(key)
}

func (cfg *appConfig) GetInt64(key string) int64 {
	return cfg.v.GetInt64(key)
}

func (cfg *appConfig) GetSizeInBytes(key string) uint {
	return cfg.v.GetSizeInBytes(key)
}

func (cfg *appConfig) GetString(key string) string {
	return cfg.v.GetString(key)
}

func (cfg *appConfig) GetStringMap(key string) map[string]interface{} {
	return cfg.v.GetStringMap(key)
}

func (cfg *appConfig) GetStringMapString(key string) map[string]string {
	return cfg.v.GetStringMapString(key)
}

func (cfg *appConfig) GetStringMapStringSlice(key string) map[string][]string {
	return cfg.v.GetStringMapStringSlice(key)
}

func (cfg *appConfig) GetStringSlice(key string) []string {
	return cfg.v.GetStringSlice(key)
}

func (cfg *appConfig) GetTime(key string) time.Time {
	return cfg.v.GetTime(key)
}

func (cfg *appConfig) GetDecodedString(key string) []byte {
	decoded, err := hex.DecodeString(cfg.GetString(key))
	if err != nil {
		return nil
	}

	return decoded
}

func (cfg *appConfig) InConfig(key string) bool {
	return cfg.v.InConfig(key)
}

func (cfg *appConfig) IsSet(key string) bool {
	return cfg.v.IsSet(key)
}

func (cfg *appConfig) Set(key string, value interface{}) {
	cfg.v.Set(key, value)
}

func (cfg *appConfig) SetDefault(key string, value interface{}) {
	cfg.v.SetDefault(key, value)
}

func (cfg *appConfig) SetEnvPrefix(in string) {
	cfg.v.SetEnvPrefix(in)
}

func (cfg *appConfig) Unmarshal(s interface{}) error {
	return cfg.v.Unmarshal(s)
}

func (cfg *appConfig) Dump(w io.Writer) {
	for k, v := range cfg.v.AllSettings() {
		w.Write([]byte(fmt.Sprintf("%s: %v\n", k, v)))
	}
}

func (cfg *appConfig) FlagBool(name string, value bool, usage string) {
	cfg.cmd.Flags().Bool(name, value, usage)
	cfg.v.BindPFlag(name, cfg.cmd.Flags().Lookup(name))
}

func (cfg *appConfig) FlagString(name string, value string, usage string) {
	cfg.cmd.Flags().String(name, value, usage)
	cfg.v.BindPFlag(name, cfg.cmd.Flags().Lookup(name))
}

func (cfg *appConfig) FlagInt(name string, value int, usage string) {
	cfg.cmd.Flags().Int(name, value, usage)
	cfg.v.BindPFlag(name, cfg.cmd.Flags().Lookup(name))
}

func (cfg *appConfig) Update() {
	cfg.cmd.ParseFlags(os.Args[1:])
}

// New create and return configuration
func New(appName string, cmd *cobra.Command, configDir ...string) (Config, error) {

	v := viper.New()
	v.SetDefault(AppListenPort, ":3000")
	// v.SetDefault(OptSSLCertificate, "")
	// v.SetDefault(OptSSLCertificateKey, "")
	// v.SetDefault(OptLogEnableRotation, true)
	// v.SetDefault(OptMode, "debug")

	v.SetConfigType("json")
	v.SetConfigName(appName)

	for _, dir := range configDir {
		v.AddConfigPath(dir)
	}

	v.AddConfigPath("/etc/" + appName)
	v.AddConfigPath("$HOME/." + appName)
	v.AddConfigPath(".")

	// make it overridable from environment variables
	v.AutomaticEnv()
	cfg := &appConfig{v, cmd}

	// set default flag
	cfg.FlagString(AppListenHost, "localhost:3000", "App host url")
	cfg.FlagString(AppListenPort, ":3000", "Port to run http server on")
	cfg.FlagString(AppName, appName, "Application Name")
	// cfg.FlagString(OptSSLCertificate, "", "SSL Certificate")
	// cfg.FlagString(OptSSLCertificateKey, "", "SSL Certificate Key")
	// cfg.FlagString(OptMode, "debug", "Runtime mode (release, debug)")
	cfg.FlagString(AppAPIBase, "/api/v1", "API base")
	cfg.FlagString(SecretKey, "secret", "Security secret key")

	// logger
	// cfg.FlagString(OptLogFileName, "", "Logging output file name")
	cfg.FlagString(LogDir, "/eventlog", "Logging output dir")
	cfg.FlagInt(LogLevel, 0, "Logging level output") // only log panic panic
	cfg.FlagBool(LogEnableRotation, true, "Enable logging file auto rotation")
	// cfg.FlagBool(OptEnableSyslog, false, "Enable logging to syslog")
	// cfg.FlagString(OptSyslogNetwork, "", "Syslog network (tcp or udp)")
	// cfg.FlagString(OptSyslogAddress, "", "Syslog address, use empty string for local syslog")
	// cfg.FlagString(OptSyslogFacility, "", "Syslog facility, default is KERN")

	// cfg.FlagString(OptTemplateDir, "./webapp/dist", "Template directory")
	// cfg.FlagString(OptStaticDir, "./webapp/dist", "Static directory")
	// cfg.FlagString(OptStaticBase, "/static", "Static path")
	// cfg.FlagInt(OptStaticCacheExpire, -1, "Static content cache expire")
	// cfg.FlagString(OptUploadDir, "./uploads", "Upload directory")
	// cfg.FlagString(OptDocumentBase, "/uploads/documents/", "document base URL")

	// // security secret & encrypt
	// cfg.FlagString(OptAuthKey, "97d6867beb4d4726c5958e03e9337b4599ae3e43a49433a07eea569eb473fcfb", "Security auth key")
	// cfg.FlagString(OptEncryptKey, "97d6867beb4d4726c5958e03e9337b4599ae3e43a49433a07eea569eb473fcfb", "Security encrypt key")
	// cfg.FlagInt(OptSessionExpiry, 10*60, "Security session timeout")
	// cfg.FlagInt(OptTokenExpiry, 2*60*60, "Seconds before token expired")
	// cfg.FlagString(OptTokenName, "auth", "Token Name")
	// cfg.FlagString(OptAuthCookieName, "uid", "Auth Cookie Name")
	// cfg.FlagString(OptSessionCookieName, "sid", "Session Cookie Name")
	// cfg.FlagString(OptPushChanel, "events", "Push notification channel")

	// database and key value store
	cfg.FlagString(DatabaseHost, "localhost:27017", "sql (mongo) url")
	cfg.FlagInt(DatabaseTimeout, 60, "mongo timeout in second")

	// cfg.FlagString(OptKvURI, "localhost:6379", "KV (redis) url")
	// cfg.FlagString(OptKvPassword, "", "KV (redis) password")
	// cfg.FlagInt(OptKvMaxOpen, 10, "KV (redis) max open")
	// cfg.FlagInt(OptKvMaxIdle, 8, "KV (redis) max idle")
	// cfg.FlagInt(OptKvIdleTimeout, 10, "KV (redis) idle timeout")
	// cfg.FlagBool(OptKvWait, false, "Wait for available kv conn")

	// cfg.FlagInt(OptPageLimit, 10, "Maximum data page size")

	// cfg.FlagString(OptDefaultPassword, "12345678", "Default password for non self-registering user")

	if err := v.ReadInConfig(); err != nil {
		return cfg, err
	}

	return cfg, nil
}
