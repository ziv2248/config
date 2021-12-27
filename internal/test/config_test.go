package test

import (
	"flag"
	"os"
	"testing"

	. "github.com/config"

	"gopkg.in/yaml.v2"
)

type DummyConfig struct {
	RedisHost     string `env:"REDIS_HOST"       yaml:"redisHost"       arg:"redis-host;the Redis server address and port"`
	RedisPassword string `env:"REDIS_PASSWORD"   yaml:"redisPassword"   arg:"redis-passowrd;the Redis password"`
	RedisDB       int    `env:"REDIS_DB"         yaml:"redisDB"         arg:"redis-db;the Redis database number"`
	RedisPoolSize int    `env:"-"                yaml:"redisPoolSize"`
	Workspace     string `env:"-"                yaml:"workspace"       arg:"workspace;the data workspace"`
	Version       string `resource:".VERSION"`
}

func TestConfigurationService(t *testing.T) {
	os.Clearenv()
	initializeEnvironment()
	initializekubernetesEnvironment()
	initializeArgs()

	conf := DummyConfig{}

	NewConfigurationService(&conf).
		LoadEnvironmentVariables("").
		LoadEnvironmentVariables("K8S").
		LoadYamlFile("config.yaml").
		LoadYamlFile("config.${ENVIRONMENT}.yaml").
		LoadCommandArguments().
		LoadResource("")

	var expectedRedisHost = "demo-kubernetes:6379"
	if conf.RedisHost != expectedRedisHost {
		t.Errorf("assert 'config.RedisHost':: expected '%v', got '%v'", expectedRedisHost, conf.RedisHost)
	}
	var expectedRedisPassword = "p@ssw0rd"
	if conf.RedisPassword != expectedRedisPassword {
		t.Errorf("assert 'config.RedisPassword':: expected '%v', got '%v'", expectedRedisPassword, conf.RedisPassword)
	}
	var expectedRedisDB = 32
	if conf.RedisDB != expectedRedisDB {
		t.Errorf("assert 'config.RedisDB':: expected '%v', got '%v'", expectedRedisDB, conf.RedisDB)
	}
	var expectedRedisPoolSize = 50
	if conf.RedisPoolSize != expectedRedisPoolSize {
		t.Errorf("assert 'config.RedisPoolSize':: expected '%v', got '%v'", expectedRedisPoolSize, conf.RedisPoolSize)
	}
	var expectedWorkspace = "demo_prod"
	if conf.Workspace != expectedWorkspace {
		t.Errorf("assert 'config.Workspace':: expected '%v', got '%v'", expectedWorkspace, conf.Workspace)
	}
	var expectedVersion = "v1.0.2"
	if conf.Version != expectedVersion {
		t.Errorf("assert 'config.Version':: expected '%v', got '%v'", expectedVersion, conf.Version)
	}
}

func initializeEnvironment() {
	os.Setenv("ENVIRONMENT", "staging")
	os.Setenv("REDIS_HOST", "127.0.0.3:6379")
	os.Setenv("REDIS_PASSWORD", "1234")
}

func initializekubernetesEnvironment() {
	os.Setenv("ENVIRONMENT", "production")
	os.Setenv("K8S_REDIS_HOST", "demo-kubernetes:6379")
	os.Setenv("K8S_REDIS_PASSWORD", "p@ssw0rd")
	os.Setenv("K8S_REDIS_DB", "6")
}

func initializeArgs() {
	os.Args = []string{"example",
		"--redis-db", "32"}

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}

func TestConfigurationService_WithUnmarshalFunc(t *testing.T) {

	conf := DummyConfig{}

	NewConfigurationService(&conf).
		LoadFile("config.yaml", yaml.Unmarshal)

	var expectedRedisHost = ""
	if conf.RedisHost != expectedRedisHost {
		t.Errorf("assert 'config.RedisHost':: expected '%v', got '%v'", expectedRedisHost, conf.RedisHost)
	}
	var expectedRedisPassword = ""
	if conf.RedisPassword != expectedRedisPassword {
		t.Errorf("assert 'config.RedisPassword':: expected '%v', got '%v'", expectedRedisPassword, conf.RedisPassword)
	}
	var expectedRedisDB = 3
	if conf.RedisDB != expectedRedisDB {
		t.Errorf("assert 'config.RedisDB':: expected '%v', got '%v'", expectedRedisDB, conf.RedisDB)
	}
	var expectedRedisPoolSize = 10
	if conf.RedisPoolSize != expectedRedisPoolSize {
		t.Errorf("assert 'config.RedisPoolSize':: expected '%v', got '%v'", expectedRedisPoolSize, conf.RedisPoolSize)
	}
	var expectedWorkspace = "demo_test"
	if conf.Workspace != expectedWorkspace {
		t.Errorf("assert 'config.Workspace':: expected '%v', got '%v'", expectedWorkspace, conf.Workspace)
	}
	var expectedVersion = ""
	if conf.Version != expectedVersion {
		t.Errorf("assert 'config.Version':: expected '%v', got '%v'", expectedVersion, conf.Version)
	}
}
