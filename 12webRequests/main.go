package main 

// import (
// 	"fmt"
// 	"net/http"
// )

// func homeHandler(w http.ResponseWriter, r *http.Request){
// 	fmt.Println("welcome to the home page ")
// }

// func aboutHandler(w http.ResponseWriter, r* http.Request){
// 	fmt.Println("about page")
// }

// func main (){
// 	http.HandleFunc("/",homeHandler)
// 	http.HandleFunc("/about",aboutHandler)

// 	fmt.Println("listening on port 8080")
// 	http.ListenAndServe(":8080",nil)
// }

///USING GIN (GO HTTP FRAMEWORK)
import (
	"net/http"
	"github.com/gin-gonic/gin"
)
func main (){
	router:=gin.Default()

	router.GET("ping",func (c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"message":"pong",
		})
	})
	router.Run(":8080")

}