package config

import (
	"encoding/json"
	"io/ioutil"
	"sync"

	"github.com/sirupsen/logrus"
)

var once sync.Once
var ConfigInstance *Config

func (c *Config) Init(filename string) bool {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil || len(bytes) == 0 {
		logrus.Errorf("read configfile err: %s", err)
		return false
	}
	err = json.Unmarshal(bytes, c)
	if err != nil {
		logrus.Errorln("Uncode Json err:", err)
		// fixit
	}
	return true
}

func GetConfig() *Config {
	once.Do(func() {
		ConfigInstance = new(Config)
		ConfigInstance.Init("config/config.json")
	})
	// fixit ConfigInstance ==
	return ConfigInstance
}
