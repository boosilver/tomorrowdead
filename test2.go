package main

import (
	"fmt"
	"net/http"
	
)
type Todo struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}
 
func (s *Client) AddTodo(todo *Todo) error {
	url := fmt.Sprintf(baseURL+"http://localhost:3754/date", s.Username)
	fmt.Println(url)
	j, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)
	return err
}