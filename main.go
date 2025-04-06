package main

import (
	"io"
	// "net/http"
	// "time"

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
// http://localhost:8080/getUrlData/name/Mark/age/30






func main() {

	// router := gin.Default()

	router := gin.New()
	router.Use(gin.Logger())


	// auth := router.BasicAuth(gin.Accounts{
	// 	"user": "password",
	// 	"user2" : "password2",
	// 	"user3" : "password",

	// })


	// router.GET("/getData"  , getData)

	// router.POST("/getDataPost", getDataPost)

	// router.GET("/getQueryString", getQueryString)

	// router.GET("/getUrlData/:name/:age", getUrlData)

	// router.Run(":8080")

	//The above line is equivalent to the following two lines
	// http.ListenAndServe(":8080", router)

	//This is the same as the above line
	// server := &http.Server{
	// 	Addr:    ":9091",
	// 	Handler: router,
	// 	ReadTimeout: 10 * time.Second,
	// 	WriteTimeout: 10 * time.Second,
	// }

	// server.ListenAndServe()






	//ROUTE GROUPING

	//without any group
	router.GET("/getData", getData)
	router.POST("/getDataPost", getDataPost)
	router.GET("/getQueryString", getQueryString)
	router.GET("/getUrlData/:name/:age", getUrlData)

	//route grouping in gin
	admin := router.Group("/admin")
	
	client := router.Group("/client")

	{
		admin.GET("/getData", getData)
		admin.POST("/getDataPost", getDataPost)
		admin.GET("/getQueryString", getQueryString)
		admin.GET("/getUrlData/:name/:age", getUrlData)
	}

	{
		client.GET("/getData", getData)
		client.POST("/getDataPost", getDataPost)
		client.GET("/getQueryString", getQueryString)
		client.GET("/getUrlData/:name/:age", getUrlData)
	}


	router.Run(":8080")
	


}







