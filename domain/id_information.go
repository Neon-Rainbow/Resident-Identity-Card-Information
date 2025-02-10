package domain

type Address struct {
	Province string `json:"province"` // 省份
	City     string `json:"city"`     // 城市
	Area     string `json:"area"`     // 区县
}

type IDInformation struct {
	Age           int     `json:"age"`           // 年龄
	Birthday      string  `json:"birthday"`      // 出生日期
	Gender        string  `json:"gender"`        // 性别
	IsAdult       bool    `json:"is_adult"`      // 是否成年
	Zodiac        string  `json:"zodiac"`        // 生肖
	Constellation string  `json:"constellation"` // 星座
	Address       Address `json:"address"`       // 地址
}
