package logs_test

import (
	"encoding/json"
	"github.com/moskvorechie/logs"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	_ = os.Remove("main.log")
	l, err := logs.New(&logs.Config{
		App:      "test",
		FilePath: "main.log",
		Clear:    true,
	})
	if err != nil {
		t.Fatal(err)
	}
	q := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec luctus sit amet augue id accumsan. Integer sed massa ipsum. Phasellus convallis pellentesque faucibus. Etiam at ipsum lacinia, feugiat elit ac, accumsan velit. In porta neque sed auctor dapibus. Aenean eleifend eget tortor a luctus. Nunc ornare, elit id maximus vulputate, nunc augue tincidunt nulla, vitae tempus risus odio vel risus. Fusce elementum, lectus ac aliquet pretium, turpis massa luctus nisl, ac pharetra risus erat quis enim. Duis non magna pharetra, pulvinar purus quis, pellentesque leo. Pellentesque placerat molestie eros tempor aliquet. Vestibulum sodales aliquam venenatis. Aliquam erat volutpat. Sed ut egestas mi. In aliquet, tortor non sollicitudin cursus, ipsum diam pretium justo, id posuere mi lectus eget sem.`
	for k := 0; k < 1000000; k++ {
		if k%100000 == 0 {
			time.Sleep(1 * time.Second)
		}
		l.Info(q)
	}

	return

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
