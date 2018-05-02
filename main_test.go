package main
/*
Create basic web generator for pure javascript
development.

Requirements:
	Create necessary files and folders
		* package.json
		* index.html
		* js/
		* css/
	Write to package.json (use npm install or perhaps
	use base template and allow user to fill in options)
*/

import (
	"testing"
	//"io/ioutil"
	//"fmt"
	"fmt"
)

func TestReadFiles (t *testing.T) {

	//find specified directory
	/*
	*/

	cf := readFiles("/Users/devhulk/go/tmp")
	if cf[1] != "test.txt"  {
		fmt.Println(cf[0])
		t.Errorf("Expected empty string but got %v", cf)
	}

}
func TestWriteFile (t *testing.T) {

	ts := []string{"package.json"}

	wf := writeFile("/Users/devhulk/go/tmp",ts)

	if wf[0] != "package.json" {
		t.Errorf("Expected empty string but got %v", wf)

	}


}
