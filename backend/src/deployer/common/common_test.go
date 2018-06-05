package common

import (
	"testing"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func Test_json_node(t *testing.T) {
	node := Node{
		Id: "one",
		Status: "loading",
		Port: 12345,
	}
	s,err := json.Marshal(node)
	assert.Nil(t, err)
	fmt.Println(fmt.Println(string(s)))
}
