package handler

import (
	"fmt"
	"net/http"
	"next-deploy/lib"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var list []map[string]interface{}
	lib.MysqlConnect().Table("product").Find(&list)
	currentTime := time.Now().Format(time.RFC850)
	fmt.Fprintf(w, currentTime,list)
}
