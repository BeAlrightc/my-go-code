package model
//定义一个结构体
type student struct{
	Name string
	score float64
}

//因为student结构体首字母是小写，因此是只能在
//model包使用，通过工厂模式解决
func NewStudent(n string,s float64) *student{
	return &student{
		Name : n,
		score : s,
	}
}

//如果score字段首字母小写，则在其他包不可以访问，我们可以提供方法
func (s *student) GetScore() float64{
	return s.score
	//实际上golang底层会将这个优化成(*s).score
	//return (*s).score
}