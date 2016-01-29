/* Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package fd

import (
	"github.com/anatma/procfs"
)

type ProcessFD struct {
	proc procfs.Proc

	Descriptors map[string]string
}

func NewProcess(proc procfs.Proc) *ProcessFD {
	pfd := &ProcessFD{proc: proc}

	fd, err := proc.NewFD()
	if err != nil {

	}

	pfd.Descriptors = fd

	return pfd
}
