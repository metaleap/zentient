package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main () {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        txt := scanner.Text()
        fmt.Println(strings.ToUpper(txt) + strings.ToLower(txt))
    }
    if err := scanner.Err() ; err != nil {
        print(err)
    }
}
