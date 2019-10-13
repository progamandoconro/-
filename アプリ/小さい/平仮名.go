package main

import ( "fmt"
	"github.com/gojp/kana")

func main () {
for {
hiragana ()
}}

func hiragana () {
var message string

     fmt.Println("Ingresa ROMANJI a transformar a HIRAGANA\n")
     fmt.Scanln(&message)
     k := kana.RomajiToHiragana(message)
     fmt.Println(message + " en Hiragana se escribe:\n")
     fmt.Println(k)
     fmt.Println( "\n")

}



