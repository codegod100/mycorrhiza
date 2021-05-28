// Code generated by qtc from "history.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/history.qtpl:1
package views

//line views/history.qtpl:1
import "net/http"

//line views/history.qtpl:3
import "github.com/bouncepaw/mycorrhiza/util"

//line views/history.qtpl:4
import "github.com/bouncepaw/mycorrhiza/user"

//line views/history.qtpl:5
import "github.com/bouncepaw/mycorrhiza/hyphae"

//line views/history.qtpl:6
import "github.com/bouncepaw/mycorrhiza/history"

//line views/history.qtpl:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/history.qtpl:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/history.qtpl:9
func StreamPrimitiveDiffHTML(qw422016 *qt422016.Writer, rq *http.Request, h *hyphae.Hypha, u *user.User, hash string) {
//line views/history.qtpl:9
	qw422016.N().S(`
`)
//line views/history.qtpl:11
	text, err := history.PrimitiveDiffAtRevision(h.TextPath, hash)
	if err != nil {
		text = err.Error()
	}

//line views/history.qtpl:15
	qw422016.N().S(`
`)
//line views/history.qtpl:16
	StreamNavHTML(qw422016, rq, h.Name, "history")
//line views/history.qtpl:16
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<article>
		<h1>Diff `)
//line views/history.qtpl:20
	qw422016.E().S(util.BeautifulName(h.Name))
//line views/history.qtpl:20
	qw422016.N().S(` at `)
//line views/history.qtpl:20
	qw422016.E().S(hash)
//line views/history.qtpl:20
	qw422016.N().S(`</h1>
		<pre class="codeblock"><code>`)
//line views/history.qtpl:21
	qw422016.E().S(text)
//line views/history.qtpl:21
	qw422016.N().S(`</code></pre>
	</article>
</main>
</div>
`)
//line views/history.qtpl:25
}

//line views/history.qtpl:25
func WritePrimitiveDiffHTML(qq422016 qtio422016.Writer, rq *http.Request, h *hyphae.Hypha, u *user.User, hash string) {
//line views/history.qtpl:25
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/history.qtpl:25
	StreamPrimitiveDiffHTML(qw422016, rq, h, u, hash)
//line views/history.qtpl:25
	qt422016.ReleaseWriter(qw422016)
//line views/history.qtpl:25
}

//line views/history.qtpl:25
func PrimitiveDiffHTML(rq *http.Request, h *hyphae.Hypha, u *user.User, hash string) string {
//line views/history.qtpl:25
	qb422016 := qt422016.AcquireByteBuffer()
//line views/history.qtpl:25
	WritePrimitiveDiffHTML(qb422016, rq, h, u, hash)
//line views/history.qtpl:25
	qs422016 := string(qb422016.B)
//line views/history.qtpl:25
	qt422016.ReleaseByteBuffer(qb422016)
//line views/history.qtpl:25
	return qs422016
//line views/history.qtpl:25
}

//line views/history.qtpl:27
func StreamRecentChangesHTML(qw422016 *qt422016.Writer, n int) {
//line views/history.qtpl:27
	qw422016.N().S(`
<div class="layout">
<main class="main-width recent-changes">
	<h1>Recent Changes</h1>

	<nav class="recent-changes__count">
		See 
	`)
//line views/history.qtpl:34
	for _, m := range []int{20, 0, 50, 0, 100} {
//line views/history.qtpl:34
		qw422016.N().S(`
	`)
//line views/history.qtpl:35
		switch m {
//line views/history.qtpl:36
		case 0:
//line views/history.qtpl:36
			qw422016.N().S(`
		<span aria-hidden="true">|</span>
		`)
//line views/history.qtpl:38
		case n:
//line views/history.qtpl:38
			qw422016.N().S(`
		<b>`)
//line views/history.qtpl:39
			qw422016.N().D(n)
//line views/history.qtpl:39
			qw422016.N().S(`</b>
		`)
//line views/history.qtpl:40
		default:
//line views/history.qtpl:40
			qw422016.N().S(`
		<a href="/recent-changes/`)
//line views/history.qtpl:41
			qw422016.N().D(m)
//line views/history.qtpl:41
			qw422016.N().S(`">`)
//line views/history.qtpl:41
			qw422016.N().D(m)
//line views/history.qtpl:41
			qw422016.N().S(`</a>
	`)
//line views/history.qtpl:42
		}
