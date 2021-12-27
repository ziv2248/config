config
=========

## Synopsis

```go
package main

import (
	"fmt"

	"github.com/config"
)

type DummyConfig struct {
	RedisHost     string `env:"REDIS_HOST"       yaml:"redisHost"       arg:"redis-host;the Redis server address and port"`
	RedisPassword string `env:"REDIS_PASSWORD"   yaml:"redisPassword"   arg:"redis-passowrd;the Redis password"`
	RedisDB       int    `env:"REDIS_DB"         yaml:"redisDB"         arg:"redis-db;the Redis database number"`
	RedisPoolSize int    `env:"-"                yaml:"redisPoolSize"`
	Workspace     string `env:"-"                yaml:"workspace"       arg:"workspace;the data workspace"`
}

func main() {
	conf := DummyConfig{}

	config.NewConfigurationService(&conf).
		LoadEnvironmentVariables("").
		LoadEnvironmentVariables("K8S").
		LoadYamlFile("config.yaml").
		LoadYamlFile("config.${ENVIRONMENT}.yaml").
		LoadCommandArguments()

	fmt.Printf("%+v\n", conf)
}
```

```bash
export ENVIRONMENT=production
export REDIS_HOST=127.0.0.3:6379
export REDIS_PASSWORD=1234
export K8S_REDIS_HOST=demo-kubernetes:6379
export K8S_REDIS_PASSWORD=p@ssw0rd
export K8S_REDIS_DB=6
```


----------
## Syntax

#### Environment Variables
使用 **env** 標示，名稱可用字元為 `[A-Za-z0-9_-]`，可使用 `*` 置於名稱前綴表示必要欄位。
```go
type Config struct {
  RedisHost     string `env:"*REDIS_HOST"`
  RedisPassword string `env:"REDIS_PASSWORD"`
  RedisDB       int    `env:"REDIS_DB"`
}
```
下面的標示與前者作用相同，唯獨使用 `,` 分隔名稱與屬性。
```go
type Config struct {
  RedisHost     string `env:"REDIS_HOST,required"`
  RedisPassword string `env:"REDIS_PASSWORD"`
  RedisDB       int    `env:"REDIS_DB"`
}
```

#### Command Arguments
使用 **arg** 標示，名稱可用字元為 `[A-Za-z0-9_-]`，可使用 `*` 置於名稱前綴表示必要欄位；於 `;` 後可加上說明文字，可以使用 `--help` 參數列出該說明。

```go
type Config struct {
	RedisHost     string `arg:"redis-host;the Redis server address and port"`
	RedisPassword string `arg:"redis-passowrd;the Redis password"`
	RedisDB       int    `arg:"redis-db;the Redis database number"`
	Workspace     string `arg:"workspace;the data workspace"`
}
```

> **NOTE**: 不要使用 `help` 作為參數名稱。


----------
### Dependency
- Yaml - https://godoc.org/gopkg.in/yaml.v2
- Json - https://golang.org/pkg/encoding/json/
