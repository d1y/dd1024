package main

import (
	"fmt"
	"os/exec"
)

func callbackPlayMpv(url string) {
	pwdCmd := exec.Command("mpv", url)
	pwdOutput, _ := pwdCmd.Output()
	fmt.Println(string(pwdOutput))
}