package main

import (
    "io/ioutil"
    "net/http"
    "fmt"
    "encoding/json"
)

type Currency struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Code string `json:"code"`
}

type DataS struct {
    Currencies [] Currency `json:"currencies"`
}

type CurrencyApiResponse struct {
    Status string `json:"status"`
    Data DataS `json:"data"`

}


func main(){
    getUrl := "api";
    client := &http.Client{}
    req, _ := http.NewRequest("GET", getUrl, nil)
    req.Header.Set("Content-Type", "application/json")
    res, _ := client.Do(req)

    data, _ := ioutil.ReadAll(res.Body)

    fmt.Println(string(data))

    var apiResponse CurrencyApiResponse;
    err := json.Unmarshal(data, &apiResponse)
    if err != nil {
        panic(err)
    }

    fmt.Println(apiResponse)
    fmt.Println(apiResponse.Data.Currencies[0])
}