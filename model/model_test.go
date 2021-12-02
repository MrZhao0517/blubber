package model_test

import (
	"encoding/json"
	"fmt"
	"github.com/whale-clouds/service.blubber.utils/model"
	"testing"
)

type name struct {
	TestModel string `json:"test_model"`
}

func TestName(t *testing.T) {
	a := &name{TestModel: "123"}
	m, err := model.Map(a)
	if err != nil {
		return
	}
	marshal, err := json.Marshal(m)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}
