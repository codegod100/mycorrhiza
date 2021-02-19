// Code generated by qtc from "common.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/common.qtpl:1
package templates

//line templates/common.qtpl:1
import "net/http"

//line templates/common.qtpl:2
import "github.com/bouncepaw/mycorrhiza/user"

//line templates/common.qtpl:3
import "github.com/bouncepaw/mycorrhiza/util"

// This is the <nav> seen on top of many pages.

//line templates/common.qtpl:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/common.qtpl:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/common.qtpl:7
type navEntry struct {
	path  string
	title string
}

var navEntries = []navEntry{
	{"page", "Hypha"},
	{"edit", "Edit"},
	{"history", "History"},
	{"revision", "NOT REACHED"},
	{"rename-ask", "Rename"},
	{"delete-ask", "Delete"},
	{"text", "Raw text"},
}

//line templates/common.qtpl:22
func streamnavHTML(qw422016 *qt422016.Writer, rq *http.Request, hyphaName, navType string, revisionHash ...string) {
//line templates/common.qtpl:22
	qw422016.N().S(`
`)
//line templates/common.qtpl:24
	u := user.FromRequest(rq)

//line templates/common.qtpl:25
	qw422016.N().S(`
	<nav class="hypha-tabs main-width">
		<ul class="hypha-tabs__flex">
`)
//line templates/common.qtpl:28
	for _, entry := range navEntries {
//line templates/common.qtpl:29
		if navType == "revision" && entry.path == "revision" {
//line templates/common.qtpl:29
			qw422016.N().S(`			<li class="hypha-tabs__tab hypha-tabs__tab_active">
				<span class="hypha-tabs__selection">`)
//line templates/common.qtpl:31
			qw422016.E().S(revisionHash[0])
//line templates/common.qtpl:31
			qw422016.N().S(`</span>
			</li>
`)
//line templates/common.qtpl:33
		} else if navType == entry.path {
//line templates/common.qtpl:33
			qw422016.N().S(`			<li class="hypha-tabs__tab hypha-tabs__tab_active">
				<span class="hypha-tabs__selection">`)
//line templates/common.qtpl:35
			qw422016.E().S(entry.title)
//line templates/common.qtpl:35
			qw422016.N().S(`</span>
			</li>
`)
//line templates/common.qtpl:37
		} else if entry.path != "revision" && u.CanProceed(entry.path) {
//line templates/common.qtpl:37
			qw422016.N().S(`			<li class="hypha-tabs__tab">
				<a class="hypha-tabs__link" href="/`)
//line templates/common.qtpl:39
			qw422016.E().S(entry.path)
//line templates/common.qtpl:39
			qw422016.N().S(`/`)
//line templates/common.qtpl:39
			qw422016.E().S(hyphaName)
//line templates/common.qtpl:39
			qw422016.N().S(`">`)
//line templates/common.qtpl:39
			qw422016.E().S(entry.title)
//line templates/common.qtpl:39
			qw422016.N().S(`</a>
			</li>
`)
//line templates/common.qtpl:41
		}
//line templates/common.qtpl:42
	}
//line templates/common.qtpl:42
	qw422016.N().S(`		</ul>
	</nav>
`)
//line templates/common.qtpl:45
}

//line templates/common.qtpl:45
func writenavHTML(qq422016 qtio422016.Writer, rq *http.Request, hyphaName, navType string, revisionHash ...string) {
//line templates/common.qtpl:45
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/common.qtpl:45
	streamnavHTML(qw422016, rq, hyphaName, navType, revisionHash...)
//line templates/common.qtpl:45
	qt422016.ReleaseWriter(qw422016)
//line templates/common.qtpl:45
}

//line templates/common.qtpl:45
func navHTML(rq *http.Request, hyphaName, navType string, revisionHash ...string) string {
//line templates/common.qtpl:45
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/common.qtpl:45
	writenavHTML(qb422016, rq, hyphaName, navType, revisionHash...)
//line templates/common.qtpl:45
	qs422016 := string(qb422016.B)
//line templates/common.qtpl:45
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/common.qtpl:45
	return qs422016
//line templates/common.qtpl:45
}

