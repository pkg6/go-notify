package dingtalk

type At struct {
	Mobiles []string `json:"atMobiles"`
	UserIds []string `json:"atUserIds"`
	IsAtAll bool     `json:"isAtAll"`
}

func (a *At) AtMobiles(mobiles []string) {
	a.Mobiles = mobiles
}
func (a *At) AtUserIds(userIds []string) {
	a.UserIds = userIds
}
func (a *At) AtAll() {
	a.IsAtAll = true
}
