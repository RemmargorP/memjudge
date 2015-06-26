package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/RemmargorP/myjudge"
	"github.com/sevlyar/go-daemon"
)

const Address = "localhost"

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Options:\nstart|stop")
		return
	}
	option := os.Args[1]

	query := Address + myjudge.MasterPort

	switch option {
	case "start":
		cmd := exec.Command("curl", query)
		state, err := cmd.CombinedOutput()
		fmt.Printf("CURL Response: %s %+v\n", state, err)
		if err != nil {
			cntx := &daemon.Context{
				PidFileName: "pid",
				PidFilePerm: 0644,
				LogFileName: "log",
				LogFilePerm: 0640,
				WorkDir:     "./",
				Umask:       027,
				Args:        []string{},
			}
			fmt.Println("Server has successfully started")
			child, _ := cntx.Reborn()
			if child != nil {
				return
			} else {
				defer cntx.Release()
				server := &myjudge.Server{}
				server.Config = myjudge.DefaultServerConfig()
				server.Serve()
			}

		} else {
			fmt.Println("Server is already running")
		}
	case "stop":
		state, err := exec.Command("curl", query+"/?option=stop").CombinedOutput()
		fmt.Printf("CURL Response: %s %+v\n", state, err)
		if err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Println("Options:\nstart|stop")
	}
}
