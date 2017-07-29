package main

import (
    "fmt"
    "os"
    "os/user"
    "gitlab.com/ishankhare07/monkey-lang/repl"
)

func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }

    fmt.Println("Hello", user.Username, "! This is the Monkey Programming Language!")
    fmt.Println("Feel free to type in commands")
    repl.Start(os.Stdin, os.Stdout)
}
