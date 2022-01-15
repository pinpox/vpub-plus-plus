// Code generated by go generate; DO NOT EDIT.

package handler

var TplCommonMap = map[string]string{
	"board_form": `{{ define "board_form" }}
<div class="field">
    <label for="name">Name</label>
    <input type="text" name="name" id="name" value="{{ .Name }}" autocomplete="off" maxlength="120" required autofocus/>
</div>
<div class="field">
    <label for="description">Description</label>
    <textarea class="editor" name="description" id="description" required>{{ .Description }}</textarea>
</div>
{{ end }}`,
	"layout": `{{ define "layout" }}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="/style.css"/>
    <title>{{ .settings.Name }}</title>
    {{ template "head" . }}
</head>
<body>
    <header>
        <span><a href="/">{{ .settings.Name }}</a></span>
        <nav>
            {{ if logged }}
            {{ if .hasNotifications }}<a href="/notifications" class="notifications">New replies</a> {{ end }} <a href="/account">{{ .logged.Name }}</a> (<a href="/logout">logout</a>)
            {{ else }}
            <a href="/login">login</a> <a href="/register">register</a>
            {{ end }}
        </nav>
    </header>
<!--    <p>{{ template "breadcrumb" . }}</p>-->
    {{ template "content" . }}
</body>
</html>
{{ end }}
{{ define "head" }}{{ end }}
{{ define "breadcrumb" }}{{ end }}`,
	"post_form": `{{ define "post_form" }}
<input type="hidden" name="topicId" value="{{ .TopicId }}">
<div class="field">
    <label for="subject">Subject</label>
    <input type="text" name="subject" id="subject" value="{{ .Subject }}" autocomplete="off" maxlength="115" required autofocus/>
</div>
<div class="field">
    <label for="content">Content</label>
    <textarea class="editor" name="content" id="content" required>{{ .Content }}</textarea>
</div>
{{ end }}`,
	"posts": `{{ define "posts" }}
<!--{{ if . }}-->
<!--<ol class="posts">-->
<!--    {{ if . }}-->
<!--    {{ range . }}-->
<!--    <li>-->
<!--        <article>-->
<!--            <header><h2><a href="/posts/{{ .Id }}">{{ .Subject }}</a></h2> ({{ .Replies }})</header>-->
<!--            <div><a href="/~{{ .User }}">{{ .User }}</a>{{ if .Topic }} in <a href="/topics/{{ .Topic }}">{{ .Topic }}</a>{{ end }} {{ timeAgo .CreatedAt }}</div>-->
<!--        </article>-->
<!--    </li>-->
<!--    {{ end }}-->
<!--    {{ else }}-->
<!--    <li>No post yet</li>-->
<!--    {{ end }}-->
<!--</ol>-->
<!--{{ end }}-->

{{ if . }}
<table class="posts">
    <thead>
        <tr>
            <th>Subject</th>
            <th>Author</th>
            <th>Replies</th>
            <th>Updated</th>
        </tr>
    </thead>
    <tbody>
    {{ range . }}
    <tr>
        <td>
            <h2><a href="/posts/{{ .Id }}">{{ .Title }}</a></h2>
        </td>
        <td style="text-align: center;"><a href="/~{{ .User }}">{{ .User }}</a></td>
        <td style="text-align: center;">{{ .Replies }}</td>
        <td style="text-align: center">{{ .DateUpdated }}</td>
    </tr>
    {{ end }}
    </tbody>
</table>
{{ else }}
<p>No post yet</p>
{{ end }}
{{ end }}`,
	"postsTopic": `{{ define "postsTopic" }}
{{ if . }}
<table class="posts">
    <thead>
    <tr>
        <th>Topic</th>
        <th>Subject</th>
        <th>Author</th>
        <th>Replies</th>
        <th>Updated</th>
    </tr>
    </thead>
    <tbody>
    {{ range . }}
    <tr>
        <td style="text-align: center;"><a href="/topics/{{ .Topic }}">{{ .Topic }}</a></td>
        <td>
            <h2><a href="/posts/{{ .Id }}">{{ .Title }}</a></h2>
        </td>
        <td style="text-align: center;"><a href="/~{{ .User }}">{{ .User }}</a></td>
        <td style="text-align: center;">{{ .Replies }}</td>
        <td style="text-align: center">{{ .DateUpdated }}</td>
    </tr>
    {{ end }}
    </tbody>
</table>
{{ else }}
<p>No post yet</p>
{{ end }}
{{ end }}`,
	"reply": `{{ define "reply" }}
    {{ if . }}
    <ol class="replies">
    {{ range . }}
        <li>
            <details class="reply" open>
                <summary><a href="/~{{ .User }}">{{ .User }}</a> on {{ .Date }}</summary>
                <div class="content">{{ gmi2html .Content }}</div>
                {{ if logged }}
                <footer>
                    {{ if hasPermission .User }}
                    <a href="/replies/{{ .Id }}/edit">edit</a>
                    <a href="/replies/{{ .Id }}/remove">remove</a>
                    {{ end }}
                </footer>
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
