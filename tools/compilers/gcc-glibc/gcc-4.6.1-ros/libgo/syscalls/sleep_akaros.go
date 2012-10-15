// sleep_rtems.go -- Sleep on RTEMS.

// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

func Sleep(nsec int64) (errno int) {
	print("Sleep not yet implemented!")
	errno = ENOSYS;
	return
}
