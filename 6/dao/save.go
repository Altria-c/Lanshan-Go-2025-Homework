package dao

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func SaveDatabase() {
	file, err := os.OpenFile("D:\\goProjects\\lansan-learning\\6\\data\\userinfo.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("open failed  erroe:", err)
		return
	}
	defer file.Close()

	write := bufio.NewWriter(file)

	for name, pwd := range database {
		content := fmt.Sprintf("%s:%s\n", name, pwd)
		_, err := write.WriteString(content)
		if err != nil {
			fmt.Println("write failed  error:", err)
			return
		}

	}
	write.Flush()

}

func LoadDatabase() {
	file, err := os.Open("D:\\goProjects\\lansan-learning\\6\\data\\userinfo.txt")
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Println("open failed error:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')

		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			username := strings.TrimSpace(parts[0])
			password := strings.TrimSpace(parts[1])
			database[username] = password
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read failed error:", err)
			return
		}
	}

}
