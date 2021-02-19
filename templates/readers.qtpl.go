// Code generated by qtc from "readers.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/readers.qtpl:1
package templates

//line templates/readers.qtpl:1
import "net/http"

//line templates/readers.qtpl:2
import "path"

//line templates/readers.qtpl:3
import "github.com/bouncepaw/mycorrhiza/user"

//line templates/readers.qtpl:4
import "github.com/bouncepaw/mycorrhiza/util"

//line templates/readers.qtpl:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/readers.qtpl:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/readers.qtpl:6
func StreamHistoryHTML(qw422016 *qt422016.Writer, rq *http.Request, hyphaName, list string) {
//line templates/readers.qtpl:6
	qw422016.N().S(`
`)
//line templates/readers.qtpl:7
	streamnavHTML(qw422016, rq, hyphaName, "history")
//line templates/readers.qtpl:7
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<article class="history">
		<h1>History of `)
//line templates/readers.qtpl:11
	qw422016.E().S(hyphaName)
//line templates/readers.qtpl:11
	qw422016.N().S(`</h1>
		`)
//line templates/readers.qtpl:12
	qw422016.N().S(list)
//line templates/readers.qtpl:12
	qw422016.N().S(`
	</article>
</main>
</div>
`)
//line templates/readers.qtpl:16
}

//line templates/readers.qtpl:16
func WriteHistoryHTML(qq422016 qtio422016.Writer, rq *http.Request, hyphaName, list string) {
//line templates/readers.qtpl:16
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/readers.qtpl:16
	StreamHistoryHTML(qw422016, rq, hyphaName, list)
//line templates/readers.qtpl:16
	qt422016.ReleaseWriter(qw422016)
//line templates/readers.qtpl:16
}

//line templates/readers.qtpl:16
func HistoryHTML(rq *http.Request, hyphaName, list string) string {
//line templates/readers.qtpl:16
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/readers.qtpl:16
	WriteHistoryHTML(qb422016, rq, hyphaName, list)
//line templates/readers.qtpl:16
	qs422016 := string(qb422016.B)
//line templates/readers.qtpl:16
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/readers.qtpl:16
	return qs422016
//line templates/readers.qtpl:16
}

//line templates/readers.qtpl:18
func StreamRevisionHTML(qw422016 *qt422016.Writer, rq *http.Request, hyphaName, naviTitle, contents, relatives, revHash string) {
//line templates/readers.qtpl:18
	qw422016.N().S(`
`)
//line templates/readers.qtpl:19
	streamnavHTML(qw422016, rq, hyphaName, "revision", revHash)
//line templates/readers.qtpl:19
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<article>
		<p>Please note that viewing binary parts of hyphae is not supported in history for now.</p>
		`)
//line templates/readers.qtpl:24
	qw422016.N().S(naviTitle)
//line templates/readers.qtpl:24
	qw422016.N().S(`
		`)
//line templates/readers.qtpl:25
	qw422016.N().S(contents)
//line templates/readers.qtpl:25
	qw422016.N().S(`
	</article>
</main>
`)
//line templates/readers.qtpl:28
	streamrelativeHyphae(qw422016, relatives)
//line templates/readers.qtpl:28
	qw422016.N().S(`
</div>
`)
//line templates/readers.qtpl:30
}

//line templates/readers.qtpl:30
func WriteRevisionHTML(qq422016 qtio422016.Writer, rq *http.Request, hyphaName, naviTitle, contents, relatives, revHash string) {
//line templates/readers.qtpl:30
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/readers.qtpl:30
	StreamRevisionHTML(qw422016, rq, hyphaName, naviTitle, contents, relatives, revHash)
//line templates/readers.qtpl:30
	qt422016.ReleaseWriter(qw422016)
//line templates/readers.qtpl:30
}

//line templates/readers.qtpl:30
func RevisionHTML(rq *http.Request, hyphaName, naviTitle, contents, relatives, revHash string) string {
//line templates/readers.qtpl:30
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/readers.qtpl:30
	WriteRevisionHTML(qb422016, rq, hyphaName, naviTitle, contents, relatives, revHash)
//line templates/readers.qtpl:30
	qs422016 := string(qb422016.B)
//line templates/readers.qtpl:30
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/readers.qtpl:30
	return qs422016
//line templates/readers.qtpl:30
}

// If `contents` == "", a helpful message is shown instead.

//line templates/readers.qtpl:33
func StreamPageHTML(qw422016 *qt422016.Writer, rq *http.Request, hyphaName, naviTitle, contents, relatives, backlinkEntries, prevHyphaName, nextHyphaName string, hasAmnt bool) {
//line templates/readers.qtpl:33
	qw422016.N().S(`
`)
//line templates/readers.qtpl:34
	streamnavHTML(qw422016, rq, hyphaName, "page")
//line templates/readers.qtpl:34
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<article>
		`)
//line templates/readers.qtpl:38
	qw422016.N().S(naviTitle)
//line templates/readers.qtpl:38
	qw422016.N().S(`
		`)
//line templates/readers.qtpl:39
	if contents == "" {
//line templates/readers.qtpl:39
		qw422016.N().S(`
			<p>This hypha has no text. Why not <a href="/edit/`)
//line templates/readers.qtpl:40
		qw422016.E().S(hyphaName)
//line templates/readers.qtpl:40
		qw422016.N().S(`">create it</a>?</p>
		`)
//line templates/readers.qtpl:41
	} else {
//line templates/readers.qtpl:41
		qw422016.N().S(`
			`)
//line templates/readers.qtpl:42
		qw422016.N().S(contents)
//line templates/readers.qtpl:42
		qw422016.N().S(`
		`)
//line templates/readers.qtpl:43
	}
