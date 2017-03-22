package z
import (
)


type ProjInfo struct {
    SrcDir      string
    CacheDir    string
    ConfigDir   string
}


type Zengine interface {
    Ids () []string
    Jsonish () interface{}
    OnFileActive (*File)
    OnFileOpen (*File)
    OnFileWrite (*File)
}


var (
    Proj        = &ProjInfo{}
    AllFiles    = map[string]*File {}
    Zengines    = map[string]Zengine {}
)


func fromZidMsg (msgargs string) (z Zengine, argstr string) {
    zid := msgargs[:2]
    if z = Zengines[zid] ; z != nil {
        argstr = msgargs[3:]
    }
    return
}

func OnFileActive (file* File) {
    file.Z.OnFileActive(file)
}

func onFileOpen (z Zengine, relpath string) {
    file := AllFiles[relpath]
    if file == nil {
        file = NewFile(z, relpath)
        AllFiles[relpath] = file
        z.OnFileOpen(file)
    }
    OnFileActive(file)
}

func onFileWrite (z Zengine, relpath string) {
    file := AllFiles[relpath]
    if (file == nil) {
        onFileOpen(z, relpath)
        file = AllFiles[relpath]
    }
    z.OnFileWrite(file)
}
