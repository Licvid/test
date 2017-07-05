package docker

import (
	"fmt"
	"strings"

	"agentlocal/bash"
	"agentlocal/utils/color"
	/*
		"github.com/docker/docker/api/types"
		"github.com/docker/docker/api/types/container"
		"github.com/docker/docker/client"
		"github.com/fsouza/go-dockerclient" */)

func Uninstall() {
	var answer string
	fmt.Println("Delete docker?(Y/N)")
	fmt.Scanln(&answer)

	if answer == "y" || answer == "Y" || answer == "yes" || answer == "Yes" || answer == "YES" {
		bash.Execute("sudo apt-get -y remove docker*", "sudo apt-get -y autoremove", "sudo apt-get autoclean")
	}

}

func InstallCore() {
	var answer string
	fmt.Print(color.BCyan("2. Install docker?(Y/N)"))
	fmt.Scanln(&answer)
	answer = strings.ToUpper(answer)
	if answer == "Y" || answer == "YES" {
		bash.Execute("sudo snap install docker")
	}

}

/*
func Install() {
	var answer string
	fmt.Println("Install docker?(Y/N)")
	fmt.Scanln(&answer)

	if answer == "y" || answer == "Y" || answer == "yes" || answer == "Yes" || answer == "YES" {
		bash.Execute("sudo apt-get update",
			"sudo apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D",
			"sudo apt-add-repository 'deb https://apt.dockerproject.org/repo ubuntu-xenial main'",
			"sudo apt-get update",
			"sudo apt-cache policy docker-engine",
			"sudo apt-get install -y docker-engine")
	}

}

func ImageList() {
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}
	imgs, err := client.ListImages(docker.ListImagesOptions{All: false})
	if err != nil {
		panic(err)
	}
	for _, img := range imgs {
		fmt.Println("ID: ", img.ID)
		fmt.Println("RepoTags: ", img.RepoTags)
		fmt.Println("--------------")
	}
}

func ContainerList() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Println(container.ID)
	}
}

func RunImage(image string) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	imageName := image

	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

}
*/
