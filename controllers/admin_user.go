package controllers

type AdminUser struct {
	Id int
	UserName string
	Password string
	Email string
	Author int
}
//
//func (user AdminUser)ValidateUser() error {
//	orm := getLink() //获得用于操作数据库的orm
//	var u AdminUser
//	orm.Where("username=? and password=?", user.UserName, user.Password).Find(&u)
//	if u.UserName == "" {
//		return errors.New("用户名或密码错误！")
//	}
//	return nil
//}