//line views/history.qtpl:42
		qw422016.N().S(`
	`)
//line views/history.qtpl:43
	}
//line views/history.qtpl:43
	qw422016.N().S(`
		recent changes
	</nav>

	<p><img class="icon" width="20" height="20" src="/static/icon/feed">Subscribe via <a href="/recent-changes-rss">RSS</a>, <a href="/recent-changes-atom">Atom</a> or <a href="/recent-changes-json">JSON feed</a>.</p>

	`)
//line views/history.qtpl:54
	qw422016.N().S(`

	`)
//line views/history.qtpl:57
	changes := history.RecentChanges(n)

//line views/history.qtpl:58
	qw422016.N().S(`
	<section class="recent-changes__list" role="feed">
	`)
//line views/history.qtpl:60
	if len(changes) == 0 {
//line views/history.qtpl:60
		qw422016.N().S(`
		<p>Could not find any recent changes.</p>
	`)
//line views/history.qtpl:62
	} else {
//line views/history.qtpl:62
		qw422016.N().S(`
		`)
//line views/history.qtpl:63
		for i, entry := range changes {
//line views/history.qtpl:63
			qw422016.N().S(`
		<ul class="recent-changes__entry rc-entry" role="article"
		    aria-setsize="`)
//line views/history.qtpl:65
			qw422016.N().D(n)
//line views/history.qtpl:65
			qw422016.N().S(`" aria-posinset="`)
//line views/history.qtpl:65
			qw422016.N().D(i)
//line views/history.qtpl:65
			qw422016.N().S(`">
			 `)
//line views/history.qtpl:66
			qw422016.N().S(recentChangesEntry(entry))
//line views/history.qtpl:66
			qw422016.N().S(`
		</ul>
		`)
//line views/history.qtpl:68
		}
//line views/history.qtpl:68
		qw422016.N().S(`
	`)
//line views/history.qtpl:69
	}
//line views/history.qtpl:69
	qw422016.N().S(`
	</section>
</main>
</div>
`)
//line views/history.qtpl:73
}

//line views/history.qtpl:73
func WriteRecentChangesHTML(qq422016 qtio422016.Writer, n int) {
//line views/history.qtpl:73
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/history.qtpl:73
	StreamRecentChangesHTML(qw422016, n)
//line views/history.qtpl:73
	qt422016.ReleaseWriter(qw422016)
//line views/history.qtpl:73
}

//line views/history.qtpl:73
func RecentChangesHTML(n int) string {
//line views/history.qtpl:73
	qb422016 := qt422016.AcquireByteBuffer()
//line views/history.qtpl:73
	WriteRecentChangesHTML(qb422016, n)
//line views/history.qtpl:73
	qs422016 := string(qb422016.B)
//line views/history.qtpl:73
	qt422016.ReleaseByteBuffer(qb422016)
//line views/history.qtpl:73
	return qs422016
//line views/history.qtpl:73
}

