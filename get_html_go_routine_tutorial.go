

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
    "sync"
)

type fileDetails struct{
    fileName string
    content string
}

func main() {
    
    fmt.Println("Started")
    start := time.Now()
    var waitgroup sync.WaitGroup
    
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
    waitgroup.Add(20)
    for i := 0; i < 10; i++ {
        go getWebPage(urlChannel, fileDetailsChannel, &waitgroup)
        go writeFile(fileDetailsChannel, &waitgroup)
    }
    waitgroup.Wait()
    elapsed := time.Since(start)
    log.Printf("Action took %s", elapsed)
}


func getWebPage(urlChannel chan string, fileDetailsChannel chan fileDetails, waitgroup *sync.WaitGroup){
    for url := range urlChannel{
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
    // close(fileDetailsChannel) //#issue 
    waitgroup.Done()
}

func writeFile(fileDetailsChannel chan fileDetails, waitgroup *sync.WaitGroup){
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
    waitgroup.Done()

}


