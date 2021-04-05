package db

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func LoadScripts(dir string) (map[string]string, error) {
	var queries = map[string]string{}
	var cwd, cwderr = os.Getwd()

	if cwderr != nil {
		return queries, cwderr
	}

	var path = fmt.Sprintf("%s/sql/%s", cwd, dir)
	var files, fileserr = ioutil.ReadDir(path)

	if fileserr != nil {
		return queries, fileserr
	}

	for _, f := range files {
		if f.IsDir() == false {
			var name = f.Name()

			var c, cerr = ioutil.ReadFile(fmt.Sprintf("%s/%s", path, name))

			if cerr != nil {
				return queries, cerr
			}

			queries[strings.ReplaceAll(name, ".sql", "")] = string(c)
		}
	}

	return queries, nil
}
