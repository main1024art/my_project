package service

import (
	"kehuguanli/model"
)

type Kehuservice struct {
	Kehus   []model.Kehu
	KehuNum int
}

func NewKehuservice() *Kehuservice {
	Kehuservice := &Kehuservice{}
	Kehuservice.KehuNum = 1
	kehu11 := model.NewKehu(1, "liu", "男", 12, "112", "12345@qq.com")
	Kehuservice.Kehus = append(Kehuservice.Kehus, kehu11)
	return Kehuservice
}

func (k *Kehuservice) List() []model.Kehu {
	return k.Kehus
}

func (k *Kehuservice) Add(ke model.Kehu) bool {
	k.KehuNum++
	ke.Id = k.KehuNum
	k.Kehus = append(k.Kehus, ke)
	return true
}
func (k *Kehuservice) Del(id int) bool {
	l := k.fin(id)
	k.Kehus = append(k.Kehus[:l], k.Kehus[l+1:]...)
	return true
}

func (k *Kehuservice) fin(id int) int {
	l := -1
	for i := 0; i < len(k.Kehus); i++ {
		if k.Kehus[i].Id == id {
			l = i
		}
	}
	return l
}
