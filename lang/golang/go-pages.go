package zgo

import (
	"net/url"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/metaleap/go-util/dev/go"
	"github.com/metaleap/go-util/fs"
	"github.com/metaleap/go-util/run"
	"github.com/metaleap/go-util/str"
	"github.com/metaleap/zentient"
)

var (
	pages goPages
)

func init() {
	pages.Impl, z.Lang.Pages = &pages, &pages
}

type goPages struct {
	z.PagesBase
}

func (me *goPages) PageBodyInnerHtml(rawUri string, path []string, query url.Values, fragment string) string {
	if len(path) > 1 {
		switch path[0] {
		case "godoc":
			subpath := strings.Join(path[1:], "/")
			return me.onGoDoc(subpath, fragment)
		}
	}
	return me.PagesBase.PageBodyInnerHtml(rawUri, path, query, fragment)
}

func (*goPages) linkifyUri(uri string) string {
	return z.Strf("command:zen.internal.page?\"%s\"", template.URLQueryEscaper(uri))
}

func (me *goPages) onGoDoc(uriPath string, identName string) string {
	cmdout, cmderr, err := urun.CmdExec("godoc", "-url", uriPath)
	if err != nil {
		return err.Error()
	} else if cmdout != "" {
		i := strings.Index(cmdout, "<div id=\"footer\">")
		if i > 0 {
			cmdout = cmdout[:i]
		}
		if i = strings.Index(cmdout, "<div id=\"short-nav\">"); i >= 0 {
			cmdout = cmdout[i:]
		}
		left, right := "", cmdout
		for right != "" {
			if i = strings.Index(right, " href=\""); i < 0 {
				left, right = left+right, ""
			} else if j := strings.Index(right[i+7:], "\""); j < 0 {
				left, right = left+right, ""
			} else {
				href := right[i+7:][:j]
				if strings.HasPrefix(href, "/") {
					var done bool
					if strings.HasPrefix(href, "/src/") {
						if u, e := url.Parse(href); e == nil {
							for _, gp := range udevgo.AllGoPaths() {
								if ln, fp := "", filepath.Join(gp, u.Path); ufs.FileExists(fp) {
									if strings.HasPrefix(u.Fragment, "L") {
										if l := ustr.ParseInt(u.Fragment[1:]); l > 0 {
											ln = z.Strf(":%d", l+10)
										}
									}
									done, href = true, z.Strf("command:zen.internal.openFileAt?\"%s%s\"", fp, ln)
									break
								}
							}
						}
					}
					if !done {
						href = "zentient://" + z.Lang.ID + "/godoc" + href
						href = me.linkifyUri(href)
					}
				}
				left, right = left+right[:i]+" href='"+href, "'"+right[i+7:][j+1:]
			}
		}
		cmdout = left
		return _PAGES_GODOC_CSS + cmdout
	}
	return cmderr
}

const _PAGES_GODOC_CSS = `
<style type="text/css">
div.collapsed { display: none !important; }
</style>
`
