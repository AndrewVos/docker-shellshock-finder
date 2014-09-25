package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/fsouza/go-dockerclient"
)

func containers() ([]*docker.Container, error) {
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		return nil, err
	}
	options := docker.ListContainersOptions{}
	apiContainers, err := client.ListContainers(options)
	if err != nil {
		return nil, err
	}
	var pids []*docker.Container
	for _, apiContainer := range apiContainers {
		container, err := client.InspectContainer(apiContainer.ID)
		if err != nil {
			return nil, err
		}
		if container.State.Running {
			pids = append(pids, container)
		}
	}
	return pids, nil
}

func main() {
	containers, err := containers()
	if err != nil {
		log.Fatal(err)
	}
	vulnerable := 0
	notVulnerable := 0

	for _, container := range containers {
		pid := strconv.Itoa(container.State.Pid)
		cmd := exec.Command(
			"sudo",
			"/tmp/nsenter",
			"--target", pid,
			"--mount",
			"--uts",
			"--ipc",
			"--net",
			"--pid",
			"/bin/bash",
			"-c",
			`/usr/bin/env x='() { :;}; /bin/echo VULNERABLE' /bin/bash -c "echo this is a test"`,
		)
		out, err := cmd.CombinedOutput()
		if strings.Contains(string(out), "VULNERABLE") {
			fmt.Printf("%v (%v) is vulnerable!\n", container.ID, container.Name)
			vulnerable += 1
		} else {
			notVulnerable += 1
			fmt.Printf("%v (%v) is NOT vulnerable!\n", container.ID, container.Name)
			if err != nil {
				log.Println(string(out))
				log.Fatal(err)
			}
		}
	}
	fmt.Printf("Found %v containers that are vulnerable to shellshock, and %v containers that are not\n", vulnerable, notVulnerable)
}
