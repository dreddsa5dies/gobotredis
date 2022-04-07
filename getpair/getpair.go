package getpair

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Получение значений валютных пар
func GetCur() (data []byte, err error) {
	token, err := ioutil.ReadFile(".secret/fixer")
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("http://data.fixer.io/api/latest?access_key=%v&base=EUR&symbols=RUB,USD", string(token))

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	log.Println("[*] pair get successfully")

	return body, nil
}
