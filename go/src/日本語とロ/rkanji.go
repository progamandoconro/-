package main

import( 

	"fmt"
	"os"
	"log"
	"io/ioutil"
	)	


func main() {
	file, err  := os.Open("kanji.txt")
	
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()


  b, err := ioutil.ReadAll(file)
  fmt.Print(b)

}


