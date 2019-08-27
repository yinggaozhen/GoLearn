package stub

import (
	"bytes"
	"os/exec"
)

const (
	User     = "root"
	Password = ""
	Database = "test"
)

func TestSetup() {
	cmd := exec.Command("/data/github/GoLearn/mysql/stub/mysql_stub.sh")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
}