//line views/history.qtpl:75
func streamrecentChangesEntry(qw422016 *qt422016.Writer, rev history.Revision) {
//line views/history.qtpl:75
	qw422016.N().S(`
<li class="rc-entry__time"><time>`)
//line views/history.qtpl:76
	qw422016.E().S(rev.TimeString())
//line views/history.qtpl:76
	qw422016.N().S(`</time></li>
<li class="rc-entry__hash">`)
//line views/history.qtpl:77
	qw422016.E().S(rev.Hash)
//line views/history.qtpl:77
	qw422016.N().S(`</li>
<li class="rc-entry__links">`)
//line views/history.qtpl:78
	qw422016.N().S(rev.HyphaeLinksHTML())
//line views/history.qtpl:78
	qw422016.N().S(`</li>
<li class="rc-entry__msg">`)
//line views/history.qtpl:79
	qw422016.E().S(rev.Message)
//line views/history.qtpl:79
	qw422016.N().S(` `)
//line views/history.qtpl:79
	if rev.Username != "anon" {
//line views/history.qtpl:79
		qw422016.N().S(`<span class="rc-entry__author">by <a href="/hypha/`)
//line views/history.qtpl:79
		qw422016.E().S(util.UserHypha)
//line views/history.qtpl:79
		qw422016.N().S(`/`)
//line views/history.qtpl:79
		qw422016.E().S(rev.Username)
//line views/history.qtpl:79
		qw422016.N().S(`" rel="author">`)
//line views/history.qtpl:79
		qw422016.E().S(rev.Username)
//line views/history.qtpl:79
		qw422016.N().S(`</a></span>`)
//line views/history.qtpl:79
	}
//line views/history.qtpl:79
	qw422016.N().S(`</li>
`)
//line views/history.qtpl:80
}

//line views/history.qtpl:80
func writerecentChangesEntry(qq422016 qtio422016.Writer, rev history.Revision) {
//line views/history.qtpl:80
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/history.qtpl:80
	streamrecentChangesEntry(qw422016, rev)
//line views/history.qtpl:80
	qt422016.ReleaseWriter(qw422016)
//line views/history.qtpl:80
}

//line views/history.qtpl:80
func recentChangesEntry(rev history.Revision) string {
//line views/history.qtpl:80
	qb422016 := qt422016.AcquireByteBuffer()
//line views/history.qtpl:80
	writerecentChangesEntry(qb422016, rev)
//line views/history.qtpl:80
	qs422016 := string(qb422016.B)
//line views/history.qtpl:80
	qt422016.ReleaseByteBuffer(qb422016)
//line views/history.qtpl:80
	return qs422016
//line views/history.qtpl:80
}

//line views/history.qtpl:82
func StreamHistoryHTML(qw422016 *qt422016.Writer, rq *http.Request, hyphaName, list string) {
//line views/history.qtpl:82
	qw422016.N().S(`
`)
//line views/history.qtpl:83
	StreamNavHTML(qw422016, rq, hyphaName, "history")
//line views/history.qtpl:83
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<article class="history">
		<h1>History of `)
//line views/history.qtpl:87
	qw422016.E().S(util.BeautifulName(hyphaName))
//line views/history.qtpl:87
	qw422016.N().S(`</h1>
		`)
//line views/history.qtpl:88
	qw422016.N().S(list)
//line views/history.qtpl:88
	qw422016.N().S(`
	</article>
</main>
</div>
`)
//line views/history.qtpl:92
}

//line views/history.qtpl:92
func WriteHistoryHTML(qq422016 qtio422016.Writer, rq *http.Request, hyphaName, list string) {
//line views/history.qtpl:92
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/history.qtpl:92
	StreamHistoryHTML(qw422016, rq, hyphaName, list)
//line views/history.qtpl:92
	qt422016.ReleaseWriter(qw422016)
//line views/history.qtpl:92
}

//line views/history.qtpl:92
func HistoryHTML(rq *http.Request, hyphaName, list string) string {
//line views/history.qtpl:92
	qb422016 := qt422016.AcquireByteBuffer()
//line views/history.qtpl:92
	WriteHistoryHTML(qb422016, rq, hyphaName, list)
//line views/history.qtpl:92
	qs422016 := string(qb422016.B)
//line views/history.qtpl:92
	qt422016.ReleaseByteBuffer(qb422016)
//line views/history.qtpl:92
	return qs422016
//line views/history.qtpl:92
}