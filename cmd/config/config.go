package config

import (
	"os"
	"sync"
	"gopkg.in/yaml.v3"
)

type CSHConfig struct {
	Connection ConnectionConfig `yaml:"connection"`
	Rest RestServer `yaml:"rest_server"`
}

type ConnectionConfig struct {
	CQL CQLConfig `yaml:"cql"`
}

type CQLConfig struct {
	AgentConfig Agent `yaml:"agent"`
}

type Agent struct {
	InstanceName string `yaml:"instance_name"`
	Type string `yaml:"type"`
	LocalDatacenter string `yaml:"local_datacenter"`
	ContactPoints []ContactPoint `yaml:"contact_points"`
	DCAware DatacenterAware `yaml:"datacenter_aware"`
}

type ContactPoint struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type DatacenterAware struct {
	Datacenters []Datacenter `yaml:"datacenters"`
}

type Datacenter struct {
	Name string `yaml:"name"`
}

type RestServer struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
}

var (
	configInstance *CSHConfig
	once           sync.Once
)

func LoadConfig(filePath string) (*CSHConfig, error) {
	var err error
	once.Do(func() {
		var cfg CSHConfig
		data, err := os.ReadFile(filePath)
		if err != nil {
			return
		}
		if err = yaml.Unmarshal(data, &cfg); err != nil {
			return
		}
		configInstance = &cfg
	})
	return configInstance, err
}

func GetConfig() *CSHConfig {
	return configInstance
}
