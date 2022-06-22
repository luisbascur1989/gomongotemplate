package conf

import (
	"io/ioutil"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// GlobalConfiguratorModel ...
type GlobalConfiguratorModel struct {
	LogLevel      int           `yaml:"logLevel"`
	Timeout       time.Duration `yaml:"timeout"`
	GzipResponse  bool          `yaml:"gzipResponse"`
	LimitBytesLog int           `yaml:"limitBytesLog"`
	Credentials   struct {
		Persist struct {
			MongoDB struct {
				User       string `yaml:"user"`
				Secret     string `yaml:"secret"`
				Database   string `yaml:"database"`
				Collection string `yaml:"collection"`
				Host       string `yaml:"host"`
			} `yaml:"mongoDB"`
		} `yaml:"persist"`
	} `yaml:"credentials"`
}

var (
	// GlobalConf ...
	GlobalConf = &GlobalConfiguratorModel{}
)

const (
	// ErrorJSONFORMAT ...
	ErrorJSONFORMAT = "file is in incorrect format"
)

func init() {
	urlconfig := os.Getenv("STACKER_CONFIG")
	if urlconfig == "" {
		pwd, _ := os.Getwd()
		urlconfig = pwd + "/configurator.local.yaml"
		log.Info("Using PWD os " + urlconfig)
	}

	file, err := ioutil.ReadFile(urlconfig)
	if err != nil {
		log.Fatal("Cannot read file")
	} else {
		if yaml.Unmarshal(file, &GlobalConf) != nil {
			println(ErrorJSONFORMAT)
			os.Exit(0)
		}
	}
}
