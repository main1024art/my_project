package aaa
import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)
type monster struct {
	Name string
	Age int
	Skill string
}

func (m *monster) Store()  bool {
    date, err := json.Marshal(&m)
    if err != nil {
		fmt.Printf("序列化错误 err=%v",err)
	}
	fmt.Println(string(date))
	
	
	filePath := "d:/monster.ser"
	err = ioutil.WriteFile(filePath,date,0666)
	if err != nil {
		fmt.Printf("WriteFile错误 err=%v",err)
	}
	return true

}
func (m *monster) Restore() bool {
	filePath := "d:/monster.ser"
	date, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("ReadFile错误 err=%v",err)
	}
	err = json.Unmarshal(date,&m)
	if err != nil {
		fmt.Printf("反序列化错误 err=%v",err)
	}
	return true
}


