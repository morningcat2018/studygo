package md5

import (
	"crypto/md5"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world")

	strValue := "hello world"
	hash := md5.New()
	_, error := hash.Write([]byte(strValue))
	if error != nil {
		fmt.Println(error)
		os.Exit(-1)
	}
	byteArray := hash.Sum(nil)
	md5Value := fmt.Sprintf("%x", byteArray)

	fmt.Println(md5Value)
}
