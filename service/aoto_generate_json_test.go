package service

import (
	"testing"
	"fmt"
	"reflect"
	"strings"
)

type UserName struct {
	Id string
	Name string
}

type BankInfo struct {
	Name string `json:"name"`
	CardNum string `json:"card_num"`
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

// 测试动态参数利用反射生成对应结构体
func TestReflect2Struct(t *testing.T){
	oneStepTest(func() {

		// 利用反射查看slice
		var s []*interface{}
		slice := reflect.TypeOf(s)
		fmt.Println(slice)
		// 利用反射看结构体

		var name UserName
		model := reflect.TypeOf(name)
		fmt.Println(model)

		//var obj []interface{}
		bins := []BankInfo{}
		m := reflect.TypeOf(bins)
		fmt.Println(m)
		//if strings.Contains(m.String(), "Bank"){
		//	obj = []m
		//}
	})
}