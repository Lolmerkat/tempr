package main

import templating "github.com/lolmerkat/tempr/internal/app/tempr/templating/types"

func getTestTemplate() templating.Template {
	return templating.Template {
		Name: "HTML5",
		Author: "lolmerkat",
		Version: "1.0",
		Languages: []string{ "html", "javascript", "css" },
		Content: []templating.FSElement{
			&templating.File{
				Name: "index.html",
				Content: []string{
					"<!DOCTYPE html>",
					"<html>",
					"<head>",
					"\t<title>MyNewProject</title>",
					"</head>",
					"<body>",
					"\t",
					"</body>",
					"</html>",
				},
			},
			&templating.Directory{
				Name: "src",
				Children: []templating.FSElement{
					&templating.Directory{
						Name: "assets",
						Children: []templating.FSElement{},
					},
					&templating.Directory{
						Name: "scripts",
						Children: []templating.FSElement{
							&templating.File{
								Name: "main.js",
								Content: []string{},
							},
						},
					},
					&templating.Directory{
						Name: "styles",
						Children: []templating.FSElement{
							&templating.File{
								Name: "main.css",
								Content: []string{
									"body {",
									"\tmargin: 0;",
									"\tbackground-color: #262740;",
									"}",
								},
							},
							&templating.File{
								Name: "colors.css",
								Content: []string{},
							},
						},
					},
				},
			},
		},
	}
}
