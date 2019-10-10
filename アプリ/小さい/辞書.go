package main

import (
    "bufio"
    "fmt"
    "os"
    "time"
    "math/rand"
)

func LinesInFile(fileName string) []string {
    f, _ := os.Open(fileName)
    // Create new Scanner.
    scanner := bufio.NewScanner(f)
    result := []string{}
    // Use Scan.
    for scanner.Scan() {
        line := scanner.Text()
        // Append line to result.
        result = append(result, line)
    }
    return result
}

func ranDic () { 
    
    lines := LinesInFile("データー/JmnedictFurigana.txt")

        var NLines int
        fmt.Print("Cantidad de Kanji (漢字の量), filas estudiar del diccionario:       ")
        fmt.Scanln(&NLines)

     for i := 1; i <= NLines ; i++ {
                s1 := rand.NewSource(time.Now().UnixNano())
                r1 := rand.New(s1)
                r := r1.Intn(6394)      
                fmt.Println(lines[r])           
                time.Sleep(60 * time.Second)

      }
}

func main() {
 ranDic()
}




