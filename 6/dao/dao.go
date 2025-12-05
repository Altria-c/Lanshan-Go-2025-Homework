package dao

// 模拟数据库
var database = map[string]string{}

func AddUser(username string, password string) {
	database[username] = password
	SaveDatabase()
}

func FindUser(username string, password string) bool {
	if pwd, ok := database[username]; ok {
		if pwd == password {
			return true
		}
	}
	return false
}
func SelectPasswordFromUsername(username string) string {
	return database[username]
}

// 修改密码
func ChangePassword(username string, newpwd string) bool {
	if _, ok := database[username]; ok {
		database[username] = newpwd
		SaveDatabase()
		return true
	}
	return false
}
