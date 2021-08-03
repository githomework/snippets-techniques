package main

import (
	"log"
	"os/exec"
)

func do() {

	cmd := exec.Command("taskkill", "/f", "/im", "excel.exe")
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err, string(b))
		return
	}
	log.Println(string(b))
}
