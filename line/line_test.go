package line

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"testing"
)

type TestData struct {
	AccessToken string
}

func init() {
}

func TestPost(t *testing.T) {

	var conf TestData
	if _, err := toml.DecodeFile("test_data.toml", &conf); err != nil {
		if err != nil {
			t.Fail()
		}
	}
	lineNotifer := New(conf.AccessToken)
	err := lineNotifer.Post("test")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
