package main

import (
	"io"
	"net/http"
	"log"
	"strings"
	"encoding/json"
	//"fmt"
)
var TasksList *Tasks
type Tasks struct {
	counter int
	list map[int][]string
}
func (t *Tasks) Init() {
	t.counter = 0
	t.list = make(map[int][]string)
}
func (t *Tasks) Add(url []string) int {
	t.counter++
	t.list[t.counter] = url
	return t.counter
}

func ParseUrlsList(r io.ReadCloser) []string {
	buf := make([]byte, 8)
	res := make([]byte, 8)
	for {
		n, err := r.Read(buf)
		res = append(res, buf[:n]...)
		if err == io.EOF {
			break
		}
	}
	return strings.Split(string(res), "\\n")
}

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		TaskId := TasksList.Add(ParseUrlsList(req.Body))
		res := make(map[string]int)
		res["TaskId"] = TaskId
		//json.NewEncoder(w).Encode(fmt.Sprintf("{ taskId: %d }", TaskId))
		json.NewEncoder(w).Encode(res)
	}
}

func main() {
	TasksList = new(Tasks)
	TasksList.Init()
	http.HandleFunc("/tasks", HelloServer)
	log.Fatal(http.ListenAndServe(":12345", nil))
}