package configs

import (
	"os"

	"gopkg.in/yaml.v2"
)

type ConfigType struct {
	Port    int
	Prefix  string
	TimeOut int `yaml:"timeOut"`
}

var Config ConfigType

const template = `
prefix: "/baseUrl"  # 路由前缀
port: 8080  # 启动端口
timeOut: 5  # 请求超时时间(s)
`

func init() {
	configFile := "./config.yml"
	data, err := os.ReadFile(configFile)
	if err != nil {
		os.Create(configFile)
		os.WriteFile(configFile, []byte(template), 0777)
		data, _ = os.ReadFile(configFile)
	}

	if err := yaml.Unmarshal([]byte(data), &Config); err != nil {
		panic(err)
	}
}