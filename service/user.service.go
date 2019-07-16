package service

type UserInfo struct {
	Name string `json:"name"`
}

func GetUserInfo(name string) (u UserInfo, err error) {
	return UserInfo{
		Name: name,
	}, err
}
