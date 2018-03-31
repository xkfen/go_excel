package service

import (
	"testing"
	"github.com/novalagung/gubrak"
	"fmt"
	"reflect"
)

func TestCreateTealegXlsx(t *testing.T) {
	oneStepTest(func() {
		CreateTealegXlsx()
	})
}

func TestReadTealegXlsx(t *testing.T) {
	oneStepTest(func() {
		ReadTealegXlsx()
	})
}

func TestChunk(t *testing.T) {
	oneStepTest(func() {
		Chunk()
	})
}

func TestConcat(t *testing.T){
	oneStepTest(func() {
		data := []int{ 1, 2, 3, 4 }
		dataConcat1 := []int{ 4, 6, 7 }
		dataConcat2 := []int{ 8, 9 }
		info, err := gubrak.Concat(data, dataConcat1, dataConcat2)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(info)
	})
}

func TestFindLastIndex(t *testing.T){
	oneStepTest(func() {
		//data := []int{ 1, 2, 2, 3, 7, 4, 5 }
		//predicate := func(each int) bool {
		//	return each == 3
		//}
		//fromIndex := 4
		//
		//result, err := gubrak.FindLastIndex(data, predicate, fromIndex)
		data := []string{ "damian", "grayson", "cass", "tim", "tim", "jason", "steph" }

		result, err := gubrak.FindLastIndex(data, func(each string) bool {
			return each == "tim"
		})
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(result)
	})
}

func TestIndexOf(t *testing.T){
	oneStepTest(func() {
		data := []float64{ 2.1, 2.2, 3, 3.00000, 3.1, 3.9, 3.95 }
		var i int
		var err error
		i, err = gubrak.IndexOf(data, 2.2)           // ===> 1
		fmt.Println(i)
		i, err = gubrak.IndexOf(data, 3)             // ===> -1
		fmt.Println(i)
		i ,err = gubrak.IndexOf(data, float64(3))    // ===> 2 (because 3 is detected as int32, not float64)
		fmt.Println(i)
		i, err = gubrak.IndexOf(data, float64(3), 2) // ===> 2
		fmt.Println(i)
		i, err = gubrak.IndexOf(data, float64(3), 3) // ===> 3
		fmt.Println(i)
		if err != nil {
			fmt.Println(err.Error())
		}
	})
}

func TestLastIndexOf(t *testing.T){
	oneStepTest(func() {
		data := []float64{ 2.1, 2.2, 3, 3.00000, 3.1, 3.9, 3.95 }
		t := reflect.TypeOf(data[3])
		fmt.Println(t)
		var i int
		i, _  = gubrak.LastIndexOf(data, 2.2)         // ===> 1
		fmt.Println(i)
		i , _ = gubrak.LastIndexOf(data, 3)             // ===> -1 (because 3 is detected as int32, not float64)
		fmt.Println(i)
		i, _ = gubrak.LastIndexOf(data, float64(3))    // ===> 3
		fmt.Println(i)
		i, _ = gubrak.LastIndexOf(data, float64(3), 2) // ===> 2
		fmt.Println(i)
		i, _ = gubrak.LastIndexOf(data, float64(3), 3) // ===> 3
		fmt.Println(i)
	})
}

func TestTakeRight(t *testing.T){
	oneStepTest(func() {
		data := []int{ 1, 2, 3, 4, 5, 6, 7, 8, 9 }
		take := 5

		result, _ := gubrak.TakeRight(data, take)
		fmt.Println(result)
	})
}