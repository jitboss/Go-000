package main

import (
	"database/sql"
	"github.com/pkg/errors"
	"log"
)

var errMsg = map[error]string{
	sql.ErrNoRows: "没有找到数据库中数据",
}
var dict map[string]string

func dao(nation string) (string, error) {
	dict = make(map[string]string)
	dict["美国"] = "华盛顿"
	dict["英国"] = "伦敦"

	value, ok := dict[nation]
	if ok {
		return value, nil
	} else {
		return "", errors.Wrapf(sql.ErrNoRows, "没有找到首都  %nation", nation)
	}

}

func service(nation string) (string, error) {
	return dao(nation)
}

func business(nation string) string {
	var (
		res string
		err error
	)
	if res, err = service(nation); err != nil {
		log.Printf("%+v", err)
		return errMsg[errors.Cause(err)]
	}

	log.Printf("首都是 %d", res)
	return "ok"
}

func main() {
	log.Println(business("美国"))
	log.Println(business("法国"))
}
