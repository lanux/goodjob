package config

import (
	"errors"
	"flag"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
	"path/filepath"
)

var Global *appConfig

type appConfig struct {
	Host  *string      `toml:"server_host"`
	Port  *string      `toml:"server_port"`
	Cas   *casConfig   `toml:"cas"`
	Mysql *mysqlConfig `toml:"mysql"`
}

type casConfig struct {
	CasServerLoginUrl  string `toml:"cas_server_login_url"`
	ServerName         string `toml:"service_url"`
	CasServerUrlPrefix string `toml:"cas_server_url_prefix"`
}

type mysqlConfig struct {
	Url             string `toml:"url"`
	Username        string `toml:"username"`
	Password        string `toml:"password"`
	MaxOpenConns    int    `toml:"maxOpenConns"`
	ConnMaxLifetime int64  `toml:"connMaxLifetime"`
	MaxIdle         int    `toml:"maxIdle"`
}

var (
	h          bool
	profile    string
	configFile string
)

// -flag //只支持bool类型
// -flag=x
// -flag x //只支持非bool类型
// 以上语法对于一个或两个‘－’号，效果是一样的
func init() {
	flag.BoolVar(&h, "h", false, "show help")
	flag.BoolVar(&h, "help", false, "show help")
	flag.StringVar(&configFile, "F", "", "config file")
	flag.StringVar(&profile, "E", os.Getenv("active.profile"), "environment")
	flag.StringVar(&profile, "env", os.Getenv("active.profile"), "environment")
	Global, _ = loadConfig()
	validate()
}

func validate() {
	if Global == nil {
		panic(errors.New("config load error"))
	}
	if Global.Cas == nil {
		panic(errors.New("cas config load error"))
	}

}

func loadConfig() (*appConfig, error) {
	//调用flag.Parse()解析命令行参数到定义的flag
	//解析函数将会在碰到第一个非flag命令行参数时停止，非flag命令行参数是指不满足命令行语法的参数，如命令行参数为cmd --flag=true abc则第一个非flag命令行参数为“abc”
	flag.Parse()
	if h {
		flag.Usage()
		os.Exit(0)
	}
	if configFile == "" {
		if profile == "" {
			profile = "dev"
		}
		configFile = "./config/" + profile + ".toml"
	}
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return nil, errors.New("configuration file '" + configFile + "' does not exist")
	}
	c := appConfig{}

	// get the abs
	// which will try to find the 'filename' from current workind dir too.
	tomlAbsPath, err := filepath.Abs(configFile)
	if err != nil {
		return nil, err
	}
	// read the raw contents of the file
	data, err := ioutil.ReadFile(tomlAbsPath)
	if err != nil {
		return nil, err
	}
	// put the file's contents as toml to the default configuration(c)
	if _, err := toml.Decode(string(data), &c); err != nil {
		return nil, err
	}
	return &c, nil
}
