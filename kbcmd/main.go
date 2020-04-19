package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/killbill/kbcli/v2/kbcmd/cmdlib"
	"github.com/killbill/kbcli/v2/kbcmd/commands"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	r := cmdlib.NewApp()
	commands.RegisterAll(r)
	err := r.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
