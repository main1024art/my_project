package model


type stu struct{
	Name string
	Score float64
}
func Say(n string, m float64) *stu {
    return &stu{
		Name : n,
		Score : m,
	}
}