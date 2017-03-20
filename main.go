package main
import (
    "bufio"
    "os"
    "path/filepath"
    "strings"

    "github.com/metaleap/go-util-fs"
    "github.com/metaleap/go-util-misc"

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
    if err = ensureDataDir() ; err != nil { return }

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


func ensureDataDir() error {
    var basedir, subdir string
    const sep = string(os.PathSeparator)
    if len(os.Args) > 1 && len(os.Args[1])>0 {
        if editordatadir , index := os.Args[1] , strings.Index(os.Args[1], sep + "Code" + sep) ; index > 0 {
            basedir = editordatadir[0 : index]
        }
    }
    if len(basedir) == 0 || !ufs.DirExists(basedir) {
        basedir = ugo.UserDataDirPath()
    }
    if volname := filepath.VolumeName(vsProjDir) ; len(volname) > 0 {
        subdir = strings.Replace(vsProjDir, volname, ufs.SanitizeFsName(volname), -1)
    } else {
        subdir = vsProjDir
    }
    dataDir = filepath.Join(basedir, "zentient", subdir)
    return ufs.EnsureDirExists(dataDir)
}


func handleRequest (queryln string) (resultln string) {
    resultln = dataDir + "::" + os.Args[0] + "::" + z.CMD_FILE_OPEN + queryln
    return
}
