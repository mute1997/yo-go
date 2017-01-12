package main

import (
    "os"
    "fmt"
)

func main(){
    command := NewCommand(os.Args)

    err := command.exec()

    if err != nil {
        fmt.Println(err)
    }
}
