package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-openapi/spec"
	"io/ioutil"
	"github.com/marshome/i-pkg/filesystem"
)

func main() {
	data, err := ioutil.ReadFile("./swagger.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s := &spec.Swagger{}
	err = json.Unmarshal(data, s)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	g := NewGenerator()
	code, err := g.Gen(s)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(code)

	filesystem.NewFile("./api.tsx", []byte(code))
}
