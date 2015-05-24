/* Inspired by tj's code: https://github.com/tj/robo/blob/master/config/config.go  */

package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	cfgdata = make(map[string]string)
)

func init() {

	// Use development as default
	env := os.Getenv("GO_BLOG_ENV")
	if env == "" {
		env = "development"
	}

	b, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal(b, &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	data, ok := m[env]
	if !ok {
		log.Fatalf("invalid env: %v", env)
	}

	for k, v := range data.(map[interface{}]interface{}) {
		cfgdata[k.(string)] = fmt.Sprintf("%v", v)
	}
}

func Get(key string) string {
	return cfgdata[key]
}
