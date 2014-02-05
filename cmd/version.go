package cmd

import "fmt"

// Please use 3 digit versioning (major, minor, patch).
// http://semver.org
const VERSION = "0.0.5"

type Version struct{}

func NewVersion() *Version {
	return &Version{}
}

func (v Version) Definition() string {
	return "Show version of this command"
}

func (v Version) Exec(args []string) error {
	fmt.Println(VERSION)
	return nil
}