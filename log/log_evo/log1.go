package log_evo

import "os"

func main() {

}

func Log(entry string) {
	logFile := "/var/log/info.log"
	file, _ := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	file.Write([]byte(entry))
}