package python

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/slog"
	"os/exec"
	"regexp"
	"strings"
)

var (
	pattern = regexp.MustCompile(`^Python (\d+)(\.\d+)(\.\d+)$`)
)

type Runner interface {
	Run(path string, arg ...string) error
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

func (i *LocalInterpreter) Run(entrypoint string, arg ...string) error {
	cmd := i.command(append([]string{entrypoint}, arg...)...)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	go func() {
		s := bufio.NewScanner(stderr)
		for s.Scan() {
			slog.Warn(s.Text())
		}
	}()
	return cmd.Wait()
}

func (i *LocalInterpreter) Version() (string, error) {
	cmd := i.command("--version")
	b, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	version := strings.TrimSpace(string(b))
	if !pattern.MatchString(version) {
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
