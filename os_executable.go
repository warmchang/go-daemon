package daemon

import (
	"os"
)

func osExecutable() (string, error) {
	return os.Executable()
}
