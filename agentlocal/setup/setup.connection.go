package setup

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"agentlocal/docker"
	"agentlocal/utils/color"
)

type connector struct {
}

func execute(cmds ...string) ([]byte, error) {
	type error interface {
		Error() []byte
	}
	var log []byte
	for _, cmd := range cmds {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			return log, err
		}
		log = append(log, out...)
	}
	return log, nil
}

func Pin() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(color.BCyan("1. Please type your connection PIN code or Email: "))
	connectionKey, _ := reader.ReadString('\n')
	fmt.Println()
	fmt.Println(connectionKey)

	docker.InstallCore()
	out, e := execute("sudo docker ps")
	if e != nil {
		fmt.Println(e)
	}
	str := string(out[:])
	fmt.Println(str)
}
