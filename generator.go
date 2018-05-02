package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"encoding/json"
	"io"
	"fmt"
)

type generator struct {
	FileMap map[string]bool
	ProjectConfig
}

type ProjectConfig struct {
	name string
	author string
	pt string
}

func getProjectConfig(config string) {
	//Not Working
	dec := json.NewDecoder(strings.NewReader(config))
	for {
		var m ProjectConfig
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.name, m.author)
	}
}

func readFiles(dir string, ftc map[string]bool) {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalln(err)
	}


	for _, f := range files {
		//fmt.Println(f.Name(), "already exists")
		if f.Name() == "gen.json" {
			const jsonStream = `
{"name":"testProject","author":"testAuthor","type":"nodejs"}
`
			getProjectConfig(jsonStream)
		}
		for k := range ftc {
			if k == f.Name() {
				ftc[f.Name()] = false
			}
		}
	}

}

func writeFiles(dir string, ftc map[string]bool) {

	for k := range ftc {
		if (ftc[k] == true) && (k != "css" && k != "js")  {
			_, err := os.Create(dir + "/" + k)
			if err != nil {
				log.Fatalln(err)
			}

		}
		if ftc[k] == true && k == "css"  {
			err := os.MkdirAll(dir + "/" + k,0700)
			if err != nil {
				log.Fatalln(err)
			}
		}
		if ftc[k] == true && k == "js"  {
			err := os.MkdirAll(dir + "/" + k,0700)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
