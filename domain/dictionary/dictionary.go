package dictionary

import (
	"fmt"
)

type Dictionary struct {
	Id       string `json:"id"`
	English  string `json:"english"`
	Japanese string `json:"japanese"`
}

func (dic Dictionary) String() string {
	return fmt.Sprintf(
		"Dictionary {\n Id : %v\n English : %v\n Japanese : %v\n}\n",
		dic.Id,
		dic.English,
		dic.Japanese,
	)
}
