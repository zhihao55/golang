package bean

type User struct {
	Id       uint
	Username string
	Password string
}
// TableName 解决gorm表明映射
func (u User) TableName() string {
	return "g_user"
}