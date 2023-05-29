package serv

import (
	"encoding/json"
	"fmt"
	"go-gin-amis/app/dal/model"
	"testing"
)

func TestE用户操作_FindOne(t *testing.T) {
	用户操作 := S用户操作
	l, c, err := 用户操作.E查询用户列表("", 1, 10, "id", "desc")
	cc, err := json.Marshal(l)
	fmt.Println(string(cc))
	fmt.Println(c)
	fmt.Println(err)

}
func TestE用户操作_ADD(t *testing.T) {
	用户操作 := S用户操作
	articleData := new(model.User)
	articleData.Username = "aaa"
	articleData.Password = "bbb"

	err := 用户操作.E创建用户(articleData)
	fmt.Println(err)
	
}
