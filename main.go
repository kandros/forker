package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	err := exec.Command("git", "remote", "rename", "origin", "upstream").Run()
	if err != nil {
		log.Fatalf("Failed to rename remote \n%v", err)
	}

	upstream, err := exec.Command("git", "remote", "get-url", "upstream").Output()
	if err != nil {
		log.Fatalf("Failed to read current upstream remote url \n%v", err)
	}
	v := reverse(string(upstream))
	if v[0] == '/' {
		v = v[1:]
	}

	projectName := strings.Split(reverse(v), "/")[0]

	err = exec.Command("git", "remote", "add", "origin", fmt.Sprintf("git@github.com:kandros/%s.git", strings.Trim(projectName, " \r\n"))).Run()
	if err != nil {
		log.Fatalf("Failed to add origin remote \n%v", err)
	}

	cmd := exec.Command("git", "remote", "-v")
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to show remotes \n%v", err)
	}

}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
