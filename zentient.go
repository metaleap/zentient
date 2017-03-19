package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "github.com/metaleap/go-util-str"
)

func main () {
    lo := false
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        lo = !lo  ;  txt := scanner.Text()
        fmt.Println( ustr.Ifs(lo, strings.ToLower(txt), strings.ToUpper(txt)) )
    }
    if err := scanner.Err() ; err != nil {
        print(err)
    }
}
