package users

type Users struct {
}

// New()函数是工厂方法，用于创建Users实例
// 控制器通过这个方法获取业务逻辑实例：users: userLogic.New()
func New() *Users {
	return &Users{}
}
