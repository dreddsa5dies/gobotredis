package getpair

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dreddsa5dies/gobotredis/storage"
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

	return body, nil
}

// Сохранение значений валютных пар и их обновление 2 раза в день
func UpdatePair() {
	for {
		currentTime := time.Now()
		key := currentTime.Format("09-07-2017")

		pair, err := GetCur()
		if err != nil {
			log.Println(err)
		}
		err = storage.SetData(pair, key)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(12 * time.Hour)
	}
}
