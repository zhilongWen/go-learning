package main

import (
	"net/http"
	"test-server/config"
	"test-server/init/init_mysql"
	"test-server/init/init_router"
	"test-server/init/register_table"
	"time"
)

func main() {
	println("Hello, world!")

	register_table.RegistTable(init_mysql.InitMySQL(config.Dbconfig.Admin))
	defer init_mysql.DB.Close()

	router := init_router.InitRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * 60 * time.Second,
		WriteTimeout:   10 * 60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}
