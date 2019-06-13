

package main

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "io"
    "log"
    "os"
    "io/ioutil"
    "net/http"
    "strings"
    "path/filepath"
    "time"
    // "sync"
)

type fileDetails struct{
    fileName string
    content string
}

func main() {
    
    fmt.Println("Started")
    start := time.Now()
    
    urlChannel := make(chan string, 100)
    fileDetailsChannel := make(chan fileDetails, 100)

    csvFile, err := os.Open("urls.csv")
    if err != nil {
        os.Exit(3)
    }
    reader := csv.NewReader(bufio.NewReader(csvFile))
    var urls []string
    for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
        }
        urls = append(urls, line[0])
    }

    for _, url := range(urls){
        urlChannel <- url

    }
    close(urlChannel)
    go getWebPage(urlChannel, fileDetailsChannel)
    go writeFile(fileDetailsChannel)
    elapsed := time.Since(start)
    log.Printf("Action took %s", elapsed)
    fmt.Scanln()
}


func getWebPage(urlChannel chan string, fileDetailsChannel chan fileDetails){
    for url := range urlChannel{
        fmt.Println(url)
        response, err := http.Get(url)
        if err != nil {
            fmt.Printf("The HTTP request failed with error %s\n", err)
        }else {
            data, _ := ioutil.ReadAll(response.Body)
            urlCompoents := strings.Split(url, "/")
            fileName := "htmls/" + urlCompoents[len(urlCompoents)-1] + ".html"
            fileDetailsChannel <- fileDetails{fileName: fileName, content: string(data)}
        }
    }
    close(fileDetailsChannel)
}

func writeFile(fileDetailsChannel chan fileDetails){
    newpath := filepath.Join(".", "htmls")
    os.MkdirAll(newpath, os.ModePerm)
    for fileDetail := range(fileDetailsChannel){
        fileName := fileDetail.fileName
        data := fileDetail.content
        f, err := os.Create(fileName)
        if err != nil {
            fmt.Println(err)
            continue
        }
        _, err1 := f.WriteString(data)
        if err1 != nil {
            f.Close()
            continue
        }
    }
}


