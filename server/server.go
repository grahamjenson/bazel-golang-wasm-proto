package server

import (
	"context"
	"encoding/json"
	"github.com/grahamjenson/bazel-golang-wasm-protoc/protos/api"
	"io/ioutil"
	"strings"
)

////
// ec2Instances file
////

type ec2Instance struct {
	PrettyName   string  `json:"pretty_name,omitempty"`
	InstanceType string  `json:"instance_type,omitempty"`
	ECU          float32 `json:"ECU,omitempty"`
	Memory       float32 `json:"memory,omitempty"`

	NetworkPerformance string `json:"network_performance,omitempty"`

	Pricing map[string]map[string]struct {
		OnDemand string `json:"ondemand,omitempty"`
	} `json:"pricing,omitempty"`
}

////
// Server
////

type Server struct {
	instances []*api.Instance
}

func (server *Server) Search(ctx context.Context, in *api.SearchRequest) (*api.Instances, error) {
	if server.instances == nil {
		server.parseInstances()
	}

	instances := []*api.Instance{}
	for _, instance := range server.instances {
		str, _ := json.Marshal(*instance)
		if strings.Contains(string(str), in.Query) {
			instances = append(instances, instance)
		}
	}

	return &api.Instances{Instances: instances}, nil
}

func (server *Server) parseInstances() {
	fileName := "external/com_github_ec2instances/file/instances.json"
	ec2Instances := []ec2Instance{}
	server.instances = []*api.Instance{}

	file, _ := ioutil.ReadFile(fileName)
	json.Unmarshal(file, &ec2Instances)

	for _, e := range ec2Instances {
		server.instances = append(server.instances, &api.Instance{
			Name:         e.PrettyName,
			InstanceType: e.InstanceType,
			Ecu:          e.ECU,
			Memory:       e.Memory,
			Network:      e.NetworkPerformance,
			Price:        e.Pricing["us-east-1"]["linux"].OnDemand,
		})
	}
}
