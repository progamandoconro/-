
package main

import (
    "bufio"
    "fmt"
    "os"
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

func main() {
    // Loop over lines in file.
    for index, line := range LinesInFile(`C:\programs\file.txt`) {
        fmt.Printf("Index = %v, line = %v\n", index , line)
    }

    // Get count of lines.
    lines := LinesInFile("データー/JmnedictFurigana.txt")
    fmt.Println(lines)
}
