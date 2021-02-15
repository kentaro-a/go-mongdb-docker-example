package templates

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetse976a941b4ea8409959a7e730ee5d0fa0b2e4e5e = "{{define \"content\"}}\n<div>\n\t<h1>Login</h1><br>\n\t<form method=\"POST\" action=\"/login\">\n\t\t<p>User Name: <input type=\"text\" name=\"name\" value=\"{{.Data.Name}}\"></p>\n\t\t<p>Password: <input type=\"password\" name=\"pass\" value=\"{{.Data.Pass}}\"></p>\n\t\t<input type=\"submit\" value=\"login\">\n\t</form>\n\n\t{{if gt (len .Data.ErrorMessages) 0 }}\n\t<div style=\"background-color:red\" >\n\t\t{{range .Data.ErrorMessages}}\n\t\t<p>{{.}}</p>\n\t\t{{end}}\n\t</div>\n\t{{end}}\n</div>\n{{end}}\n"
var _Assetse189cfd5cc1aefc7cba4f446e8fd1399488e39a8 = "{{define \"layout\"}}\n\n<!DOCTYPE html>\n<html lang=\"ja\">\n\t<head>\n\t    <meta charset=\"UTF-8\">\n\t\t<title>{{.Title}}</title>\n\t</head>\n\t<body>\n\t\t<div class=\"container\">\n\t\t\t{{template \"content\" .}}\t\n\t\t</div>\n\t</body>\n</html>\n{{end}}\n"
var _Assets1081c5b10b8e22e502790bd42a6e9ed8487aa0c3 = "{{define \"content\"}}\n<div>\n\t<h1>HOME</h1><br>\n\t<p>Logged in as..</p>\n\t<p>User ID: {{.Data.usr_id}}</p>\n\t<p>User Name: {{.Data.usr_name}}</p>\n</div>\n<div>\n\t<a href=\"/logout\">logout</a>\n</div>\n{{end}}\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"templates"}, "/templates": []string{"htmls"}, "/templates/htmls": []string{"login.html", "layout.html", "home.html"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1613365852, 1613365852929303997),
		Data:     nil,
	}, "/templates": &assets.File{
		Path:     "/templates",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1613370528, 1613370528266014623),
		Data:     nil,
	}, "/templates/htmls": &assets.File{
		Path:     "/templates/htmls",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1613370696, 1613370696774887203),
		Data:     nil,
	}, "/templates/htmls/login.html": &assets.File{
		Path:     "/templates/htmls/login.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1613370696, 1613370696773887223),
		Data:     []byte(_Assetse976a941b4ea8409959a7e730ee5d0fa0b2e4e5e),
	}, "/templates/htmls/layout.html": &assets.File{
		Path:     "/templates/htmls/layout.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1613370154, 1613370154019962906),
		Data:     []byte(_Assetse189cfd5cc1aefc7cba4f446e8fd1399488e39a8),
	}, "/templates/htmls/home.html": &assets.File{
		Path:     "/templates/htmls/home.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1613370324, 1613370324933435906),
		Data:     []byte(_Assets1081c5b10b8e22e502790bd42a6e9ed8487aa0c3),
	}}, "")
