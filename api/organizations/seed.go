package organizations

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/aacebo/equinox-api/log"
)

type Seed struct {
	Organizations []Model `json:"organizations"`
}

func NewSeed() *Seed {
	var seed = new(Seed)
	var file, err = os.Open("./seeds/organizations.json")

	if err != nil {
		log.Error.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		log.Error.Fatal(err)
	}

	err = json.Unmarshal(bytes, &seed)

	if err != nil {
		log.Error.Fatal(err)
	}

	return seed
}
