package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	clearScreen()

	fmt.Println("Please enter tasks, separated by a new line. Blank line when finished")
	reader := bufio.NewReader(os.Stdin)

	var tasks []string
	for {
		task, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		if len(strings.TrimSpace(task)) == 0 {
			break
		}

		tasks = append(tasks, task)
	}

	for _, task := range tasks {
		createTask(task)
	}
	fmt.Println("All tasks added successfully!")
}

func clearScreen() {
	switch runtime.GOOS {
	case "darwin":
		runCmd("clear")
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func createTask(task string) {
	params := url.Values{}
	params.Add("text", task)

	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "https://api.todoist.com/sync/v9/quick/add", body)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", os.ExpandEnv("Bearer ${TODOIST_TOKEN}"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
}
