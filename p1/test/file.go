package test

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Define struct for Hosts
type Host struct {
	PeerIP     string `yaml:"peer_ip"`
	ExternalIP string `yaml:"external_ip"`
	Port       int    `yaml:"port"`
	SSHCommand string `yaml:"ssh_command"`
}

// Define struct for S3
type S3 struct {
	Enabled bool   `json:"enabled"`
	Bucket  string `json:"bucket"`
	Region  string `json:"region"`
}

// Define top-level config
type Config struct {
	Hosts map[string]Host `yaml:"hosts"`
	S3    S3              `yaml:"s3"`
}

// Parse JSON file into Config struct
func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

// Print helper for debugging
func (c *Config) Print() {
	for name, host := range c.Hosts {
		fmt.Printf("Host: %s, PeerIP: %s, ExternalIP: %s, Port: %d, SSH: %s\n",
			name, host.PeerIP, host.ExternalIP, host.Port, host.SSHCommand)
	}
	fmt.Printf("S3: Enabled=%v, Bucket=%s, Region=%s\n",
		c.S3.Enabled, c.S3.Bucket, c.S3.Region)
}
