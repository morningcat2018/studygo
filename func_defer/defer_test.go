package func_defer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"runtime"
	"sort"
	"testing"
)

func Test1(t *testing.T) {
	if v := math.Pow(2, 10); v > 1000 {
		fmt.Printf("2^10=%f  > 1000\n", v) // v的作用域只在当前 {} 内
	}

	var osName = ""
	switch os := runtime.GOOS; os {
	case "darwin":
		osName = "MacOs"
	case "linux":
		osName = "Linux"
	default:
		osName = os
	}
	fmt.Println(osName)
}

func Test2(t *testing.T) {
	var buffer1 bytes.Buffer // 类似 java 中的 StringBuilder
	buffer1.WriteByte('m')
	buffer1.WriteString("cat")
	fmt.Println(string(buffer1.Bytes()))

	var buffer2 *bytes.Buffer = new(bytes.Buffer)
	buffer2.Write(buffer1.Bytes())
	buffer2.WriteRune(48)
	fmt.Println(string(buffer2.Bytes()))

	var buffer3 *bytes.Buffer = bytes.NewBuffer(buffer2.Bytes())
	buffer3.WriteString(" hi")
	fmt.Println(string(buffer3.Bytes()))
}

func Test3(t *testing.T) {
	// map -> json string
	var map1 = make(map[string]interface{})
	map1["name"] = "mc2018"
	map1["age"] = 28
	map1["hobby"] = [...]string{"吃", "玩", "看动漫", "读书"}
	bytes, err := json.Marshal(map1)
	if err != nil {
		return
	}
	fmt.Println(string(bytes))

	// obj -> json string
	type Person struct {
		Name  string   `json:"姓名"`
		Age   uint8    `json:"年龄"`
		Hobby []string `json:"爱好"`
	}
	p1 := Person{"mc2018", 28, []string{"read", "eat"}}
	bytes2, err2 := json.Marshal(p1)
	if err2 != nil {
		return
	}
	fmt.Println(string(bytes2))
}

func Test4(t *testing.T) {
	var x = 10
	v := reflect.ValueOf(x)
	fmt.Println(v.Kind())
	fmt.Println(v.Kind() == reflect.Int)

	w := reflect.TypeOf(x)
	fmt.Println(w.Name())
	fmt.Println(w.Name() == "int")
}

func Test5(t *testing.T) {
	var ints = [...]int{7, 3, 9, 2, 7, 1, 0, 45, 6}
	a := sort.IntSlice(ints[:])
	sort.Sort(a)
	fmt.Println(a)
}