//line templates/readers.qtpl:43
	qw422016.N().S(`
	</article>
	<section class="prevnext">
		`)
//line templates/readers.qtpl:46
	if prevHyphaName != "" {
//line templates/readers.qtpl:46
		qw422016.N().S(`
		<a class="prevnext__el prevnext__prev" href="/hypha/`)
//line templates/readers.qtpl:47
		qw422016.E().S(prevHyphaName)
//line templates/readers.qtpl:47
		qw422016.N().S(`" rel="prev">← `)
//line templates/readers.qtpl:47
		qw422016.E().S(util.BeautifulName(path.Base(prevHyphaName)))
//line templates/readers.qtpl:47
		qw422016.N().S(`</a>
		`)
//line templates/readers.qtpl:48
	}
//line templates/readers.qtpl:48
	qw422016.N().S(`
		`)
//line templates/readers.qtpl:49
	if nextHyphaName != "" {
//line templates/readers.qtpl:49
		qw422016.N().S(`
		<a class="prevnext__el prevnext__next" href="/hypha/`)
//line templates/readers.qtpl:50
		qw422016.E().S(nextHyphaName)
//line templates/readers.qtpl:50
		qw422016.N().S(`" rel="next">`)
//line templates/readers.qtpl:50
		qw422016.E().S(util.BeautifulName(path.Base(nextHyphaName)))
//line templates/readers.qtpl:50
		qw422016.N().S(` →</a>
		`)
//line templates/readers.qtpl:51
	}
//line templates/readers.qtpl:51
	qw422016.N().S(`
	</section>
`)
//line templates/readers.qtpl:53
	if u := user.FromRequest(rq); !user.AuthUsed || u.Group != "anon" {
//line templates/readers.qtpl:53
		qw422016.N().S(`
	<form action="/upload-binary/`)
//line templates/readers.qtpl:54
		qw422016.E().S(hyphaName)
//line templates/readers.qtpl:54
		qw422016.N().S(`"
			method="post" enctype="multipart/form-data"
			class="upload-amnt">
		`)
//line templates/readers.qtpl:57
		if hasAmnt {
//line templates/readers.qtpl:57
			qw422016.N().S(`
			<a class="upload-amnt__unattach" href="/unattach-ask/`)
//line templates/readers.qtpl:58
			qw422016.E().S(hyphaName)
//line templates/readers.qtpl:58
			qw422016.N().S(`">Unattach current attachment?</a>
		`)
//line templates/readers.qtpl:59
		}
//line templates/readers.qtpl:59
		qw422016.N().S(`
		<label for="upload-binary__input">Upload a new attachment</label>
		<br>
		<input type="file" id="upload-binary__input" name="binary"/>
		<input type="submit"/>
	</form>
`)
//line templates/readers.qtpl:65
	}
//line templates/readers.qtpl:65
	qw422016.N().S(`
</main>
`)
//line templates/readers.qtpl:67
	streamrelativeHyphae(qw422016, relatives)
//line templates/readers.qtpl:67
	qw422016.N().S(`
`)
//line templates/readers.qtpl:68
	streambacklinks(qw422016, backlinkEntries)
//line templates/readers.qtpl:68
	qw422016.N().S(`
</div>
`)
//line templates/readers.qtpl:70
}

//line templates/readers.qtpl:70
func WritePageHTML(qq422016 qtio422016.Writer, rq *http.Request, hyphaName, naviTitle, contents, relatives, backlinkEntries, prevHyphaName, nextHyphaName string, hasAmnt bool) {
//line templates/readers.qtpl:70
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/readers.qtpl:70
	StreamPageHTML(qw422016, rq, hyphaName, naviTitle, contents, relatives, backlinkEntries, prevHyphaName, nextHyphaName, hasAmnt)
//line templates/readers.qtpl:70
	qt422016.ReleaseWriter(qw422016)
//line templates/readers.qtpl:70
}

//line templates/readers.qtpl:70
func PageHTML(rq *http.Request, hyphaName, naviTitle, contents, relatives, backlinkEntries, prevHyphaName, nextHyphaName string, hasAmnt bool) string {
//line templates/readers.qtpl:70
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/readers.qtpl:70
	WritePageHTML(qb422016, rq, hyphaName, naviTitle, contents, relatives, backlinkEntries, prevHyphaName, nextHyphaName, hasAmnt)
//line templates/readers.qtpl:70
	qs422016 := string(qb422016.B)
//line templates/readers.qtpl:70
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/readers.qtpl:70
	return qs422016
//line templates/readers.qtpl:70
}
