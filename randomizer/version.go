package randomizer

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Version struct {
	Major       int
	Minor       int
	Patch       int
	Branch      string
	BranchMajor int
	BranchMinor int
}

func (v *Version) IsBranchVersion() bool {
	return v.Branch != ""
}

func (v *Version) String() string {
	if v.IsBranchVersion() {
		return fmt.Sprintf(
			"%d.%d.%d-%s-%d.%d",
			v.Major, v.Minor, v.Patch,
			v.Branch, v.BranchMajor, v.BranchMinor,
		)
	} else {
		return fmt.Sprintf(
			"%d.%d.%d",
			v.Major, v.Minor, v.Patch,
		)
	}
}

var (
	versionPattern = regexp.MustCompile(`^v?\d+(\.\d+){2}(-\w+-\d+(\.\d+))?$`)
)

func ParseVersionFromString(s string) (v Version, err error) {
	if !versionPattern.MatchString(s) {
		return v, errors.New("version does not match pattern")
	}
	s = strings.TrimPrefix(s, "v")
	versionStr, branch, branchVersionStr := versionTokens(s)
	v.Major, v.Minor, v.Patch, err = parseVersion(versionStr)
	if err != nil {
		return
	}
	if branch != "" && branchVersionStr != "" {
		v.Branch = branch
		v.BranchMajor, v.BranchMinor, err = parseBranchVersion(branchVersionStr)
	}
	return
}

func versionTokens(s string) (version string, branch string, branchVersion string) {
	tokens := strings.Split(s, "-")
	if len(tokens) >= 1 {
		version = tokens[0]
	}
	if len(tokens) >= 3 {
		branch = tokens[1]
		branchVersion = tokens[2]
	}
	return
}

func parseVersion(s string) (major int, minor int, patch int, err error) {
	if s == "" {
		return 0, 0, 0, errors.New("string cannot be empty")
	}
	tokens := strings.Split(s, ".")
	if len(tokens) != 3 {
		return 0, 0, 0, errors.New("invalid version string")
	}
	major, err = strconv.Atoi(tokens[0])
	if err != nil {
		return
	}
	minor, err = strconv.Atoi(tokens[1])
	if err != nil {
		return
	}
	patch, err = strconv.Atoi(tokens[2])
	return
}

func parseBranchVersion(s string) (major int, minor int, err error) {
	if s == "" {
		return 0, 0, errors.New("string cannot be empty")
	}
	tokens := strings.Split(s, ".")
	major, err = strconv.Atoi(tokens[0])
	if err != nil {
		return
	}
	minor, err = strconv.Atoi(tokens[1])
	return
}
