package db

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func LoadScripts(dir string) (map[string]string, error) {
	var queries = map[string]string{}
	var cwd, cwde = os.Getwd()

	if cwde != nil {
		return queries, cwde
	}

	var path = fmt.Sprintf("%s/sql/%s", cwd, dir)
	var files, filese = ioutil.ReadDir(path)

	if filese != nil {
		return queries, filese
	}

	for _, f := range files {
		if f.IsDir() == false {
			var name = f.Name()

			var c, ce = ioutil.ReadFile(fmt.Sprintf("%s/%s", path, name))

			if ce != nil {
				return queries, ce
			}

			queries[strings.ReplaceAll(name, ".sql", "")] = string(c)
		}
	}

	return queries, nil
}
