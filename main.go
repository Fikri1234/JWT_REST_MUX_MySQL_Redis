package main

import (
	"JWT_REST_MUX_MySQL/configuration"
	"JWT_REST_MUX_MySQL/router"
	"JWT_REST_MUX_MySQL/util"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func init() {

	// read config environment
	configuration.ReadConfig()

	util.Pool = util.SetupRedisJWT()

}

func main() {

	var err error

	// Setup database
	configuration.DB, err = configuration.SetupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer configuration.DB.Close()

	// Router configuration
	router := router.NewRouter()
	port := viper.GetString("PORT")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
