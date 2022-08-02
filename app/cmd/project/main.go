package main

import "my_project/pkg/logging"

func main() {

	logging.Log().Warn("hi")
	test()

}

func test() {
	logging.Log().Warn("bye")
}
