package main

import (
    "fmt";
    "flag";
    "os";
    //"net/http";
    "net/url";
    //"strings"
)

type ParamError string

func (e ParamError) Error() string {
    return string(e)
}

var baseUrl, _ = url.Parse("https://api.nexmo.com")

func main() {
    if len(os.Args) <= 1 {
        fmt.Println("Expecting action as first argument")
        os.Exit(1)
    }

    action := os.Args[1]

    apiKey, key_err := parseParameter("apiKey", "NEXMO_API_KEY", "Nexmo API Key")

    if len(key_err) != 0 {
        fmt.Println("Failed to get Nexmo API Key: ", key_err)
        os.Exit(1)
    }

    apiSecret, sec_err := parseParameter("apiSecret", "NEXMO_API_SECRET", "Nexmo API Secret")
    if len(sec_err) != 0 {
        fmt.Println("Failed to get Nexmo API Secret: " + sec_err)
        os.Exit(1)
    }

    executeAction(action, apiKey, apiSecret)
}

func parseParameter(cmd_flag string, env_var string, description string) (param, err string) {
    flag.StringVar(&param, cmd_flag, "", description)

    if len(param) == 0 {
        param = os.Getenv(env_var)

        if len(param) == 0 {
            err = "Could not find " + description + " in -" + cmd_flag + " parameter or " + env_var + " environment variable"
            return
        }
    }

    return
}

func executeAction(action, apiKey, apiSecret string) {
    switch action {
    case "list-applications":
        fmt.Println("will do list-applications")
        os.Exit(0)

    default:
        fmt.Println("you said nothing interesting bye")
        os.Exit(2)
    }
}
