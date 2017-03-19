package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main () {


    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer([]byte{}, 1024*1024*128)
    for scanner.Scan() {
        txt := scanner.Text()
        fmt.Println(strings.ToUpper(txt) + strings.ToLower(txt))
    }
    if err := scanner.Err() ; err != nil {
        print(err)
    }
}
