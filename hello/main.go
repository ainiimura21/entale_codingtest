package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)
type Article struct {
    ID int `json:"id"`
    Title string `json:"title"`
    Body  string `json:"body"`
    Medias []Media `json:"medias"`
}
type Media struct {
    ID int `json:"id"`
    ContentURL string `json:"contentUrl"`
    ContentType string `json:"contentType"`
}

func main() {
    response, err := http.Get("https://gist.githubusercontent.com/gotokatsuya/cc78c04d3af15ebe43afe5ad970bc334/raw/dc39bacb834105c81497ba08940be5432ed69848/articles.json")
    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    var responseObject Article
    json.Unmarshal(responseData, &responseObject)

    fmt.Println(responseObject.ID)
    fmt.Println(len(responseObject.Medias))

    for i := 0; i < len(responseObject.ID); i++ {
        fmt.Println(responseObject.Medias[i].ContentURL)
    }

}