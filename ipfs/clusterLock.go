package ipfs

import (
	"errors"
	"fmt"
	"io"
	"path"

	fslock "github.com/ipfs/go-fs-lock"
	"gitlab.com/vocdoni/go-dvote/log"
)

// lock logic heavily inspired by go-ipfs/repo/fsrepo/lock/lock.go

// The name of the file used for locking
const lockFileName = "cluster.lock"

var locker *lock

// lock helps to coordinate procees via a lock file
type lock struct {
	lockCloser io.Closer
	path       string
}

func (l *lock) lock() {
	if l.lockCloser != nil {
		checkErr("", errors.New("cannot acquire lock twice"))
	}

	// we should have a config folder whenever we try to lock
	makeConfigFolder()

	// set the lock file within this function
	log.Debugf("checking lock")
	lk, err := fslock.Lock(l.path, lockFileName)
	if err != nil {
		log.Debug(err)
		l.lockCloser = nil
		errStr := "%s. If no other "
		errStr += "%s process is running, remove %s, or make sure "
		errStr += "that the config folder is writable for the user "
		errStr += "running %s."
		errStr = fmt.Sprintf(
			errStr,
			err,
			ProgramName,
			path.Join(l.path, lockFileName),
			ProgramName,
		)
		checkErr("obtaining execution lock", errors.New(errStr))
	}
	log.Debugf("%s execution lock acquired", ProgramName)
	l.lockCloser = lk
}

func (l *lock) tryUnlock() error {
	// Noop in the uninitialized case
	if l.lockCloser == nil {
		log.Debug("locking not initialized, unlock is noop")
		return nil
	}
	err := l.lockCloser.Close()
	if err != nil {
		return err
	}
	log.Debug("successfully released execution lock")
	l.lockCloser = nil
	return nil
}
