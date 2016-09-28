package main

import (
    "fmt";
    "flag";
    "os";
    "net/http";
    "net/url";
    "io/ioutil";
    "encoding/json";
    //"strings"
)

type ParamError string

type Application struct {
    Id string
    Name string
    Keys map[string]string
}

type ApplicationList struct {
    Embedded struct {
        Applications []Application
    } `json:"_embedded"`
}

func (appList ApplicationList) Applications() []Application {
    return appList.Embedded.Applications
}

func (e ParamError) Error() string {
    return string(e)
}

var baseUrlPtr, _ = url.Parse("https://api.nexmo.com")
var baseUrl = *baseUrlPtr

func main() {
    action, apiKey, apiSecret, err := parseCommandArgs()
    if len(err) != 0 {
        fmt.Println("Not all required arguments were found: ", err)
        os.Exit(2)
    }

    v := url.Values{}
    v.Add("api_key", apiKey)
    v.Add("api_secret", apiSecret)
    baseUrl.RawQuery = v.Encode()

    executeAction(action, apiKey, apiSecret)
}

func parseCommandArgs() (action, apiKey, apiSecret string, err ParamError) {
    flag.StringVar(&apiKey, "apiKey", os.Getenv("NEXMO_API_KEY"), "Nexmo API Key")
    flag.StringVar(&apiSecret, "apiSecret", os.Getenv("NEXMO_API_SECRET"), "Nexmo API Secret")
    flag.StringVar(&action, "action", "", "Action to execute")
    flag.Parse()

    if len(apiKey) == 0 {
        err = "Could not find api key in parameter '-apiKey' or environment variable NEXMO_API_KEY"
        return
    }

    if len(apiSecret) == 0 {
        err = "Could not find api secret in parameter '-apiSecret' or environment variable NEXMO_API_SECRET"
        return
    }

    if len(action) == 0 {
        err = "No action provided."
    }

    return
}

func executeAction(action, apiKey, apiSecret string) bool {
    switch action {
    case "list-applications":
        return listApplications()

    default:
        fmt.Println("Unknown action", action)
    }

    return false
}

func listApplications() bool {
    applicationUrl := baseUrl
    applicationUrl.Path = "v1/applications"

    fmt.Println(baseUrl.String(), applicationUrl.String())
    resp, err := http.Get(applicationUrl.String())
    if err != nil {
        fmt.Println("Failed to list applications: ", err)
        return false
    }

    bytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Failed to read body of request: ", err)
        return false
    }

    if resp.StatusCode < 200 || resp.StatusCode >= 300 {
        fmt .Println("Failed to perform http request. Failed with error: ", resp.Status)
        return false
    }

    var appList ApplicationList
    err = json.Unmarshal(bytes, &appList)

    for _, app := range appList.Applications() {
        fmt.Println(app.Id)
    }

    return true
}
