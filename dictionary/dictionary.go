package dictionary

import (
	"fmt"
)

type Dictionary struct {
	Id       string
	English  string
	Japanese string
}

func (dic Dictionary) String() string {
	return fmt.Sprintf("Dictionary {\n Id : %v\n English : %v\n Japanese : %v\n}\n", dic.Id, dic.English, dic.Japanese)
}
