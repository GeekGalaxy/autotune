/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"log"
)

type Files struct{}

func (k *Files) Synopsis() string {
	return "Show the file changes needed for the signature"
}

func (k *Files) Help() string {
	return "Show the file changes needed for the signature"
}

func (k *Files) Run(args []string) int {
	if len(args) < 1 {
		log.Println("need to pass the signature")
		return -1
	}

	gaInvokeEvent("files", args[0])

	profile, err := profiles.Get(args[0], false)
	if err != nil {
		log.Println(err)
		return -1
	}

	profile.PrintFiles()

	return 0
}
