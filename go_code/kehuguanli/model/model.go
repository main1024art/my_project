package model
import (
	"fmt"
)

type Kehu struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}
func NewKehu(id int, name string, gender string, age int, phone string, email string) Kehu {
	return Kehu{
		Id : id,
		Name : name,
		Gender : gender,
		Age : age,
		Phone : phone,
		Email : email,
	}
}

func (k *Kehu) Info() string {
	aa := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", k.Id, k.Name, k.Gender, k.Age, k.Phone, k.Email)
	return aa

}
func NewKehu2(name string, gender string, age int, phone string, email string) Kehu {
	return Kehu{
		Name : name,
		Gender : gender,
		Age : age,
		Phone : phone,
		Email : email,
	}
}