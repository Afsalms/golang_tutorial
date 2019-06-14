

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
    var waitgroup2 sync.WaitGroup
    
    // urlChannel := make(chan string, 100)
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

    // for _, url := range(urls){
    //     urlChannel <- url

    // }
    // close(urlChannel)
    waitgroup.Add(len(urls))
    for _, url := range(urls){
        go getWebPage(url, fileDetailsChannel, &waitgroup)
    }
    waitgroup.Wait()
    close(fileDetailsChannel)
    elapsed1 := time.Since(start)
    log.Printf("Fetching time %s", elapsed1)
    waitgroup2.Add(10)
    go writeFile(fileDetailsChannel, &waitgroup2)
    go writeFile(fileDetailsChannel, &waitgroup2)
    go writeFile(fileDetailsChannel, &waitgroup2)
    go writeFile(fileDetailsChannel, &waitgroup2)
    go writeFile(fileDetailsChannel, &waitgroup2)
    go writeFile(fileDetailsChannel, &waitgroup2)
    go writeFile(fileDetailsChannel, &waitgroup2)
    go writeFile(fileDetailsChannel, &waitgroup2)
    go writeFile(fileDetailsChannel, &waitgroup2)
    go writeFile(fileDetailsChannel, &waitgroup2)
    waitgroup2.Wait()
    elapsed := time.Since(start)
    log.Printf("Total time %s", elapsed)
}


func getWebPage(url string, fileDetailsChannel chan fileDetails, waitgroup *sync.WaitGroup){
    
    response, err := http.Get(url)
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    }else {
        data, _ := ioutil.ReadAll(response.Body)
        urlCompoents := strings.Split(url, "/")
        fileName := "htmls/" + urlCompoents[len(urlCompoents)-1] + ".html"
        fileDetailsChannel <- fileDetails{fileName: fileName, content: string(data)}
    }
    // close(fileDetailsChannel)
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


