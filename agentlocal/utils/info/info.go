package info

import (
	"fmt"
	"runtime"
)
// #include <unistd.h>
import "C"

type GoInfoObject struct {
	GoOS     string
	Kernel   string
	Core     string
	Platform string
	OS       string
	Hostname string
	CPUs     int
	RAM      int64
}

func (gi *GoInfoObject) VarDump() {
	fmt.Println("GoOS:", gi.GoOS)
	fmt.Println("Kernel:", gi.Kernel)
	fmt.Println("Core:", gi.Core)
	fmt.Println("Platform:", gi.Platform)
	fmt.Println("OS:", gi.OS)
	fmt.Println("Hostname:", gi.Hostname)
	fmt.Println("CPUs:", gi.CPUs)
}

type IInfo interface {
	Get() *GoInfoObject
}

func Platform() *GoInfoObject {
	var platform IInfo
	switch runtime.GOOS {
	case "linux":
		platform = &linux{}
	case "darwin":
		platform = &darwin{}
	case "windows":
		platform = &windows{}
	default:
		panic("OS not supported")
	}

	result:=  platform.Get()
	result.RAM = int64(C.sysconf(C._SC_PHYS_PAGES)*C.sysconf(C._SC_PAGE_SIZE))
	return result
}
