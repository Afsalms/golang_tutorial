package main

import (
    "io/ioutil"
    "net/http"
    "fmt"
    "encoding/json"
    "bytes"
)


type LoginRequest struct {
    Uid string `json:"uid"`
    Password string `json:"password"`
    UserTyoe string `json:"user_type"`
}

func main(){
    
    inputData := LoginRequest{"username", "password", "usr_type1"}
    requestBody, _ := json.Marshal(inputData)
    postUrl := "url";
    client := &http.Client{}
    req, _ := http.NewRequest("POST", postUrl, bytes.NewBuffer(requestBody))
    req.Header.Set("Content-Type", "application/json")
    res, _ := client.Do(req)
    data, _ := ioutil.ReadAll(res.Body)
    fmt.Println(string(data))
}