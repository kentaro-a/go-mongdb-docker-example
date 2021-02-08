package templates

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetse976a941b4ea8409959a7e730ee5d0fa0b2e4e5e = "<!DOCTYPE html>\n<html lang=\"ja\">\n\t<head>\n\t    <meta charset=\"UTF-8\">\n\t\t<title>login</title>\n\t</head>\n\t<body>\n\t\t<div>\n\t\t\t<h1>Login</h1><br>\n\t\t\t<form method=\"POST\" action=\"/login\">\n\t\t\t\t<p>User Name: <input type=\"text\" name=\"name\" value=\"{{.Name}}\"></p>\n\t\t\t\t<p>Password: <input type=\"password\" name=\"pass\" value=\"{{.Pass}}\"></p>\n\t\t\t\t<input type=\"submit\" value=\"login\">\n\t\t\t</form>\n\n\t\t\t{{if gt (len .ErrorMessages) 0 }}\n\t\t\t\t<div style=\"background-color:red\" >\n\t\t\t\t\t{{range .ErrorMessages}}\n\t\t\t\t\t\t<p>{{.}}</p>\n\t\t\t\t\t{{end}}\n\t\t\t\t</div>\n\t\t\t{{end}}\n\t\t</div>\n\t</body>\n</html>\n"
var _Assets1081c5b10b8e22e502790bd42a6e9ed8487aa0c3 = "<!DOCTYPE html>\n<html lang=\"ja\">\n\t<head>\n\t    <meta charset=\"UTF-8\">\n\t\t<title>users</title>\n\t</head>\n\t<body>\n\t\t<div>\n\t\t\t<h1>HOME</h1><br>\n\t\t\t<p>Logged in as..</p>\n\t\t\t<p>User ID: {{.usr_id}}</p>\n\t\t\t<p>User Name: {{.usr_name}}</p>\n\t\t</div>\n\t\t<div>\n\t\t\t<a href=\"/logout\">logout</a>\n\t\t</div>\n\t</body>\n</html>\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"templates"}, "/templates": []string{"htmls"}, "/templates/htmls": []string{"login.html", "home.html"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1612765538, 1612765538624715188),
		Data:     nil,
	}, "/templates": &assets.File{
		Path:     "/templates",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1612764433, 1612764433386518569),
		Data:     nil,
	}, "/templates/htmls": &assets.File{
		Path:     "/templates/htmls",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1612766601, 1612766601033580293),
		Data:     nil,
	}, "/templates/htmls/login.html": &assets.File{
		Path:     "/templates/htmls/login.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1612766601, 1612766601032580191),
		Data:     []byte(_Assetse976a941b4ea8409959a7e730ee5d0fa0b2e4e5e),
	}, "/templates/htmls/home.html": &assets.File{
		Path:     "/templates/htmls/home.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1612766130, 1612766130712386446),
		Data:     []byte(_Assets1081c5b10b8e22e502790bd42a6e9ed8487aa0c3),
	}}, "")
