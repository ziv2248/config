package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/config/internal/env"
	"github.com/config/internal/flag"
	"github.com/config/internal/json"
	"github.com/config/internal/resource"
	"github.com/config/internal/yaml"
)

type ConfigurationService struct {
	target interface{}
}

func NewConfigurationService(target interface{}) *ConfigurationService {
	instance := ConfigurationService{
		target: target,
	}
	return &instance
}

func (service *ConfigurationService) LoadEnvironmentVariables(prefix string) *ConfigurationService {
	err := env.Process(prefix, service.target)
	if err != nil {
		panic(fmt.Errorf("config: %v\n", err))
	}
	return service
}

func (service *ConfigurationService) LoadCommandArguments() *ConfigurationService {
	err := flag.Process(service.target)
	if err != nil {
		panic(fmt.Errorf("config: %v\n", err))
	}
	return service
}

func (service *ConfigurationService) LoadJsonFile(filepath string) *ConfigurationService {
	err := json.LoadFile(filepath, service.target)
	if err != nil && os.IsExist(err) {
		panic(fmt.Errorf("config: %v\n", err))
	}
	return service
}

func (service *ConfigurationService) LoadJsonBytes(buffer []byte) *ConfigurationService {
	err := json.LoadBytes(buffer, service.target)
	if err != nil {
		panic(fmt.Errorf("config: %v\n", err))
	}
	return service
}

func (service *ConfigurationService) LoadYamlFile(filepath string) *ConfigurationService {
	err := yaml.LoadFile(filepath, service.target)
	if err != nil && os.IsExist(err) {
		panic(fmt.Errorf("config: %#v\n", err))
	}
	return service
}

func (service *ConfigurationService) LoadYamlBytes(buffer []byte) *ConfigurationService {
	err := yaml.LoadBytes(buffer, service.target)
	if err != nil {
		panic(fmt.Errorf("config: %v\n", err))
	}
	return service
}

func (service *ConfigurationService) LoadResource(baseDir string) *ConfigurationService {
	err := resource.Process(baseDir, service.target)
	if err != nil {
		panic(fmt.Errorf("config: %v\n", err))
	}
	return service
}

func (service *ConfigurationService) LoadFile(fullpath string, unmarshal Unmarshal) *ConfigurationService {
	path := os.ExpandEnv(fullpath)
	buffer, err := ioutil.ReadFile(path)
	if err != nil && os.IsExist(err) {
		panic(fmt.Errorf("config: %#v\n", err))
	}

	err = unmarshal(buffer, service.target)
	if err != nil {
		panic(fmt.Errorf("config: %#v\n", err))
	}
	return nil
}

func (service *ConfigurationService) LoadBytes(buffer []byte, unmarshal Unmarshal) *ConfigurationService {
	err := unmarshal(buffer, service.target)
	if err != nil {
		panic(fmt.Errorf("config: %#v\n", err))
	}
	return nil
}
