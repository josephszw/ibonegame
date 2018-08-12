<h1>Edit {{.Post.Name}}</h1>

<h1>New Blog Post</h1>
<form action="" method="post">
昵称:<input type="text" name="title" value="{{.Post.Nick_name}}"><br>
名字：<textarea name="content" colspan="3" rowspan="10">{{.Post.Name}}</textarea>
<input type="hidden" name="id" value="{{.Post.Player_id}}">
<input type="submit">
</form>