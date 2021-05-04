package db

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func LoadScripts(dir string) (map[string]string, error) {
	queries := map[string]string{}
	cwd, err := os.Getwd()

	if err != nil {
		return queries, err
	}

	path := fmt.Sprintf("%s/sql/%s", cwd, dir)
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return queries, err
	}

	for _, f := range files {
		if f.IsDir() == false {
			name := f.Name()
			c, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", path, name))

			if err != nil {
				return queries, err
			}

			queries[strings.ReplaceAll(name, ".sql", "")] = string(c)
		}
	}

	return queries, nil
}
