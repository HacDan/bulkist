package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type Todo struct {
	Content   string `json:"content"`
	DueString string `json:"due_string"`
	DueLang   string `json:"due_lang"`
	Priority  int    `json:"priority"`
}

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
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func createTask(task string) {
	todo := Todo{
		Content: task,
	}
	payloadBytes, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api.todoist.com/rest/v2/tasks", body)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", os.ExpandEnv("Bearer ${TODOIST_TOKEN}"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
}
