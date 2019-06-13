

package main

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "io"
    "log"
    "os"
    // "io/ioutil"
    "net/http"
    // "strings"
    // "path/filepath"
    "time"
    "sync"
)

func main() {
    var waitgroup sync.WaitGroup
    start := time.Now()
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

    waitgroup.Add(len(urls))
    for _, url := range(urls){
        go getWebPage(url, &waitgroup)
    }
    waitgroup.Wait()
    elapsed := time.Since(start)
    log.Printf("Binomial took %s", elapsed)
}


func getWebPage(url string ,waitgroup *sync.WaitGroup){
    _, err := http.Get(url)
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    // } else {
    //     data, _ := ioutil.ReadAll(response.Body)
        // fmt.Println(data)
        // urlCompoents := strings.Split(url, "/")
        // newpath := filepath.Join(".", "htmls")
        // os.MkdirAll(newpath, os.ModePerm)
        // fileName := "htmls/" + urlCompoents[len(urlCompoents)-1] + ".html"
        // f, err := os.Create(fileName)
        // if err != nil {
            // fmt.Println(err)
            // return
        // }
        // _, err1 := f.WriteString(string(data))
        // if err1 != nil {
            // f.Close()
            // return
        // }
    }
    waitgroup.Done()

}