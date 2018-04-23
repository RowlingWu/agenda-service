package entities

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type serv interface {
	load()
}

var (
	DB       *gorm.DB
	servlist []serv
)

func addServ(newserv serv) {
	servlist = append(servlist, newserv)
}

func Initial(dbpath string) {
	db, err := gorm.Open("sqlite3", dbpath)
	checkError(err)
	DB = db
	var erro interface{}
	finished := make(chan bool)
	for _, s := range servlist {
		go func(s serv) {
			defer func() {
				if e := recover(); e != nil {
					erro = e
				}
				finished <- true
			}()
			s.load()
		}(s)
	}
	for _ = range servlist {
		<-finished
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
