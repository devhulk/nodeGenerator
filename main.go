package main

/*
Create basic nodejs generator.

Requirements:
	Write to package.json (use npm install or perhaps
	use base template and allow user to fill in options)
*/

func main() {
	t := generator{map[string]bool{
			"package.json": true,
			"app.js": true,
			"index.html": true,
			"js": true,
			"css": true,
		},
		 ProjectConfig{
			name: "",
			author: "",
			pt: "",
		},
	}
	readFiles("/Users/devhulk/go/tmp", t.FileMap)
	writeFiles("/Users/devhulk/go/tmp", t.FileMap)
}
