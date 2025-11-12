package log

import (
	"E-04/richerror"
	"encoding/json"
	"os"
)

type Log struct {
	Errors []richerror.RichError
}

func (l Log) Save() {
	f, _ := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	defer f.Close()
	data, _ := json.Marshal(l.Errors)
	f.Write(data)
}

func (l Log) Append(r *richerror.RichError) {
	l.Errors = append(l.Errors, *r)
}
