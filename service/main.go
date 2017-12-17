package main
import (
    "github.com/RowlingWu/agenda-service/service/entities"
    "github.com/RowlingWu/agenda-service/service/service"
    "os"
    flag "github.com/spf13/pflag"
)

const (
  PORT string = "8080"
)

func main() {
  port := os.Getenv("PORT")
  if len(port) == 0 {
		port = PORT
	}

	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	dbpath := flag.StringP("db", "d", "", "sqlite database file")
	flag.Parse()

	if len(*pPort) != 0 {
		port = *pPort
	}
	if len(*dbpath) == 0 {
		os.Mkdir("data", 0755)
		*dbpath = "data/agenda.db"
	}

	entities.Initial(*dbpath)
	server := service.NewServer()
	server.Run(":" + port)
}
