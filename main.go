package gomongocrud

import "github.com/julienschmidt/httprouter"

func main () {
	router := httprouter.New()
	router.POST("/signup", signup)
	router.POST("/login", login)
}