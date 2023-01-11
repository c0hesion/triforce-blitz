package python

import (
	"fmt"
	"io"
	"os/exec"
	"regexp"
	"strings"
)

var (
	versionPattern = regexp.MustCompile(`^Python (\d+)(\.\d+)(\.\d+)$`)
)

type Runner interface {
	RunScript(path string, stdout io.Writer, stderr io.Writer, arg ...string) error
	Run(stdin io.Reader, stdout io.Writer, stderr io.Writer, arg ...string) error
}

type Interpreter interface {
	Runner
	Version() (string, error)
}

type LocalInterpreter struct {
	path string
}

func (i *LocalInterpreter) command(arg ...string) *exec.Cmd {
	return exec.Command(i.path, arg...)
}

func (i *LocalInterpreter) RunScript(entrypoint string, stdout io.Writer, stderr io.Writer, arg ...string) error {
	cmd := i.command(append([]string{entrypoint}, arg...)...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	if err := cmd.Start(); err != nil {
		return err
	}
	return cmd.Wait()
}

func (i *LocalInterpreter) Run(stdin io.Reader, stdout io.Writer, stderr io.Writer, arg ...string) error {
	cmd := i.command(arg...)
	cmd.Stdin = stdin
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	if err := cmd.Start(); err != nil {
		return err
	}
	return cmd.Wait()
}

func (i *LocalInterpreter) Version() (string, error) {
	cmd := i.command("--version")
	b, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	version := strings.TrimSpace(string(b))
	if !versionPattern.MatchString(version) {
		return version, fmt.Errorf("output does not match valid Python version")
	}
	return strings.TrimPrefix(version, "Python "), nil
}

func FindInterpreter(finder Finder) (Interpreter, error) {
	for _, name := range [...]string{"python", "python3", "python2"} {
		if interpreter, err := finder.Find(name); err == nil {
			return interpreter, nil
		}
	}
	return nil, fmt.Errorf("could not find Python executable")
}
