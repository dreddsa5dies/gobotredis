package getpair

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/peterhellberg/fixer"
)

func GetCur() {
	ftkn, err := ioutil.ReadFile(".secret/fixer")
	if err != nil {
		log.Printf("unable to read fixer token: %v", err)
	}

	fixer.AccessKey(string(ftkn))

	resp, err := fixer.Latest(context.Background(),
		fixer.Base(fixer.RUB),
		fixer.Symbols(
			fixer.USD,
			fixer.EUR,
		),
	)
	if err != nil {
		return
	}

	encode(resp)
}

func encode(v interface{}) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", " ")
	enc.Encode(v)
}
