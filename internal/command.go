// package internal command.go
package internal

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Command struct.
type Command struct {
	Name string
	Args []string
	env  []string
	dir  string
}

// NewCommand new a Command instance.
//  @param name
//  @param args
//  @return *Command
func NewCommand(name string, args ...string) *Command {
	return &Command{
		Name: name,
		Args: args,
		env:  os.Environ(),
		dir:  ".",
	}
}

// SetDir set command workspace.
//  @receiver c
//  @param dir
//  @return *cmd
func (c *Command) SetDir(dir string) *Command {
	c.dir = dir

	return c
}

// AddEnv add an environment config.
//  @receiver c
//  @param env
//  @return *Command
func (c *Command) AddEnv(env []string) *Command {
	c.env = append(c.env, env...)

	return c
}

// Run command.
//  @receiver c
//  @return string
//  @return error
func (c *Command) Run() (string, error) {
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)

	e := exec.Command(c.Name, c.Args...)
	log.Printf("exec command: %+v", e)

	e.Env = c.env
	e.Dir = c.dir
	e.Stdout = stdout
	e.Stderr = stderr

	err := e.Run()
	if err != nil {
		log.Printf("exec command failed: err = %+v, stderr = %+v", err, stderr.String())
		return "", err
	}

	return strings.Trim(stdout.String(), "\n"), nil
}
