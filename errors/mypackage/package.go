package mypackage

import "fmt"

func Run() {
	defer func() {
		println("end my function")
		r := recover()
		if r != nil {
			fmt.Println("error in run: ", r)
		}
	}()
	panic("panic in run function")
}
