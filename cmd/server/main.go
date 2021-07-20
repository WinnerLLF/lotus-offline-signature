package main

import (
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/gin-gonic/gin"
	"lotus-offline-signature/api"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

var (
	// The generated HTML save directory
	htmlOutPath = "./html"
	// Static file template directory
	templatePath = "./html"
)

func init() {
	address.CurrentNetwork = address.Mainnet
}

func main() {
	fmt.Printf("server start...\n")
	// init off-line wallet
	api.InitCache()

	// init api
	router := api.RouterApiServer()

	// web html
	router.LoadHTMLGlob("html/*")
	router.GET("/index", func(c *gin.Context) {
		GetGenerateHtml()
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// run server
	if err := router.Run(":8787"); err != nil {
		fmt.Printf("server start err:%v\n", err)
	}
}

// A method to generate static files
func GetGenerateHtml() {
	//1.Access to the template
	contenstTmp, err := template.ParseFiles(filepath.Join(templatePath, "index.html"))
	if err != nil {
		fmt.Println("Failed to get template file")
	}
	//2.Gets the HTML generation path
	fileName := filepath.Join(htmlOutPath, "htmlindex.html")
	//4.Generate static files
	generateStaticHtml(contenstTmp, fileName, gin.H{})
}

// Generate static files
func generateStaticHtml(template *template.Template, fileName string, product map[string]interface{}) {
	// Determine whether the static file exists
	if exist(fileName) {
		err := os.Remove(fileName)
		if err != nil {
			fmt.Println("Failed to remove file")
		}
	}
	// Generate static files
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("Failed to open file")
	}
	defer file.Close()
	template.Execute(file, &product)
}

// Determine whether the file exists
func exist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}
