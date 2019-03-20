package main

import "fmt"

func replaceSpace(str string) string {
	length := len(str)
	count := 0
	for i := 0; i < length; i++ {
		if str[i] == ' ' {
			count ++
		}
	}
	newLength := length + count*2
	strByte := make([]byte, newLength)

	j := 0
	for i := length - 1; i >= 0; i-- {
		if str[i] == ' ' {
			strByte[newLength-1-j-2] = '%'
			strByte[newLength-1-j-1] = '2'
			strByte[newLength-1-j] = '0'

			j += 3
		} else {
			strByte[newLength-1-j] = str[i]
			j++
		}
	}
	return string(strByte)
}

func main() {
	str := "We Are Happy."
	fmt.Println(replaceSpace(str))
}