//line templates/common.qtpl:47
func streamuserMenuHTML(qw422016 *qt422016.Writer, u *user.User) {
//line templates/common.qtpl:47
	qw422016.N().S(`
`)
//line templates/common.qtpl:48
	if user.AuthUsed {
//line templates/common.qtpl:48
		qw422016.N().S(`
<li class="header-links__entry header-links__entry_user">
	`)
//line templates/common.qtpl:50
		if u.Group == "anon" {
//line templates/common.qtpl:50
			qw422016.N().S(`
	<a href="/login" class="header-links__link">Login</a>
	`)
//line templates/common.qtpl:52
		} else {
//line templates/common.qtpl:52
			qw422016.N().S(`
	<a href="/page/`)
//line templates/common.qtpl:53
			qw422016.E().S(util.UserHypha)
//line templates/common.qtpl:53
			qw422016.N().S(`/`)
//line templates/common.qtpl:53
			qw422016.E().S(u.Name)
//line templates/common.qtpl:53
			qw422016.N().S(`" class="header-links__link">`)
//line templates/common.qtpl:53
			qw422016.E().S(u.Name)
//line templates/common.qtpl:53
			qw422016.N().S(`</a>
	`)
//line templates/common.qtpl:54
		}
//line templates/common.qtpl:54
		qw422016.N().S(`
</li>
`)
//line templates/common.qtpl:56
	}
//line templates/common.qtpl:56
	qw422016.N().S(`
`)
//line templates/common.qtpl:57
}

//line templates/common.qtpl:57
func writeuserMenuHTML(qq422016 qtio422016.Writer, u *user.User) {
//line templates/common.qtpl:57
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/common.qtpl:57
	streamuserMenuHTML(qw422016, u)
//line templates/common.qtpl:57
	qt422016.ReleaseWriter(qw422016)
//line templates/common.qtpl:57
}

//line templates/common.qtpl:57
func userMenuHTML(u *user.User) string {
//line templates/common.qtpl:57
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/common.qtpl:57
	writeuserMenuHTML(qb422016, u)
//line templates/common.qtpl:57
	qs422016 := string(qb422016.B)
//line templates/common.qtpl:57
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/common.qtpl:57
	return qs422016
//line templates/common.qtpl:57
}

//line templates/common.qtpl:59
func streamrelativeHyphae(qw422016 *qt422016.Writer, relatives string) {
//line templates/common.qtpl:59
	qw422016.N().S(`
<aside class="relative-hyphae layout-card">
	<h2 class="relative-hyphae__title layout-card__title">Relative hyphae</h2>
	`)
//line templates/common.qtpl:62
	qw422016.N().S(relatives)
//line templates/common.qtpl:62
	qw422016.N().S(`
</aside>
`)
//line templates/common.qtpl:64
}

//line templates/common.qtpl:64
func writerelativeHyphae(qq422016 qtio422016.Writer, relatives string) {
//line templates/common.qtpl:64
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/common.qtpl:64
	streamrelativeHyphae(qw422016, relatives)
//line templates/common.qtpl:64
	qt422016.ReleaseWriter(qw422016)
//line templates/common.qtpl:64
}

//line templates/common.qtpl:64
func relativeHyphae(relatives string) string {
//line templates/common.qtpl:64
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/common.qtpl:64
	writerelativeHyphae(qb422016, relatives)
//line templates/common.qtpl:64
	qs422016 := string(qb422016.B)
//line templates/common.qtpl:64
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/common.qtpl:64
	return qs422016
//line templates/common.qtpl:64
}

//line templates/common.qtpl:66
func streambacklinks(qw422016 *qt422016.Writer, backlinkEntries string) {
//line templates/common.qtpl:66
	qw422016.N().S(`
<aside class="backlinks layout-card">
	<h2 class="backlinks__title layout-card__title">Backlinks</h2>
	<nav class="backlinks__nav">
		<ul class="backlinks__list">
		`)
//line templates/common.qtpl:71
	qw422016.N().S(backlinkEntries)
//line templates/common.qtpl:71
	qw422016.N().S(`
		</ul>
	</nav>
</aside>
`)
//line templates/common.qtpl:75
}

//line templates/common.qtpl:75
func writebacklinks(qq422016 qtio422016.Writer, backlinkEntries string) {
//line templates/common.qtpl:75
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/common.qtpl:75
	streambacklinks(qw422016, backlinkEntries)
//line templates/common.qtpl:75
	qt422016.ReleaseWriter(qw422016)
//line templates/common.qtpl:75
}

//line templates/common.qtpl:75
func backlinks(backlinkEntries string) string {
//line templates/common.qtpl:75
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/common.qtpl:75
	writebacklinks(qb422016, backlinkEntries)
//line templates/common.qtpl:75
	qs422016 := string(qb422016.B)
//line templates/common.qtpl:75
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/common.qtpl:75
	return qs422016
//line templates/common.qtpl:75
}
