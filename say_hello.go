package go_say_hello

import "strconv"

func SayHello(name string, age int) string {
	return "Hello " + name + " you are " + strconv.Itoa(age) + " years old"
}
