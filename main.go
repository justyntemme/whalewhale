package main

import "flag"

func init() {
	keypath		:= flag.String("keypath", "~/.ssh/id_rsa", "Location of ssh key for public key encryption")
	pass		:= flag.String("pass", "", "Password for ssh authentication")
	remoteMachine	:= flag.String("r", "", "Remote machine to connect to")
}

func main() {
	flag.Parse()
}
