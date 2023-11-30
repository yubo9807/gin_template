package configs

import (
	"os"

	"gopkg.in/yaml.v2"
)

type ConfigType struct {
	Port   int
	Prefix string

	TokenValidTime         int64 `yaml:"tokenValidTime"`
	TokenExceedRefreshTime int64 `yaml:"tokenExceedRefreshTime"`
}

var Config ConfigType

const template = `
prefix: "/baseUrl"  # 路由前缀
port: 8080  # 启动端口

tokenValidTime: 7200  # 令牌有效时间(s)
TokenExceedRefreshTime: 86400  # 令牌超过刷新时间(s)
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
