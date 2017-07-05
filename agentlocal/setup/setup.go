package setup

import (
	"agentlocal/utils/color"
	"agentlocal/utils/info"
	"fmt"
)

func Run() {
	platform := info.Platform()
	fmt.Println()
	fmt.Println(color.BCyan("Welcome to Sqdron Agent setup."))
	fmt.Println()

	fmt.Println(color.BCyan("Machine"))
	fmt.Println("Host:", platform.Hostname)
	fmt.Println("OS:", platform.Kernel, platform.Core)
	fmt.Println("CPU:", platform.CPUs, platform.Platform)
	fmt.Println("RAM:", platform.RAM/1024/1024)
	fmt.Println()

	// 1. Validate Connection
	Pin()

	// 2. Get connection data

	// 3. Install service
}
