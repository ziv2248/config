package flag

import (
	"flag"
	"os"
	"testing"
)

const (
	ROLE_NONE  Role = 0
	ROLE_USER  Role = 1
	ROLE_ADMIN Role = 2
)

var (
	roleNames = map[int64]string{
		0: "None",
		1: "User",
		2: "Admin",
	}

	roleNameValues = map[string]Role{
		"None":  ROLE_NONE,
		"User":  ROLE_USER,
		"Admin": ROLE_ADMIN,
	}
)

type Role int64

func (r *Role) String() string {
	v, ok := roleNames[int64(*r)]
	if ok {
		return v
	}
	return "None"
}

func (r *Role) Set(name string) error {
	v, ok := roleNameValues[name]
	if ok {
		*r = v
	} else {
		*r = ROLE_NONE
	}
	return nil
}

type config struct {
	RedisHost   string `arg:"redis-host;the Redis server address and port"`
	RedisSecret string `arg:"redis-passowrd;the Redis password"`
	RedisDB     int    `arg:"redis-db;the Redis database number"`
	Workspace   string `arg:"*workspace;the data workspace"`
	Role        Role   `arg:"role;the role"`
}

func TestLoad(t *testing.T) {
	os.Args = []string{"example",
		"--redis-host", "192.168.56.53:6379",
		"--redis-passowrd", "foobared",
		"--redis-db", "3",
		"--role", "User",
	}

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	if flag.Parsed() {
		t.Errorf("assert flag.Parse():: expected '%v', got '%v'", false, flag.Parsed())
	}

	c := config{
		RedisSecret: "p@ssw0rd",
		Workspace:   "demo_test",
	}
	err := Process(&c)
	if err != nil {
		t.Error(err)
	}

	if c.RedisHost != "192.168.56.53:6379" {
		t.Errorf("assert 'config.RedisHost':: expected '%v', got '%v'", "192.168.56.53:6379", c.RedisHost)
	}
	if c.RedisSecret != "foobared" {
		t.Errorf("assert 'config.RedisSecret':: expected '%v', got '%v'", "foobared", c.RedisSecret)
	}
	if c.RedisDB != 3 {
		t.Errorf("assert 'config.RedisDB':: expected '%v', got '%v'", 3, c.RedisDB)
	}
	if c.Workspace != "demo_test" {
		t.Errorf("assert 'config.Workspace':: expected '%v', got '%v'", "demo_test", c.Workspace)
	}
	var expectedRole = ROLE_USER
	if c.Role != expectedRole {
		t.Errorf("assert 'config.Role':: expected '%v', got '%v'", expectedRole, c.Role)
	}
}
