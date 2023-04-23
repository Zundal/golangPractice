package main

import "fmt"

func main() {
	// fmt format string go에서 문자열을 나타내주는 라이브러리
	var texts string = "Test"
	var nums int = 10

	type Person struct {
		name string
		age  int
	}
	p := Person{name: "Alice", age: 30} // 구조체 인스턴스 생성

	fmt.Println(texts)
	fmt.Println(nums)

	fmt.Println(p.name, p.age) // 구조체 필드 출력

}
