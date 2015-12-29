/*
 * Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */
package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	CmdName = "autotune"
)

func subCmd(cmds ...string) string {
	return fmt.Sprintf("%s %s", CmdName, strings.Join(cmds, " "))
}

func usage() {
	usage := `
Usage: %s [command]

Available commands:
    network stats Get network utilization over a period of time.
    network pid   Figure out the profile of the machine based on
                  network processes that are running on the machine.

    memory stats  FUTURE
    memory pid    FUTURE

    io stats      FUTURE
    io pid        FUTURE

    cpu stats     FUTURE
    cpu pid       FUTURE

    profile       FUTURE: Guess signature of the machine based on memory,
                  network and IO usage.

    server        Update settings based on profile of server.

Autotune by Anatma.
Copyright (c) 2015-2016. Abhi Yerra.
https://anatma.co/autotune
`

	fmt.Printf(usage, CmdName)
}

func main() {
	var (
		err error
	)

	if len(os.Args) < 2 {
		usage()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "network":
		switch os.Args[2] {
		case "stats":
			network := NewNetworkStats()
			network.ParseArgs(os.Args[3:])
			err = network.Run()
		case "pid":
			network := NewNetworkPid()
			network.ParseArgs(os.Args[3:])
			err = network.Run()
		}
	case "server":
		agent := NewServer()
		agent.ParseArgs(os.Args[2:])
		err = agent.Run()
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
