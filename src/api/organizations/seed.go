package organizations

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Seed struct {
	Organizations []Model `json:"organizations"`
}

func NewSeed() *Seed {
	var seed = new(Seed)
	var file, ferr = os.Open("./seeds/organizations.json")

	if ferr != nil {
		log.Error(ferr)
	}

	var bytes, berr = ioutil.ReadAll(file)

	if berr != nil {
		log.Error(berr)
	}

	var jerr = json.Unmarshal(bytes, &seed)

	if jerr != nil {
		log.Error(jerr)
	}

	return seed
}
