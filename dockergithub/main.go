package main

import(
"fmt"
"os")

func main(){
    gopath:=os.Getenv("GOPATH")
    fmt.Println("go path",gopath)
    goroot:= os.Getenv("GOROOT")
    fmt.Println("go root",goroot)
    testvar:= os.Getenv("TEST_VAR")
    fmt.Println("test variable",testvar)
}