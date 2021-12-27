package env

import (
	"os"
	"testing"
)

type config struct {
	RedisHost    string `env:"*REDIS_HOST"`
	RedisSecret  string `env:"RESID_SECRET"`
	RedisDB      int    `env:"REDIS_DB"`
	Workspace    string `env:"*WORKSPACE"`
	IgnoredField string `env:"-"`
}

func TestLoad(t *testing.T) {
	os.Setenv("REDIS_HOST", "192.168.56.53")
	os.Setenv("RESID_SECRET", "foobar")
	os.Setenv("REDIS_DB", "3")
	os.Setenv("WORKSPACE", "demo_test")

	c := config{}
	err := Process("", &c)
	if err != nil {
		t.Error(err)
	}

	if c.RedisHost != "192.168.56.53" {
		t.Errorf("assert 'config.RedisHost':: expected '%v', got '%v'", "192.168.56.53", c.RedisHost)
	}
	if c.RedisSecret != "foobar" {
		t.Errorf("assert 'config.RedisSecret':: expected '%v', got '%v'", "foobar", c.RedisSecret)
	}
	if c.RedisDB != 3 {
		t.Errorf("assert 'config.RedisDB':: expected '%v', got '%v'", 3, c.RedisDB)
	}
	if c.Workspace != "demo_test" {
		t.Errorf("assert 'config.Workspace':: expected '%v', got '%v'", "demo_test", c.Workspace)
	}
	if c.IgnoredField != "" {
		t.Errorf("assert 'config.IgnoredField':: expected '%v', got '%v'", "", c.IgnoredField)
	}
}

func TestLoadWithPrefix(t *testing.T) {
	os.Setenv("K8S_REDIS_HOST", "192.168.56.53")
	os.Setenv("K8S_RESID_SECRET", "foobar")
	os.Setenv("K8S_REDIS_DB", "3")
	os.Setenv("K8S_WORKSPACE", "demo_test")

	c := config{}
	err := Process("K8S", &c)
	if err != nil {
		t.Error(err)
	}

	if c.RedisHost != "192.168.56.53" {
		t.Errorf("assert 'config.RedisHost':: expected '%v', got '%v'", "192.168.56.53", c.RedisHost)
	}
	if c.RedisSecret != "foobar" {
		t.Errorf("assert 'config.RedisSecret':: expected '%v', got '%v'", "foobar", c.RedisSecret)
	}
	if c.RedisDB != 3 {
		t.Errorf("assert 'config.RedisDB':: expected '%v', got '%v'", 3, c.RedisDB)
	}
	if c.Workspace != "demo_test" {
		t.Errorf("assert 'config.Workspace':: expected '%v', got '%v'", "demo_test", c.Workspace)
	}
}
