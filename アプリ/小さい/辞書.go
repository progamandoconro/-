package main

import (
        "fmt"
        "time"
	//"strings"
        "io/ioutil"
)

func main () {
f := "データー/JmnedictFurigana.txt"
d, err := ioutil.ReadFile(f)
 
if err != nil {
        fmt.Println("File reading error", err)
        return
    }

time.Sleep (1* time.Second)
fmt.Println("Contents of file:", string(d))


}

