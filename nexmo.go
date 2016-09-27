package main

import (
    "fmt";
    "os";
    //"net/http";
    "net/url";
    //"strings"
)

var baseUrl, _ = url.Parse("https://api.nexmo.com")

func main() {
    if len(os.Args) <= 1 {
        fmt.Println("Need an action kthxbai")
        os.Exit(1)
    }

    action := os.Args[1]
    executeAction(action)
}

func executeAction(action string) {
    switch action {
    case "list-applications":
        fmt.Println("mkay will do list-applications")
        fmt.Println("AHAHAHA JUST KID I DO NOTHING BYE!")
        os.Exit(0)

    default:
        fmt.Println("you said nothing interesting bye")
        os.Exit(2)
    }
}
