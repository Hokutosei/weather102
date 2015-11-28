package getters

import (
	"encoding/json"
	"fmt"
	"net/http"

	utils "github.com/Hokutosei/hokutoseiUtils"
)

// HTTPGetter main http getter or GET request
func HTTPGetter(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		utils.Error(fmt.Sprintf("HTTP GET %v", err))
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
