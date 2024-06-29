package model

type Account struct {
	hao string
	yuer float64
	pwd string
}


func (a *Account) Set(h string, y float64, p string) {
	if len(h) > 6 && len(h) < 10 {
		a.hao = h
	}
	if y > 20 {
		a.yuer = y
	}
	if len(p) == 6 {
		a.pwd = p
	}
}
func (a *Account) Get() *Account {
	return a
}