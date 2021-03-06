package userDao

import (
	"gorm.io/gorm"
	userModel2 "kChatRoom/app/client/model/userModel"
	"kChatRoom/utils/db"
)

type userDao struct {
	Db *gorm.DB
}

//NewUserDao new user dao
func NewUserDao() *userDao {
	return &userDao{
		Db: db.NewDB(),
	}
}

// GetUserByName 通过用户名获取用户
func (t *userDao) GetUserByName(username string) *userModel2.UserModel {
	user := &userModel2.UserModel{}
	Db := t.Db //拷贝全局Db 避免变量污染
	Db = Db.Where("username = ?", username)
	_ = Db.Take(user)
	return user
}

// GetUserByMail 通过邮箱查找用户
func (t *userDao) GetUserByMail(mail string) *userModel2.UserModel {
	user := &userModel2.UserModel{}
	Db := t.Db //拷贝全局Db 避免变量污染
	Db = Db.Where("mail = ?", mail)
	_ = Db.Take(user)
	return user
}

// GetUserById 通过用户id获取用户
func (t *userDao) GetUserById(id int) *userModel2.UserModel {
	Db := t.Db
	user := &userModel2.UserModel{}
	_ = Db.First(user, id)
	return user
}

// SaveHead 保存用户头像
func (t *userDao) SaveHead(mail, head string) error {
	Db := t.Db
	res := Db.Model(&userModel2.UserModel{}).Where("mail = ?", mail).Select("head").Update("head", head)
	return res.Error
}

// AddUser 添加用户
func (t *userDao) AddUser(user *userModel2.UserModel) (bool, string) {
	//判断是否存在
	findUser := t.GetUserByName(user.UserName)
	if findUser.UserName != "" {
		return false, ",用户名已存在！"
	}
	Db := t.Db //拷贝全局Db 避免变量污染
	res := Db.Create(user)
	if err := res.Error; err != nil {
		return false, ""
	} else {
		return true, ""
	}
}

//SaveName 保存用户名
func (t *userDao) SaveName(mail, name string) error {
	//fmt.Println(mail)
	Db := t.Db
	res := Db.Model(&userModel2.UserModel{}).Where("mail = ?", mail).Select("username").Update("username", name)
	return res.Error
}
