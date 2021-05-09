package bitm

import (
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

// Run takes the invocation args and transparently runs the intended command while intercepting sensitive info.
func Run(args []string) error {
	ourPath := args[0]
	cmdPath, err := LookPathSkip(filepath.Base(ourPath), ourPath)
	if err != nil {
		return err
	}

	cmd := exec.Command(cmdPath, args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func findExecutable(file string) error {
	d, err := os.Stat(file)
	if err != nil {
		return err
	}
	if m := d.Mode(); !m.IsDir() && m&0111 != 0 {
		return nil
	}
	return fs.ErrPermission
}

// LookPathSkip is the same as exec/LookPath but skips the given skipPath during lookup.
func LookPathSkip(file, skipPath string) (string, error) {
	path := os.Getenv("PATH")
	for _, dir := range filepath.SplitList(path) {
		if filepath.Clean(dir) == filepath.Clean(skipPath) {
			continue
		}
		if dir == "" {
			// Unix shell semantics: path element "" means "."
			dir = "."
		}
		path := filepath.Join(dir, file)
		if err := findExecutable(path); err == nil {
			return path, nil
		}
	}
	return "", &exec.Error{file, exec.ErrNotFound}
}
