package getpair

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Структура значений валютных пар
type CUR struct {
	Success   bool   `json:"success"`
	Timestamp int    `json:"timestamp"`
	Base      string `json:"base"`
	Date      string `json:"date"`
	Rates     struct {
		Rub float64 `json:"RUB"`
		Usd float64 `json:"USD"`
	} `json:"rates"`
}

// Получение значений валютных пар
func (data *CUR) GetCur() (err error) {
	token, err := ioutil.ReadFile(".secret/fixer")
	if err != nil {
		return err
	}

	url := fmt.Sprintf("http://data.fixer.io/api/latest?access_key=%v&base=EUR&symbols=RUB,USD", string(token))

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	json.Unmarshal(body, &data)

	return nil
}
