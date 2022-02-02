package pluginmanager

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	network "github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	Plugins struct {
		Weather struct {
			Name  string
			Image string
		}
	}
}

func GetPlugins() {
	ymlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		fmt.Errorf("something went wrong while reading config file")
	}

	var config Config

	err = yaml.Unmarshal(ymlFile, &config)
	if err != nil {
		fmt.Errorf("Something went wrong unmarshalling config")
	}
	fmt.Println(config.Plugins.Weather.Name)
	startPlugin(config.Plugins.Weather.Image, config.Plugins.Weather.Name)
}

func startPlugin(image string, name string) (string, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Println("Unable to create docker client")
	}
	if err != nil {
		fmt.Println("Unable to get port")
	}
	config := &container.Config{
		Image:    image,
		Hostname: "weather",
	}
	networkConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{},
	}
	gatewayConfig := &network.EndpointSettings{
		Gateway: "gatewayname",
	}
	networkConfig.EndpointsConfig["smart-home-network"] = gatewayConfig

	cont, err := cli.ContainerCreate(
		context.Background(),
		config,
		nil,
		networkConfig,
		nil,
		name,
	)

	if err != nil {
		fmt.Println(err)
	}

	cli.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{})
	fmt.Printf("Container %s is started", cont.ID)
	return cont.ID, nil
}
