package resource

import (
	"reflect"
	"testing"
)

type config struct {
	ResString string `resource:".RES_STR"`
	ResInt    int    `resource:".RES_INT"`
	ResBytes  []byte `resource:".RES_BIN"`
}

func TestLoad_WithFromCurrentDir(t *testing.T) {
	c := config{}

	err := Process(".", &c)
	if err != nil {
		t.Error(err)
	}

	var expectedResString = "192.168.56.53:6379"
	if c.ResString != expectedResString {
		t.Errorf("assert 'config.ResString':: expected '%v', got '%v'", expectedResString, c.ResString)
	}
	var expectedResInt = 3
	if c.ResInt != expectedResInt {
		t.Errorf("assert 'config.ResInt':: expected '%v', got '%v'", expectedResInt, c.ResInt)
	}
	var expectedResBytes = []byte("the quick brown fox jumps over the lazy dog")
	if !reflect.DeepEqual(c.ResBytes, expectedResBytes) {
		t.Errorf("assert 'config.ResBytes':: expected '%v', got '%v'", expectedResBytes, c.ResBytes)
	}
}

func TestLoad_WithEmptyStringBaseDir(t *testing.T) {
	c := config{}

	err := Process("", &c)
	if err != nil {
		t.Error(err)
	}

	var expectedResString = "192.168.56.53:6379"
	if c.ResString != expectedResString {
		t.Errorf("assert 'config.ResString':: expected '%v', got '%v'", expectedResString, c.ResString)
	}
	var expectedResInt = 3
	if c.ResInt != expectedResInt {
		t.Errorf("assert 'config.ResInt':: expected '%v', got '%v'", expectedResInt, c.ResInt)
	}
	var expectedResBytes = []byte("the quick brown fox jumps over the lazy dog")
	if !reflect.DeepEqual(c.ResBytes, expectedResBytes) {
		t.Errorf("assert 'config.ResBytes':: expected '%v', got '%v'", expectedResBytes, c.ResBytes)
	}
}

func TestLoad_WithFromSpecifiedDir(t *testing.T) {
	c := config{}

	err := Process("conf", &c)
	if err != nil {
		t.Error(err)
	}

	var expectedResString = "192.168.56.53:9200"
	if c.ResString != expectedResString {
		t.Errorf("assert 'config.ResString':: expected '%v', got '%v'", expectedResString, c.ResString)
	}
	var expectedResInt = 6
	if c.ResInt != expectedResInt {
		t.Errorf("assert 'config.ResInt':: expected '%v', got '%v'", expectedResInt, c.ResInt)
	}
	var expectedResBytes = []byte("the quick brown fox jumps over the lazy dog")
	if !reflect.DeepEqual(c.ResBytes, expectedResBytes) {
		t.Errorf("assert 'config.ResBytes':: expected '%v', got '%v'", expectedResBytes, c.ResBytes)
	}
}

func TestLoad_WithFromEnvironmentVar(t *testing.T) {
	c := config{}

	err := Process("${Environment}", &c)
	if err != nil {
		t.Error(err)
	}

	var expectedResString = "192.168.56.112:6379"
	if c.ResString != expectedResString {
		t.Errorf("assert 'config.ResString':: expected '%v', got '%v'", expectedResString, c.ResString)
	}
	var expectedResInt = 2
	if c.ResInt != expectedResInt {
		t.Errorf("assert 'config.ResInt':: expected '%v', got '%v'", expectedResInt, c.ResInt)
	}
	var expectedResBytes = []byte("the quick brown fox jumps over the lazy dog")
	if !reflect.DeepEqual(c.ResBytes, expectedResBytes) {
		t.Errorf("assert 'config.ResBytes':: expected '%v', got '%v'", expectedResBytes, c.ResBytes)
	}
}
