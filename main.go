package main

import (
	
	"io"

	"github.com/gin-gonic/gin"
)



func getData(c *gin.Context) {

	c.JSON(200, gin.H{
		"data": "Hi I am GIN Framework",
	})


}

func getDataPost(c *gin.Context) {
	//var json map[string]interface{}
	body := c.Request.Body
	value , _ := io.ReadAll(body)

	// if err := c.ShouldBindJSON(&json); err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(200, gin.H{ 
		"message": "Hi I am GIN Framework POST METHOD",
		"bodyData": string(value),
	})				



}

// http://localhost:8080/getQueryString?name=John&age=30
func getQueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"data" : "Hi I am GIN Framework GET METHOD",
		"message": "Query String",
		"queryString": c.Request.URL.RawQuery,
		"name": name,
		"age":  age,
	})
}



// http://localhost:8080/getUrlData/name/Mark/age/30
func getUrlData(c *gin.Context){
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"data" : "Hi I am GIN Framework GET METHOD",
		"message": "URL Data",	
		"queryString": c.Request.URL.RawQuery,
		"name": name,
		"age":  age,
	})
}






func main() {

	// router := gin.Default()

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/getData"  , getData)

	router.POST("/getDataPost", getDataPost)

	router.GET("/getQueryString", getQueryString)

	router.GET("/getUrlData/:name/:age", getUrlData)

	router.Run(":8080")


}







