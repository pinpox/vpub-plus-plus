// Code generated by go generate; DO NOT EDIT.

package handler

var TplMap = map[string]string{
	"account": `{{ define "title" }}Account{{ end }}

{{ define "content" }}
<h1>Account</h1>

<form action="/update-account" method="post">
    {{ .csrfField }}
    <div class="field">
        <label for="picture">Picture</label>
        <input type="url" name="picture" id="picture" value="{{ .form.Picture }}">
    </div>
    <div class="field">
        <label for="about">About</label>
        <textarea name="about" id="about" autofocus>{{ .form.About }}</textarea>
    </div>
    <input type="submit" value="Submit">
</form>
</section>
{{ end }}
`,
	"admin": `{{ define "breadcrumb" }} > Admin{{ end }}
{{ define "content"}}
<h1>Admin</h1>
<nav>
  <ul>
    <li><a href="/admin/settings/edit">Edit settings</a></li>
    <li><a href="/admin/keys">Manage keys</a></li>
    <li><a href="/admin/boards">Manage boards</a></li>
    <li><a href="/admin/forums">Manage forums</a></li>
    <li><a href="/admin/users">Manage users</a></li>
  </ul>
</nav>
{{ end }}`,
	"admin_board": `{{ define "content"}}
<nav class="breadcrumb">
    <ul>
        <li>
            <a href="/admin">Admin</a>
            <ul>
                <li>
                    Boards
                </li>
            </ul>
        </li>
    </ul>
</nav>
<h1>Boards</h1>
<p>
    {{ if .hasForums }}
    <a href="/admin/boards/new">New board</a>
    {{ else }}
    <a href="/admin/forums">Create a forum</a> to create boards
    {{ end }}
</p>
<table>
    <thead>
    <tr>
        <th class="grow">Board</th>
        <th>Edit</th>
    </tr>
    </thead>
    <tbody>
    {{ if .forums }}
    {{ range .forums }}
    <tr class="forum">
        <td colspan="4">{{ .Name }}</td>
    </tr>
    {{ range .Boards }}
    <tr>
        <td colspan="grow">
            <a href="/boards/{{ .Id }}">{{ .Name }}</a><br>{{ .Description }}
        </td>
        <td class="center"><a href="/admin/boards/{{ .Id }}/edit">Edit</a></td>
    </tr>
    {{ end }}
    {{ end }}
    {{ else }}
    <tr>
        <td colspan="2">No boards yet.</td>
    </tr>
    {{ end }}
    </tbody>
</table>
{{ end }}`,
	"admin_board_create": `{{ define "title" }}New board{{ end }}
{{ define "content" }}
<nav class="breadcrumb">
  <ul>
    <li>
      <a href="/admin">Admin</a>
      <ul>
        <li>
          <a href="/admin/boards">Boards</a>
          <ul>
            <li>Create board</li>
          </ul>
        </li>
      </ul>
    </li>
  </ul>
</nav>
<h1>Create board</h1>
{{ if .errorMessage }}
    <p class="errors">{{ .errorMessage }}</p>
{{ end }}
<form action="/admin/boards/save" method="post">
  {{ .csrfField }}
  {{ template "board_form" .form }}
  <input type="submit" value="Submit">
</form>
{{ end }}
`,
	"admin_board_edit": `{{ define "title" }}Edit board{{ end }}
{{ define "content" }}
<nav class="breadcrumb">
    <ul>
        <li>
            <a href="{{ .root }}/admin">Admin</a>
            <ul>
                <li>
                    <a href="/admin/boards">Boards</a>
                    <ul>
                        <li>Edit board</li>
                    </ul>
                </li>
            </ul>
        </li>
    </ul>
</nav>
<h1>Edit board</h1>
{{ if .errorMessage }}
    <p class="errors">{{ .errorMessage }}</p>
{{ end }}
<form action="/admin/boards/{{ .board.Id }}/update" method="post">
    {{ .csrfField }}
    {{ template "board_form" .form }}
    <input type="submit" value="Submit">
</form>
{{ end }}
`,
	"admin_forum": `{{ define "content"}}
<nav class="breadcrumb">
    <ul>
        <li>
            <a href="/admin">Admin</a>
            <ul>
                <li>
                    Forums
                </li>
            </ul>
        </li>
    </ul>
</nav>
<h1>Forums</h1>
<p><a href="/admin/forums/new">New forum</a></p>
<table>
    <thead>
    <tr>
        <th class="grow">Forum</th>
        <th>Edit</th>
    </tr>
    </thead>
    <tbody>
    {{ if .forums }}
    {{ range .forums }}
    <tr>
        <td colspan="grow">
            {{ .Name }}
        </td>
        <td class="center"><a href="/admin/forums/{{ .Id }}/edit">Edit</a></td>
    </tr>
    {{ end }}
    {{ else }}
    <tr>
        <td colspan="2">No boards yet.</td>
    </tr>
    {{ end }}
    </tbody>
</table>
{{ end }}`,
	"admin_forum_create": `{{ define "title" }}New forum{{ end }}
{{ define "content" }}
<nav class="breadcrumb">
    <ul>
        <li>
            <a href="/admin">Admin</a>
            <ul>
                <li>
                    <a href="/admin/forums">Forums</a>
                    <ul>
                        <li>Create forum</li>
                    </ul>
                </li>
            </ul>
        </li>
    </ul>
</nav>
<h1>Create forum</h1>
{{ if .errorMessage }}
    <p class="errors">{{ .errorMessage }}</p>
{{ end }}
<form action="/admin/forums/save" method="post">
    {{ .csrfField }}
    {{ template "forum_form" .form }}
    <input type="submit" value="Submit">
</form>
{{ end }}
`,
	"admin_forum_edit": `{{ define "title" }}Edit forum{{ end }}
{{ define "content" }}
<nav class="breadcrumb">
    <ul>
        <li>
            <a href="/admin">Admin</a>
            <ul>
                <li>
                    <a href="/admin/forums">Forums</a>
                    <ul>
                        <li>Edit forum</li>
                    </ul>
                </li>
            </ul>
        </li>
    </ul>
</nav>
<h1>Edit forum</h1>
{{ if .errorMessage }}
    <p class="errors">{{ .errorMessage }}</p>
{{ end }}
<form action="/admin/forums/{{ .forum.Id }}/update" method="post">
    {{ .csrfField }}
    {{ template "forum_form" .form }}
    <input type="submit" value="Submit">
</form>
{{ end }}
`,
	"admin_keys": `{{ define "title" }}Keys{{ end }}
{{ define "content" }}
<nav class="breadcrumb">
    <ul>
        <li>
            <a href="/admin">Admin</a>
            <ul>
                <li>
                    Keys
                </li>
            </ul>
        </li>
    </ul>
</nav>
<h1>Keys</h1>
<form action="/admin/keys/save" method="post" class="action">
    {{ .csrfField }}
    <input type="submit" value="Create key">
</form>

<table>
    <thead>
    <tr>
        <th class="grow">Key</th>
        <th>Created</th>
        <th>Delete</th>
    </tr>
    </thead>
    <tbody>
    {{ range .keys }}
    <tr>
        <td colspan="grow">{{ .Key }}</td>
        <td class="center">{{ iso8601 .CreatedAt }}</td>
        <td class="center"><a href="/admin/keys/{{ .Id }}/remove">Delete</a></td>
    </tr>
    {{ end }}
    </tbody>
</table>
{{ end }}
`,
	"admin_settings_edit": `{{ define "breadcrumb" }} > <a href="/admin">Admin</a> > Settings{{ end }}
{{ define "content"}}
<nav class="breadcrumb">
    <ul>
        <li>
            <a href="/admin">Admin</a>
            <ul>
                <li>
                    Edit Settings
                </li>
            </ul>
        </li>
    </ul>
</nav>
<h1>Edit Settings</h1>
<form action="/admin/settings/update" method="post">
    {{ .csrfField }}
    <div class="field">
        <label for="name">Name</label>
        <input type="text" name="name" id="name" value="{{ .form.Name }}" autocomplete="off" maxlength="120" autofocus/>
    </div>
    <div class="field">
        <label for="name">URL</label>
        <input type="url" name="url" id="url" value="{{ .form.URL }}" autocomplete="off"/>
    </div>
    <div class="field">
        <label for="css">Footer</label>
        <textarea class="editor" name="footer" id="footer">{{ .form.Footer }}</textarea>
    </div>
    <div class="field">
        <label for="css">CSS</label>
        <textarea class="editor" name="css" id="css">{{ .form.Css }}</textarea>
    </div>
    <div class="field">
        <label for="per-page">Per page</label>
        <input type="number" name="per-page" id="per-page" value="{{ .form.PerPage }}" autocomplete="off" required/>
    </div>
    <input type="submit" value="Submit">
</form>
</table>
{{ end }}`,
	"admin_user": `{{ define "content"}}
<nav class="breadcrumb">
    <ul>
        <li>
            <a href="/admin">Admin</a>
            <ul>
                <li>Users</li>
            </ul>
        </li>
    </ul>
</nav>
<h1>Users</h1>
<table>
    <thead>
    <tr>
        <th class="grow">User</th>
        <th>Edit</th>
        <th>Password</th>
        <th>Delete</th>
    </tr>
    </thead>
    <tbody>
    {{ range .users }}
    <tr>
        <td colspan="grow">{{ .Name }}</td>
        <td class="center"><a href="/admin/users/{{ .Id }}/edit">Edit</a></td>
        <td class="center"><a href="/reset-password?hash={{ .Hash }}">Reset</a></td>
        <td class="center"><a href="/admin/users/{{ .Id }}/remove">Delete</a></td>
    </tr>
    {{ end }}
    </tbody>
</table>
{{ end }}`,
	"admin_user_edit": `{{ define "title" }}Edit user{{ end }}
{{ define "content" }}
<nav class="breadcrumb">
    <ul>
        <li>
            <a href="/admin">Admin</a>
            <ul>
                <li>
                    <a href="/admin/users">Users</a>
                    <ul>
                        <li>Edit users</li>
                    </ul>
                </li>
            </ul>
        </li>
    </ul>
</nav>
<h1>Edit user</h1>
<form action="/admin/users/{{ .user.Id }}/update" method="post">
    {{ .csrfField }}
    <div class="field">
        <label for="name">Name</label>
        <input type="text" name="name" id="name" value="{{ .form.Username }}" autocomplete="off" maxlength="120" required autofocus/>
    </div>
    <div class="field">
        <label for="name">Picture</label>
        <input type="text" name="picture" id="picture" value="{{ .form.Picture }}" autocomplete="off"/>
    </div>
    <div class="field">
        <label for="about">About</label>
        <textarea class="editor" name="about" id="about">{{ .form.About }}</textarea>
    </div>
    <input type="submit" value="Submit">
</form>
{{ end }}
`,
	"admin_user_remove": `{{ define "content" }}

{{ if .errorMessage }}
    <p class="errors">{{ .errorMessage }}</p>
{{ end }}

Are you sure you you want to delete the following user?
<p>{{ .user.Name }}</p>
<form action="/admin/users/{{ .user.Id }}/remove" method="post">
    {{ .csrfField }}
    <input type="submit" value="Submit">
</form>
{{ end }}`,
	"board": `{{ define "breadcrumb" }}<a href="/">boards</a> > {{ .board.Name }}{{ end }}
{{ define "content" }}
<nav class="breadcrumb">
    <ul>
        <li>
            <a href="/">All forums</a>
            <ul>
                <li>
                    <a href="/forums/{{ .board.Forum.Id }}">{{ .board.Forum.Name }}</a>
                    <ul>
                        <li>{{ .board.Name }}</li>
                    </ul>
                </li>
            </ul>
        </li>
    </ul>
</nav>
<h1>{{ .board.Name }}</h1>

{{ if logged }}
{{ if .board.Forum.IsLocked }}
<p>This forum is locked.</p>
{{ end }}
{{ if .board.IsLocked }}
<p>This board is locked.</p>
{{ end }}
{{ if or (and (not .board.IsLocked) (not .board.Forum.IsLocked)) .logged.IsAdmin }}
<form action="/boards/{{ .board.Id }}/new-topic" method="get" class="action">
    <input type="submit" value="New topic">
</form>
{{ end }}
{{ end }}

<section>
    <table>
        <thead>
        <tr>
            <th class="grow">Subject</th>
            <th>Author</th>
            <th>Replies</th>
            <th>Updated</th>
        </tr>
        </thead>
        <tbody>
        {{ if .topics }}
        {{ range .topics }}
        <tr{{ if .IsSticky }} class="sticky"{{ end }}>
            <td colspan="grow"><a href="/topics/{{ .Id }}">{{ .Post.Subject }}</a></td>
            <td class="center">{{ .Post.User.Name }}</td>
            <td class="center">{{ dec .Posts }}</td>
            <td><a href="/topics/{{ .Id }}/newest">{{ iso8601 .UpdatedAt }}</a></td>
        </tr>
        {{ end }}
        {{ else }}
        <tr>
            <td colspan="4">No topics yet.</td>
        </tr>
        {{ end }}
        </tbody>
    </table>
    {{ template "pagination" .pagination }}
</section>
{{ end }}`,
	"boards": `{{ define "content"}}
<nav class="breadcrumb">
    <ul>
        <li>
            <a href="/">All forums</a>
            <ul>
                <li>
                    {{ .forum.Name }}
                </li>
            </ul>
        </li>
    </ul>
</nav>
<h1>{{ .forum.Name }}</h1>
{{ if .forum.IsLocked }}
<p>This forum is locked.</p>
{{ end }}
<table>
    <thead>
    <tr>
        <th class="grow">Board</th>
        <th>Topics</th>
        <th>Posts</th>
        <th>Updated</th>
    </tr>
    </thead>
    <tbody>
    {{ if .boards }}
    {{ range .boards }}
    <tr>
        <td colspan="grow">
            <a href="/boards/{{ .Id }}">{{ .Name }}</a><br>{{ .Description }}
        </td>
        <td class="center">{{ .Topics }}</td>
        <td class="center">{{ .Posts }}</td>
        <td class="center"><a href="/boards/{{ .Id }}/newest">{{ iso8601 .UpdatedAt }}</a></td>
    </tr>
    {{ end }}
    {{ else }}
    <tr>
        <td colspan="4">No boards yet.</td>
    </tr>
    {{ end }}
    </tbody>
</table>
{{ end }}`,
	"confirm_remove_post": `{{ define "content" }}
    {{ if .errorMessage }}
        <p class="errors">{{ .errorMessage }}</p>
    {{ end }}

    Are you sure you you want to delete the following post?
    <p>{{ syntax .post.Content }}</p>
    <form action="/posts/{{ .post.Id }}/remove" method="post">
        {{ .csrfField }}
        <input type="submit" value="Submit">
    </form>
{{ end }}`,
	"create_post": `{{ define "title" }}Create Post{{ end }}
{{ define "content" }}
<nav class="breadcrumb">
  <ul>
    <li>
      <a href="/">All forums</a>
      <ul>
        <li>
          <a href="/forums/{{ .board.Forum.Id }}">{{ .board.Forum.Name }}</a>
          <ul>
            <li>
              <a href="/boards/{{ .board.Id }}">{{ .board.Name }}</a>
              <ul>
                <li>{{ .topic.Post.Subject }}</li>
              </ul>
            </li>
          </ul>
        </li>
      </ul>
    </li>
  </ul>
</nav>

<h1>Create Post</h1>

{{ if .errorMessage }}
    <p class="errors">{{ .errorMessage }}</p>
{{ end }}

<form action="/posts/save" method="post">
  {{ .csrfField }}
  {{ template "post_form" .form }}
  <input type="submit" value="Reply">
</form>
{{ end }}
`,
	"create_topic": `{{ define "title" }}New topic{{ end }}
{{ define "breadcrumb" }} > <a href="/boards/{{ .board.Id }}">{{ .board.Name }}</a>{{ end }}
{{ define "content" }}
<nav class="breadcrumb">
    <ul>
        <li>
            <a href="/">All forums</a>
            <ul>
                <li>
                    <a href="/forums/{{ .board.Forum.Id }}">{{ .board.Forum.Name }}</a>
                    <ul>
                        <li>
                            <a href="/boards/{{ .board.Id }}">{{ .board.Name }}</a>
                            <ul>
                                <li>New topic</li>
                            </ul>
                        </li>
                    </ul>
                </li>
            </ul>
        </li>
    </ul>
</nav>
<h1>New topic</h1>
{{ if .errorMessage }}
    <p class="errors">{{ .errorMessage }}</p>
{{ end }}
<form action="/boards/{{ .board.Id }}/save-topic" method="post">
    {{ .csrfField }}
    <input type="hidden" name="boardId" value="{{ .form.BoardId }}">
    {{ template "post_form" .form.PostForm }}
    {{ if .logged.IsAdmin }}
    {{ template "topic_form" .form }}
    {{ end }}
    <input type="submit" value="Submit">
</form>
{{ end }}
`,
	"edit_post": `{{ define "title" }}Edit Post{{ end }}
{{ define "content" }}
<nav class="breadcrumb">
    <ul>
        <li>
            <a href="/">All forums</a>
            <ul>
                <li>
                    <a href="/forums/{{ .board.Forum.Id }}">{{ .board.Forum.Name }}</a>
                    <ul>
                        <li>
                            <a href="/boards/{{ .board.Id }}">{{ .board.Name }}</a>
                            <ul>
                                <li>{{ .topic.Post.Subject }}</li>
                            </ul>
                        </li>
                    </ul>
                </li>
            </ul>
        </li>
    </ul>
</nav>

<h1>Edit Post</h1>

{{ if .errorMessage }}
    <p class="errors">{{ .errorMessage }}</p>
{{ end }}

<form action="/posts/{{ .post.Id }}/update" method="post">
    {{ .csrfField }}
    {{ template "post_form" .form }}
    <input type="submit" value="Reply">
</form>
{{ end }}
`,
	"edit_topic": `{{ define "title" }}Edit topic{{ end }}
{{ define "content" }}
<h1>Edit topic</h1>
{{ if .errorMessage }}
    <p class="errors">{{ .errorMessage }}</p>
{{ end }}
<form action="/topics/{{ .form.Id }}/update" method="post">
    {{ .csrfField }}
    <input type="hidden" name="boardId" value="{{ .form.BoardId }}">
    {{ template "post_form" .form.PostForm }}
    {{ if .logged.IsAdmin }}
    {{ template "topic_form" .form }}
    {{ end }}
    <input type="submit" value="Submit">
</form>
{{ end }}
`,
	"index": `{{ define "content"}}
{{ if .settings.Name }}
<h1>{{ .settings.Name }}</h1>
{{ else }}
<h1>Forums</h1>
{{ end }}
<table>
    <thead>
        <tr>
            <th class="grow">Forum</th>
            <th>Topics</th>
            <th>Posts</th>
            <th>Updated</th>
        </tr>
    </thead>
    <tbody>
        {{ if .forums }}
        {{ range .forums }}
        <tr class="forum">
            <td colspan="4"><a href="/forums/{{ .Id }}">{{ .Name }}</a></td>
        </tr>
        {{ range .Boards }}
        <tr>
            <td colspan="grow">
                <a href="/boards/{{ .Id }}">{{ .Name }}</a><br>{{ .Description }}
            </td>
            <td class="center">{{ .Topics }}</td>
            <td class="center">{{ .Posts }}</td>
            <td class="center">
                {{ if .Topics }}
                <a href="/boards/{{ .Id }}/newest">{{ iso8601 .UpdatedAt }}</a>
                {{ else }}
                {{ iso8601 .UpdatedAt }}
                {{ end }}
            </td>
        </tr>
        {{ end }}
        {{ end }}
        {{ else }}
        <tr>
            <td colspan="4">No forums yet.</td>
        </tr>
        {{ end }}
    </tbody>
</table>
{{ end }}`,
	"login": `{{ define "title" }}Login{{ end }}

{{ define "content" }}
<h1>Login</h1>
{{ if .errorMessage }}
    <p class="errors">{{ .errorMessage }}</p>
{{ end }}
<form action="/login" method="post" class="auth-form">
    {{ .csrfField }}
    <div class="field">
        <label for="name">Username</label>
        <input type="text" name="name" id="name" autocomplete="off" required/>
    </div>
    <div class="field">
        <label for="password">Password</label>
        <input type="password" name="password" id="password" required/>
    </div>
    <input type="submit" value="Login">
</form>
{{ end }}`,
	"post": `{{ define "breadcrumb" }} > <a href="/topics/{{ .post.Topic }}">{{ .post.Topic }}</a>{{ end }}
{{ define "content"}}
<h1>{{ .post.Subject }}</h1>
<table class="thread">
    <tr class="post">
        <td class="post-aside">
            <p>{{ .post.User }}</p>
            <p>{{ timeAgo .post.CreatedAt }}</p>
        </td>
        <td class="post-content">
            {{ syntax .content }}
        </td>
    </tr>
    {{ range .replies }}
    <tr class="post">
        <td class="post-aside">
            <p>{{ .User }}</p>
            <p>{{ timeAgo .CreatedAt }}</p>
        </td>
        <td class="post-content">
            {{ syntax .Content }}
        </td>
    </tr>
    {{ end }}
</table>
<form action="/posts/{{ .post.Id }}/reply" method="post">
    {{ .csrfField }}
    <div class="field">
        <textarea name="reply"></textarea>
    </div>
    <input type="submit" value="Reply">
</form>
{{ end }}`,
	"posts": `{{ define "content" }}
<h1>Posts</h1>

<section>
    <table>
        <thead>
        <tr>
            <th class="grow">Subject</th>
            <th>Author</th>
            <th>Posted</th>
        </tr>
        </thead>
        <tbody>
        {{ if .posts }}
        {{ range .posts }}
        <tr>
        <td colspan="grow"><a href="/topics/{{ .TopicId }}#{{ .Id }}">{{ .Subject }}</a></td>
        <td class="center">{{ .User.Name }}</td>
        <td class="center">{{ iso8601 .CreatedAt }}</td>
        </tr>
        {{ end }}
        {{ else }}
        <tr>
            <td colspan="4">No posts yet.</td>
        </tr>
        {{ end }}
        </tbody>
    </table>
    {{ template "pagination" .pagination }}
</section>
{{ end }}`,
	"register": `{{ define "title" }}Register{{ end }}

{{ define "content" }}
<h1>Register</h1>
{{ if .errorMessage }}
    <p class="errors">{{ .errorMessage }}</p>
{{ end }}
<form action="/register" method="post" class="auth-form">
    {{ .csrfField }}
    <div class="field">
        <label for="name">Username</label>
        <input type="text" id="name" name="name" autocomplete="off" value="{{ .form.Username }}" maxlength="15" required/>
    </div>
    <div class="field">
        <label for="password">Password</label>
        <input type="password" id="password" name="password" required/>
    </div>
    <div class="field">
        <label for="confirm">Confirm password</label>
        <input type="password" id="confirm" name="confirm" required/>
    </div>
    <div class="field">
        <label for="key">Key</label>
        <input type="text" id="key" name="key" required/>
    </div>
    <input type="submit" value="Submit">
</form>
{{ end }}
`,
	"reset_password": `{{ define "title" }}Reset password{{ end }}

{{ define "content" }}
<h1>Reset password</h1>
{{ if .error }}
<p class="error">{{ .error }}</p>
{{ end }}
<form action="/reset-password" method="post" class="auth-form">
    {{ .csrfField }}
    <input name="hash" type="hidden" value="{{ .hash }}">
    <div class="field">
        <label for="password">New password</label>
        <input type="password" id="password" name="password"/>
    </div>
    <div class="field">
        <label for="confirm">Confirm password</label>
        <input type="password" id="confirm" name="confirm" required/>
    </div>
    <input type="submit" value="Submit">
</form>
{{ end }}`,
	"topic": `{{ define "breadcrumb" }}<a href="/">boards</a> > <a href="/boards/{{ .board.Id }}">{{ .board.Name }}</a>{{ end }}
{{ define "content"}}
<nav class="breadcrumb">
    <ul>
        <li>
            <a href="/">All forums</a>
            <ul>
                <li>
                    <a href="/forums/{{ .board.Forum.Id }}">{{ .board.Forum.Name }}</a>
                    <ul>
                        <li>
                            <a href="/boards/{{ .board.Id }}">{{ .board.Name }}</a>
                            <ul>
                                <li>{{ .topic.Post.Subject }}</li>
                            </ul>
                        </li>
                    </ul>
                </li>
            </ul>
        </li>
    </ul>
</nav>
<h1>{{ .topic.Post.Subject }}</h1>
<table class="topic">
    <thead>
        <th>Author</th>
        <th>Topic</th>
    </thead>
    <tbody>
    {{ range .posts }}
    <tr id="{{ .Id }}">
        <td class="col-author">
            {{ .User.Name }}
            <p><img src="{{ .User.Picture }}" width="80"/></p>
        </td>
        <td>
            <div class="posted">
                Posted {{ iso8601Time .CreatedAt }}
                {{ if and (eq $.topic.Post.Id .Id) $.logged.IsAdmin }}
                <a href="/topics/{{ $.topic.Id }}/edit">edit</a> <a href="/posts/{{ .Id }}/remove">remove</a>
                {{ else }}
                {{ if or (hasPermission .User.Name) $.logged.IsAdmin }}
                <a href="/posts/{{ .Id }}/edit">edit</a> <a href="/posts/{{ .Id }}/remove">remove</a>
                {{ end }}
                {{ end }}
                <hr/>
            </div>
            <div>{{ syntax .Content }}</div>
            {{ if .User.About }}
            <div class="signature">
                {{ sig .User.About }}
            </div>
            {{ end }}
        </td>
    </tr>
    {{ end }}
    </tbody>
</table>
{{ if .board.Forum.IsLocked }}
<p>This forum is locked.</p>
{{ end }}
{{ if .board.IsLocked }}
<p>This board is locked.</p>
{{ end }}
{{ if .topic.IsLocked }}
<p>This topic is locked.</p>
{{ end }}
{{ if logged }}
{{ if or (and (not .topic.IsLocked) (not .board.IsLocked) (not .board.Forum.IsLocked)) .logged.IsAdmin }}
<section style="margin-top: 1em;">
    {{ if .errorMessage }}
        <p class="errors">{{ .errorMessage }}</p>
    {{ end }}
    <form action="/posts/save" method="post">
        {{ .csrfField }}
        <input type="hidden" name="topicId" value="{{ .topic.Id }}">
        <input type="hidden" name="boardId" value="{{ .board.Id }}">
        <input type="hidden" name="subject" value="Re: {{ .topic.Post.Subject }}">
        <div class="field">
            <label for="content">Reply to this topic</label>
            <textarea name="content" id="content" style="height: 150px;"></textarea>
        </div>
        <input type="submit" value="Reply">
    </form>
</section>
{{ end }}
{{ end }}
{{ end }}`,
}
