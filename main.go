package main

import (
	"flag"
	"fmt"

	"github.com/justyntemme/whalewhale/ssh"
)

func main() {
	keypath		:= flag.String("keypath", "~/.ssh/id_rsa", "Location of ssh key for public key encryption")
	pass		:= flag.String("pass", "", "Password for ssh authentication")
	remoteMachine	:= flag.String("r", "", "Remote machine to connect to")

	flag.Parse()
	ssh.CreateSession(*keypath,*pass,*remoteMachine)

}
