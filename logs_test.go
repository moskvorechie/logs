package logs_test

import (
	"encoding/json"
	"github.com/moskvorechie/logs"
	"io/ioutil"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	_ = os.Remove("main.log")
	l, err := logs.New(&logs.Config{
		App:      "test",
		FilePath: "main.log",
	})
	if err != nil {
		t.Fatal(err)
	}
	l.Info("123")
	body, err := ioutil.ReadFile("main.log")
	if err != nil {
		t.Fatal(err)
	}
	type Message struct {
		Level    string `json:"level"`
		Time     int    `json:"time"`
		Datetime string `json:"datetime"`
		App      string `json:"app"`
		Message  string `json:"message"`
	}
	var m Message
	err = json.Unmarshal(body, &m)
	if err != nil {
		t.Fatal(err)
	}
	if m.Message != "123" {
		t.Fatal("Not eq Message")
	}
	if m.App != "test" {
		t.Fatal("Not eq App")
	}
}
