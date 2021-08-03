package main

import (
  "log"
	"strings"
	"github.com/mitchellh/go-ps"
)


func excelRunning() bool {

	processes, err := ps.Processes()
	if err != nil {
		log.Println(err)
		return false
	}
	for _, v := range processes {
		if strings.ToLower(v.Executable()) == "excel.exe" {
			return true
		}
		//log.Println(v.Executable())

	}
	return false
}
