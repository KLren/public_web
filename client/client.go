package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SideBar struct {
	FirstLayer  string
	SecondLayer []string
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("/go/src/website/view/*") // Load HTML files(template)
	router.Static("../../../home/pic", "/home/pic")

	router.GET("/index", indexRender)

	router.Run(":443")
}

// Handler for rendering HTML
func indexRender(c *gin.Context) {
	// Define SideBar content
	sidebar := []SideBar{
		{"Home", []string{"Overview", "Updates", "Reports"}},
		{"Dashboard", []string{"Overview", "Weekly", "Monthly", "Annually"}},
		{"Orders", []string{"New", "Processed", "Shipped", "Returned"}},
	}

	productList := GetDataFromAPI("http://172.18.0.4/api/v1/products")
	//productList := GetDataFromAPI("http://172.18.0.3/api/v1/products/1")

	c.HTML(http.StatusOK, "index.html", gin.H{
		// Define Header content
		"Headers":     []string{"Features", "Pricing", "FAQs", "About"},
		"SideBar":     sidebar,
		"ProductList": productList,
	})
}

func GetDataFromAPI(apiUrl string) map[string]interface{} {
	resp, err := http.Get(apiUrl)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var sresp map[string]interface{}
	err = json.Unmarshal(body, &sresp)
	if err != nil {
		log.Println(err)
	}

	return sresp
}
