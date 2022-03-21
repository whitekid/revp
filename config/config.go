package config

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/whitekid/go-utils/flags"
)

var (
	Server *ServerConfig
	Client *ClientConfig

	envReplacer *strings.Replacer
)

func init() {
	Server = &ServerConfig{viper: viper.New()}
	Client = &ClientConfig{viper: viper.New()}

	envReplacer = strings.NewReplacer("-", "_")
	initViper(viper.GetViper())
	initViper(Client.viper)
	initViper(Server.viper)
}

func initViper(v *viper.Viper) {
	v.SetEnvPrefix("rp")
	v.SetEnvKeyReplacer(envReplacer)
	v.AutomaticEnv()
}

// global configs
const (
	keySecret = "secret"
)

func Secret() string { return viper.GetString(keySecret) }

func Init(use string, flagset *pflag.FlagSet) {
	configs := map[string][]flags.Flag{
		"revp": {
			{keySecret, "", "", "secret key"},
		},
	}

	flags.InitDefaults(nil, configs)
	flags.InitFlagSet(nil, configs, use, flagset)
}

// ClientConfig is configurations for client
type ClientConfig struct {
	viper *viper.Viper
}

const (
	keyServerAddr             = "server-addr"
	keyClientKeepaliveTime    = "keepalive"
	keyClientKeepaliveTimeout = "keepalive-timeout"
)

func (client *ClientConfig) Init(use string, flagset *pflag.FlagSet) {
	clientConfigs := map[string][]flags.Flag{
		"revp": {
			{keyServerAddr, "s", "revp.woosum.net:49999", "server address"},
			{keyClientKeepaliveTime, "", 10 * time.Second, "keepalive time"},
			{keyClientKeepaliveTimeout, "", 1 * time.Second, "keepalive timeout"},
		},
	}

	flags.InitDefaults(client.viper, clientConfigs)
	flags.InitFlagSet(client.viper, clientConfigs, use, flagset)
}

func (client *ClientConfig) ServerAddr() string {
	return client.viper.GetString(keyServerAddr)
}

func (client *ClientConfig) KeepaliveTime() time.Duration {
	return client.viper.GetDuration(keyClientKeepaliveTime)
}

func (client *ClientConfig) KeepaliveTimeout() time.Duration {
	return client.viper.GetDuration(keyClientKeepaliveTimeout)
}

// ServerConfig is configurations for server
type ServerConfig struct {
	viper *viper.Viper
}

// server
const (
	keyRootURL                = "root-url"
	keyBindAddr               = "bind-addr"
	keyPortRange              = "port-range"
	keyServerKeepaliveTime    = "keepalive"
	keyServerKeepaliveTimeout = "keepalive-timeout"
	keyServerDemoSecret       = "demo"
	keyServerDemoTimeout      = "demo-timeout"
)

func (server *ServerConfig) Init(use string, flagset *pflag.FlagSet) {
	serverConfigs := map[string][]flags.Flag{
		"revps": {
			{keyRootURL, "", "https://revp.woosum.net", "remote server root url without ports"},
			{keyBindAddr, "b", ":49999", "server bind address"},
			{keyPortRange, "", "50000:59999", "server port range"},
			{keyServerKeepaliveTime, "", 5 * time.Second, "keepalive time"},
			{keyServerKeepaliveTimeout, "", 1 * time.Second, "keepalive timeout"},
			{keyServerDemoSecret, "", "demo", "demo secret"},
			{keyServerDemoTimeout, "", 5 * time.Minute, "timeout for demo mode"},
		},
	}

	flags.InitDefaults(server.viper, serverConfigs)
	flags.InitFlagSet(server.viper, serverConfigs, use, flagset)
}

func (server *ServerConfig) RootURL() string {
	return server.viper.GetString(keyRootURL)
}

func (server *ServerConfig) BindAddr() string {
	return server.viper.GetString(keyBindAddr)
}

func (server *ServerConfig) KeepaliveTime() time.Duration {
	return server.viper.GetDuration(keyServerKeepaliveTime)
}

func (server *ServerConfig) KeepaliveTimeout() time.Duration {
	return server.viper.GetDuration(keyServerKeepaliveTimeout)
}

func (server *ServerConfig) PortRange() []int {
	r := server.viper.GetString(keyPortRange)
	portRange := strings.Split(r, ":")
	if len(portRange) != 2 {
		log.Fatalf("invalid port range: %s", portRange)
	}

	intRange := make([]int, len(portRange))
	for i := 0; i < len(portRange); i++ {
		p, err := strconv.ParseInt(portRange[i], 10, 56)
		if err != nil {
			log.Fatal(err)
		}

		intRange[i] = int(p)
	}

	return intRange
}

func (server *ServerConfig) DemoSecret() string {
	return server.viper.GetString(keyServerDemoSecret)
}

func (server *ServerConfig) DemoTimeout() time.Duration {
	return server.viper.GetDuration(keyServerDemoTimeout)
}
