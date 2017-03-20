package main
import (
    "bufio"
    "encoding/json"
    "os"
    "path/filepath"
    "strings"

    "github.com/metaleap/go-util-fs"
    "github.com/metaleap/go-util-misc"

    "github.com/metaleap/zentient/z"
    "github.com/metaleap/zentient/go"
    "github.com/metaleap/zentient/hs"
)

const bufferCapacity = 1024*1024*4


func main () {
    var err error

    if z.ProjDir,err = os.Getwd() ; err != nil { return }
    if err = ensureDataDirs() ; err != nil { return }
    z.Zengines = map[string]z.Zengine { "hs": zhs.Zengine, "go": zgo.Zengine }

    stdin := bufio.NewScanner(os.Stdin)
    stdin.Buffer(make([]byte, 1024*1024, bufferCapacity), bufferCapacity)
    stdout := bufio.NewWriterSize(os.Stdout, bufferCapacity)
    z.Out = json.NewEncoder(stdout)
    z.Out.SetEscapeHTML(false)
    z.Out.SetIndent("","")

    for stdin.Scan() {
        if err = z.HandleRequest(stdin.Text()) ; err == nil {
            err = stdout.Flush()
        }
        if err != nil {
            z.Out.Encode(err.Error())
            err = stdout.Flush()
            break
        }
    }
}


func ensureDataDirs() error {
    var basedir, subdir string

    //  coming from VScode?
    if len(os.Args) > 1 && len(os.Args[1])>0 {
        const sep = string(os.PathSeparator)
        if editordatadir , index := os.Args[1] , strings.Index(os.Args[1], sep + "Code" + sep) ; index > 0 {
            basedir = editordatadir[0 : index]
        }
    }
    //  otherwise..
    if len(basedir) == 0 || !ufs.DirExists(basedir) {
        basedir = ugo.UserDataDirPath()
    }

    z.DataDir = filepath.Join(basedir, "zentient")
    //  cache for this session is in datadir/user/proj/dir
    if volname := filepath.VolumeName(z.ProjDir) ; len(volname) > 0 {
        subdir = strings.Replace(z.ProjDir, volname, ufs.SanitizeFsName(volname), -1)
    } else {
        subdir = z.ProjDir
    }
    z.DataProjDir = filepath.Join(z.DataDir, subdir)
    return ufs.EnsureDirExists(z.DataProjDir) //  this also creates dataDir
}
