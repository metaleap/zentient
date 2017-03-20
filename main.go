package main
import (
    "bufio"
    "os"
    "path/filepath"
    "strings"

    "github.com/metaleap/go-util-fs"

    "github.com/metaleap/zentient/z"
)

const bufferCapacity = 1024*1024*4

var (
    vsProjDir   string
    dataDir     string
)


func main () {
    var err error
    if vsProjDir,err = os.Getwd() ; err != nil { return }
    if len(os.Args) != 2 { return }

    vn := filepath.VolumeName(vsProjDir)
    subdir := strings.Replace(vsProjDir, vn, ufs.SanitizeFsName(vn), -1)
    dataDir = filepath.Join(os.Args[1], subdir)
    if err = ufs.EnsureDirExists(dataDir) ; err != nil { return }

    stdout := bufio.NewWriterSize(os.Stdout, bufferCapacity)
    stdin := bufio.NewScanner(os.Stdin)
    stdin.Buffer(make([]byte, 1024*1024, bufferCapacity), bufferCapacity)
    for stdin.Scan() {
        if response := handleRequest(stdin.Text()) ; len(response) > 0 {
            if _,err = stdout.WriteString(response + "\r\n") ; err != nil { break }
            if err = stdout.Flush() ; err != nil { break }
        }
    }
}

func handleRequest (queryln string) (resultln string) {
    resultln = vsProjDir + "::" + z.CMD_FILE_OPEN + "---not yet dudes! "
    return
}
