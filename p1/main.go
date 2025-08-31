package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {

	fmt.Printf("----------start of testcase--------------")
	type host_info struct {
		Peer_ip     string `yaml:"peer_ip"`
		External_ip string `yaml:"external_ip"`
		Port        int    `yaml:"port"`
		Ssh_command string `yaml:"ssh_command"`
	}

	type s3 struct {
		Enabled bool   `yaml:"enabled"`
		Bucket  string `yaml:"bucket"`
		Region  string `yaml:"region"`
	}

	//var hosts map[string]host_info

	type data struct {
		Hosts map[string]host_info `yaml:"hosts"`
		S3    s3                   `yaml:"s3"`
	}

	file_data, err := os.ReadFile("data.yaml")
	if err != nil {
		fmt.Printf("An error occured")
	}

	if len(file_data) == 0 {
		fmt.Println("File is empty")
	} else {
		fmt.Println("File read successfully, size:", len(file_data), "bytes")
	}

	var data_obj data
	err = yaml.Unmarshal(file_data, &data_obj)
	if err != nil {
		fmt.Println("An error occured")
	}
	fmt.Printf("length of host %d", len(data_obj.Hosts))

	fmt.Println(" individual value: " + data_obj.Hosts["host-1"].Peer_ip)

	for name, host := range data_obj.Hosts {
		fmt.Printf(" name :  %s, Peer_ip: %s, external_ip %s,port %d, ssh_command %s \n",
			name, host.Peer_ip, host.External_ip, host.Port, host.Ssh_command)
	}
	fmt.Printf("S3 buckets , enable: %t , bucket: %s, region %s",
		data_obj.S3.Enabled, data_obj.S3.Bucket, data_obj.S3.Region)

	fmt.Printf("----------end of testcase--------------")

}
