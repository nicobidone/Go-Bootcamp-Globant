package thereader

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/nicob/theRESTservice/entity"
)

var myClient = &http.Client{Timeout: 2 * time.Second}

func GetJson(url string) *[]entity.Product {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "products-tutorial")

	res, getErr := myClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var products []entity.Product
	jsonErr := json.Unmarshal(body, &products)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	//fmt.Println(products)
	return &products
}
