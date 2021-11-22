package safety

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_Encrypt(t *testing.T) {
	obj:=make(map[string]interface{})
	obj["test"]=1
	obj["测试"]=2
	s := "12354t543532"
	k := "123456fdsafdsajkklfjadsa"
	b := Encrypt([]byte(k),s)
	fmt.Println("b=", b)
	d := Decrypt([]byte(k),b)
	fmt.Println("d=",d)
	if string(d) != string(s) {
		t.Error("加密失敗")
	}
	b1,_:=json.Marshal(obj)
	b=Encrypt([]byte(k),string(b1))
	fmt.Println("b1=",b)
	d=Decrypt([]byte(k),b)
	fmt.Println("d1=",d)
}
