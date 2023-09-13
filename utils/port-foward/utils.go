package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
)

func ReadServicesFromJson(filePath string) Services {
	jsonFile, err := os.Open(filePath)
	defer jsonFile.Close()

	if err != nil {
		log.Fatal(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var services Services

	json.Unmarshal(byteValue, &services)

	if err := services.validate(); err != nil {
		log.Fatal(err)
	}

	return services
}

func PortForwardService(service Service, wg *sync.WaitGroup, done chan struct{}) {
	defer wg.Done()

	commandArray := []string{
		"kubectl",
		"port-forward",
		service.ServiceName,
		fmt.Sprintf("%d:%d", service.LocalhostPort,
			service.ContainerPort),
		"-n", service.Namespace,
	}

	cmd := exec.Command(
		commandArray[0],
		commandArray[1],
		commandArray[2],
		commandArray[3],
		commandArray[4],
		commandArray[5],
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	commandString := fmt.Sprintf(
		"%s %s %s %s %s %s",
		commandArray[0],
		commandArray[1],
		commandArray[2],
		commandArray[3],
		commandArray[4],
		commandArray[5],
	)

	fmt.Println(commandString)
	err := cmd.Run()
	if err != nil {
		fmt.Printf(
			"Error running the commandString %s: %v\n",
			commandString,
			err,
		)
	}

	select {
	case <-done:
		cmd.Process.Signal(os.Kill)
	default:
	}
}

func PortForwardServices(services Services) {
	var wg sync.WaitGroup

	done := make(chan struct{})

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signals
		fmt.Printf("Received signal %v, stopping...\n", sig)
		close(done)
	}()

	for _, service := range services.Services {
		wg.Add(1)
		go PortForwardService(service, &wg, done)
	}

	wg.Wait()
}
