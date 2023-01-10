package python

import (
	"fmt"
	"os/exec"
)

// Finder is the interface that wraps the basic Find method.
//
// Find seeks for a Python interpreter executable. It returns an Interpreter if it finds one, and must a return non-nil
// error value if it fails to find one.
type Finder interface {
	Find(name string) (Interpreter, error)
}

// A LocalFinder is an implementation of Finder that seeks for a Python interpreter on the local file system by
// scanning the system executable path.
type LocalFinder struct {
}

// Find seeks for a Python interpreter named name on the local filesystem.
//
// Internally Find is just a thin wrapper around exec.LookPath.
func (f *LocalFinder) Find(name string) (Interpreter, error) {
	path, err := exec.LookPath(name)
	if err != nil {
		return nil, fmt.Errorf("could not find Python on local filesystem: %v", err)
	}
	return &LocalInterpreter{
		path: path,
	}, nil
}
