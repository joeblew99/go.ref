// Copyright 2016 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lockutil

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"syscall"
)

// v1 improves upon v0 in two ways:
//
// - reduces the false positive rate of StillHeld by using the lock holder's PID
//   even when the 'ps' information is not available.
//
// - reduces the false negative rate of StillHeld by trying to determine when
//   it's appropriate to compare the lock holder's process information (the lock
//   may be on NFS or on a volume shared among containers).  Ideally, we'd seek
//   out and compare this information if and only if the lock seeker is sharing
//   a PID space with the lock holder.  In practice, this is hard to determine
//   correctly on all systems.  The price of a false positive when deciding if
//   the PID space is the same is much higher than that of a false negative
//   (since it results in a false negative for StillHeld), and hence we prefer
//   to err on the side of false negatives.  The strategy employed is:
//
//   - if a system id is available (machine-id on linux, serial number on
//     darwin), then rely on that when deciding if comparing process information
//     is appropriate
//
//   - otherwise, compare the timestamp of /proc/1 (in linux) when deciding if
//     comparing process information is appropriate.
//
//   Notes:
//
//   On linux, matching machine-id gives high confidence that process
//   information can be compared.  However, there are a few corner cases that
//   can result in false negatives:
//
//   - some systems are missing the machine-id (generated by dbus-uuidgen in
//     /var/lib/dbus/machine-id or by systemd-machine-id-setup in
//     /etc/machine-id); e.g. some Docker containers.  Falling back on comparing
//     /proc/1's timestamp helps here, though fails to work when the system has
//     rebooted; it also has a theoretical possibility of false positives if two
//     systems' /proc/1 timestamps happen to match
//
//   - on some systems (where /etc is mounted on tmpfs) the machine-id changes
//     with each reboot
//
//   - some systems may share PID space even if they have different machine-ids
//     (e.g., Docker containers can be configured to share PID space with the
//     host)
//
//   On darwin, the serial number should give high confidence that process
//   information can be compared. MacOS installations inside VM should have
//   their own serial numbers.  The only concern is if we fail to identify this,
//   in which case we have no fallback and default to assuming the process
//   information is not comparable.

const (
	systemIDLabel = "SYSTEM ID"
	pidLabel      = "PID"
	unknownID     = "UNKNOWN"
)

func makePsCommandV1(pid int) *exec.Cmd {
	return exec.Command("ps", "-o", "pid,lstart,user,comm", "-p", strconv.Itoa(pid))
}

// createV1 writes information about the current process (like host-identifying
// information and the process' PID) to the specified writer.  If some of the
// information cannot be determined, UNKNOWN is written instead.
func createV1(w io.Writer) error {
	// Write some system-dependent ID.
	sysID, err := getSystemID()
	if err != nil {
		return err
	}
	if _, err := fmt.Fprintf(w, "%s:%s\n", systemIDLabel, sysID); err != nil {
		return err
	}
	// Write the PID.
	if _, err := fmt.Fprintf(w, "%s:%d\n", pidLabel, os.Getpid()); err != nil {
		return err
	}
	if _, err := exec.LookPath("ps"); err != nil {
		// No 'ps' command available.
		return nil
	}
	cmd := makePsCommandV1(os.Getpid())
	cmd.Stdout = w
	cmd.Stderr = nil
	return cmd.Run()
}

var pidRegexV1 = regexp.MustCompile("\n\\s*(\\d+)")

func stillHeldV1(info []byte) (bool, error) {
	sysID, infoLeft, err := parseValue(info, systemIDLabel)
	if err != nil {
		return false, err
	}
	mySysID, err := getSystemID()
	if err != nil {
		return false, err
	}
	if sysID != mySysID || sysID == unknownID {
		// The locker's systemID doesn't match ours.  Assume the lock is
		// on a shared filesystem, created by a process on another
		// system.  Since we can't verify that process' liveness, assume
		// the lock is still held.
		return true, nil
	}
	pidStr, infoLeft, err := parseValue(infoLeft, pidLabel)
	if err != nil {
		return false, err
	}
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return false, fmt.Errorf("couldn't parse PID from %s", pidStr)
	}
	switch err := syscall.Kill(pid, 0); err {
	case syscall.ESRCH:
		// No such PID.
		return false, nil
	case nil, syscall.EPERM:
		// Process pid is running, proceed to compare process details.
	default:
		// Unexpected error.
		return false, err
	}
	if len(infoLeft) == 0 {
		// No details available, assume locker process still running.
		return true, nil
	}
	if _, err := exec.LookPath("ps"); err != nil {
		// No 'ps' command available, assume locker process still
		// running.
		// We could just let the ps invocation below fail, but this may
		// be faster and avoid returning an error gratuitously.
		return true, nil
	}
	cmd := makePsCommandV1(pid)
	out, err := cmd.Output()
	if err != nil {
		return false, err
	}
	return bytes.Equal(infoLeft, out), nil
}
