package pcc

import (
	"fmt"

	date "github.com/fxtlabs/date"
)

//Process exemplo
func Process(c int) {
	hoje := date.Today()
	fmt.Println(c, "dias a frente = ", hoje.AddDate(0, 0, c).Format(date.ISO8601))
}
