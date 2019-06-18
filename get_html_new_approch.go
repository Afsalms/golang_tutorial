

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
    waitgroup.Add(2 * len(urls))
    for _, url := range(urls){
        go getWebPage(url, fileDetailsChannel, &waitgroup)
        go writeFile(fileDetailsChannel, &waitgroup)
    }
    waitgroup.Wait()
    close(fileDetailsChannel)
    elapsed := time.Since(start)
    log.Printf("Total time %s", elapsed)
}


func getWebPage(url string, fileDetailsChannel chan fileDetails, waitgroup *sync.WaitGroup){
    
    response, err := http.Get(url)
    urlCompoents := strings.Split(url, "/")
    fileName := "htmls/" + urlCompoents[len(urlCompoents)-1] + ".html"
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
        fileDetailsChannel <- fileDetails{fileName: fileName, content: "Failed to load page"}
    }else {
        data, _ := ioutil.ReadAll(response.Body)
        fileDetailsChannel <- fileDetails{fileName: fileName, content: string(data)}
    }
    waitgroup.Done()
}

func writeFile(fileDetailsChannel chan fileDetails, waitgroup *sync.WaitGroup){
    newpath := filepath.Join(".", "htmls")
    os.MkdirAll(newpath, os.ModePerm)
    fileDetail, is_open := <- fileDetailsChannel
    fmt.Println(is_open)
    fileName := fileDetail.fileName
    data := fileDetail.content
    f, err := os.Create(fileName)
    if err != nil {
        fmt.Println(err)
    }
    _, err1 := f.WriteString(data)
    if err1 != nil {
        f.Close()
    }
    waitgroup.Done()

}


