package main
import (
    "bufio"
    "os"
)

const bufferCapacity = 1024*1024*4

var vsProjDir string

func main () {
    var err error
    if vsProjDir,err = os.Getwd() ; err != nil { return }

    stdout := bufio.NewWriterSize(os.Stdout, bufferCapacity)
    stdin := bufio.NewScanner(os.Stdin)
    stdin.Buffer(make([]byte, 1024*1024, bufferCapacity), bufferCapacity)
    for stdin.Scan() {
        if _,err = stdout.WriteString(handleRequest(stdin.Text()) + "\r\n") ; err != nil { break }
        if err = stdout.Flush() ; err != nil { break }
    }
}

func handleRequest (queryln string) (resultln string) {
    resultln = vsProjDir + "---not yet dudes! "
    return
}
