package service

import (
	"testing"
	"fmt"
)

type UserName struct {
	Id string
	Name string
}



func TestGetTagName(t *testing.T){
	oneStepTest(func() {
		name := getTagName("UserName", "json")
		fmt.Println(name)
	})
}

func TestProduceStructTag(t *testing.T) {
	var name UserName
	oneStepTest(func() {
		json := ProduceStructTag(name, "json")
		fmt.Println(json)
	})
}