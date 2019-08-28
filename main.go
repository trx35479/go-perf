package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-api-perf/models"
)

type Headers struct {
	MagicToken string
	Xv         string
	Url        string
}

func (h Headers) GetToken() []interface{} {
	sl := make([]interface{}, 0)
	sl = append(sl, h.Xv)
	sl = append(sl, h.MagicToken)
	sl = append(sl, h.Url)

	return sl
}

func GetProductIDs(b []interface{}) (*models.Products, error) {
	log.Printf("Getting the product ids of: %s", b[2])

	client := &http.Client{}
	req, err := http.NewRequest("GET", b[2].(string), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-v", b[0].(string))
	req.Header.Add("anz-magic-token", b[1].(string))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var products *models.Products
	err = json.Unmarshal([]byte(body), &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func main() {
	h := Headers{"bc339c8d57843580d0020e95778aa22e503a581cf82679e6c76f168544de3b20",
		"1",
		"https://api-np.anz/cds-au/v1/banking/products/"}
	t := h.GetToken()

	values, err := GetProductIDs(t)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range values.Data.Products {
		fmt.Println(v.ProductId)
	}
}

//func main() {
//	h := Headers{"bc339c8d57843580d0020e95778aa22e503a581cf82679e6c76f168544de3b20",
//		"1",
//		"https://api-np.anz/cds-au/v1/banking/products/e36042da-fa74-d28b-30b1-0b41851e098b"}
//	t := h.GetToken()

//	client := &http.Client{}

//	for i := 0; i < 30; i++ {
//		req, err := http.NewRequest("GET", t[2].(string), nil)
//		if err != nil {
//			log.Fatal(err)
//		}
// for non-prod | if prod do not include the magic token
//		req.Header.Add("x-v", t[0].(string))
//		req.Header.Add("anz-magic-token", t[1].(string))

//		resp, err := client.Do(req)
//		if err != nil {
//			log.Fatal(err)
//		}

//		if resp.StatusCode != 200 {
//			fmt.Printf("The resquest responded with %d error\n", resp.StatusCode)
//		}

//		for k, v := range resp.Header {
//			if k == "Etag" {
//				fmt.Printf("The request was served from the cached with Etag of: %s\n", v)
//				break
//			}
//		}

//		time.Sleep(1 * time.Second)
//	}
//}
