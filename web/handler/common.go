// Code generated by go generate; DO NOT EDIT.

package handler

var TplCommonMap = map[string]string{
	"layout": `{{ define "layout" }}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="/style.css"/>
    <title>{{ .boardTitle }}</title>
    {{ template "head" . }}
</head>
<body>
    <header>
        <span><a href="/">{{ .boardTitle }}</a></span>
        <nav>
            {{ if .logged }}
            {{ if .hasNotifications }}<a href="/notifications" class="notifications">New replies</a> {{ end }} {{ .logged }} (<a href="/logout">logout</a>)
            {{ else }}
            <a href="/login">login</a> <a href="/register">register</a>
            {{ end }}
        </nav>
    </header>
    {{ template "content" . }}
    <footer><p>Powered by vpub</p></footer>
</body>
</html>
{{ end }}
{{ define "head" }}{{ end }}`,
	"meta": `{{ define "meta" }}
<a href="/~{{ .User }}">{{ .User }}</a> on {{ .Date }}{{ if .Topic }} in <a href="/topics/{{ .Topic }}">{{ .Topic }}</a>{{ end }}
{{ end }}`,
	"post_form": `{{ define "post_form" }}
    {{ if .Topics }}
    <select name="topic">
        {{ range .Topics }}
        <option value="{{ . }}" {{ if eq . $.Topic }}selected{{ end }}>{{ . }}</option>
        {{ end }}
    </select>
    {{ end }}

    <label for="title">Title</label>
    <input type="text" name="title" id="title" value="{{ .Title }}" autocomplete="off" required autofocus/>
    <textarea class="editor" name="content" id="content" required>{{ .Content }}</textarea>
    <br>
{{ end }}`,
	"posts": `{{ define "posts" }}
<ol class="posts">
    {{ if . }}
    {{ range . }}
    <li>
        <h2><a href="/posts/{{ .Id }}">{{ .Title }}</a></h2>
        <ul class="key-value">
            <li><span class="key">From: </span><span class="value"><a href="/~{{ .User }}">{{ .User }}</a></span></li>
            <li><span class="key">On: </span><span class="value">{{ .Date }}</span></li>
            {{ if .Topic }}<li><span class="key">Topic: </span><span class="value"><a href="/topics/{{ .Topic }}">{{ .Topic }}</a></span></li>{{ end }}
            <li><span class="key">Replies: </span><span class="value">{{ .Replies }}</span></li>
        </ul>
    </li>
    {{ end }}
    {{ else }}
    <li>No post yet</li>
    {{ end }}
</ol>
{{ end }}`,
	"reply": `{{ define "reply" }}
    {{ if . }}
    <ol class="replies">
    {{ range . }}
        <li>
            <details class="reply" open>
                <summary><a href="/~{{ .User }}">{{ .User }}</a> on {{ .Date }}</summary>
                <div class="reply-content">{{ gmi2html .Content }}</div>
                {{ if logged }}
                <footer>
                    <a href="/replies/{{ .Id }}">reply</a>
                    {{ if hasPermission .User }}
                    <a href="/replies/{{ .Id }}/edit">edit</a>
                    <a href="/replies/{{ .Id }}/remove">Remove</a>
                    {{ end }}
                </footer>
                {{ end }}
                {{ if .Thread }}
                <div class="thread">
                    {{ template "reply" .Thread }}
                </div>
                {{ end }}
            </details>
        </li>
    {{ end }}
    </ol>
    {{ end }}
{{ end }}`,
	"topics": `{{ define "topics" }}
{{ if .topics }}
<nav class="topics">
  {{ range .topics }}
  {{ if .Selected }}
  <span class="selected">{{ .Name }}</span>
  {{ else }}
  <a href="/topics/{{ .Name }}">{{ .Name }}</a>
  {{ end }}
  {{ end }}
</nav>
{{ end }}
{{ end }}`,
}
