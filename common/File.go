package common

import "os"


//判断本地文件是否存在
func Isfile(path string) bool  {
	_,err:=os.Stat(path)
	if err != nil {
		return false
	}
	return true
}