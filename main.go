package main

//import "os"

func main() {
	a := App{}
	a.Initialize("./spa.sqlite")
	a.Run(":8080")
}