package usecase

import (
	"fmt"
	"testing"
	"time"
)

func Test_calculateConstellation(t *testing.T) {
	tests := []struct {
		name     string
		idNumber string
		want     string
	}{
		// 白羊座测试用例 (3.21-4.19)
		{name: "白羊座开始", idNumber: "11111119210321111", want: "白羊座"},
		{name: "白羊座结束", idNumber: "11111119210419111", want: "白羊座"},

		// 金牛座测试用例 (4.20-5.20)
		{name: "金牛座开始", idNumber: "11111119210420111", want: "金牛座"},
		{name: "金牛座结束", idNumber: "11111119210520111", want: "金牛座"},

		// 双子座测试用例 (5.21-6.21)
		{name: "双子座开始", idNumber: "11111119210521111", want: "双子座"},
		{name: "双子座结束", idNumber: "11111119210621111", want: "双子座"},

		// 巨蟹座测试用例 (6.22-7.22)
		{name: "巨蟹座开始", idNumber: "11111119210622111", want: "巨蟹座"},
		{name: "巨蟹座结束", idNumber: "11111119210722111", want: "巨蟹座"},

		// 狮子座测试用例 (7.23-8.22)
		{name: "狮子座开始", idNumber: "11111119210723111", want: "狮子座"},
		{name: "狮子座结束", idNumber: "11111119210822111", want: "狮子座"},

		// 处女座测试用例 (8.23-9.22)
		{name: "处女座开始", idNumber: "11111119210823111", want: "处女座"},
		{name: "处女座结束", idNumber: "11111119210922111", want: "处女座"},

		// 天秤座测试用例 (9.23-10.23)
		{name: "天秤座开始", idNumber: "11111119210923111", want: "天秤座"},
		{name: "天秤座结束", idNumber: "11111119211023111", want: "天秤座"},

		// 天蝎座测试用例 (10.24-11.22)
		{name: "天蝎座开始", idNumber: "11111119211024111", want: "天蝎座"},
		{name: "天蝎座结束", idNumber: "11111119211122111", want: "天蝎座"},

		// 射手座测试用例 (11.23-12.21)
		{name: "射手座开始", idNumber: "11111119211123111", want: "射手座"},
		{name: "射手座结束", idNumber: "11111119211221111", want: "射手座"},

		// 摩羯座测试用例 (12.22-1.19)
		{name: "摩羯座开始", idNumber: "11111119211222111", want: "摩羯座"},
		{name: "摩羯座结束", idNumber: "11111119210119111", want: "摩羯座"},

		// 水瓶座测试用例 (1.20-2.18)
		{name: "水瓶座开始", idNumber: "11111119210120111", want: "水瓶座"},
		{name: "水瓶座结束", idNumber: "11111119210218111", want: "水瓶座"},

		// 双鱼座测试用例 (2.19-3.20)
		{name: "双鱼座开始", idNumber: "11111119210219111", want: "双鱼座"},
		{name: "双鱼座结束", idNumber: "11111119210320111", want: "双鱼座"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateConstellation(tt.idNumber); got != tt.want {
				t.Errorf("calculateConstellation() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_calculateAge(t *testing.T) {
	birthdayToday := time.Now().AddDate(-20, 0, 0).Format("20060102")
	birthdayPassed := time.Now().AddDate(-20, -11, 0).Format("20060102")
	birthdayNotYet := time.Now().AddDate(-20, 1, 0).Format("20060102")

	tests := []struct {
		name     string
		idNumber string
		want     int
	}{
		{name: "Birthday today", idNumber: fmt.Sprintf("111111%s1111", birthdayToday), want: 20},
		{name: "Birthday passed this year", idNumber: fmt.Sprintf("111111%s1111", birthdayPassed), want: 20},
		{name: "Birthday not yet this year", idNumber: fmt.Sprintf("111111%s1111", birthdayNotYet), want: 19},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateAge(tt.idNumber); got != tt.want {
				t.Errorf("calculateAge() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_calculateGender(t *testing.T) {
	tests := []struct {
		name     string
		idNumber string
		want     string
	}{
		{name: "Male", idNumber: "111111199901011111", want: "男"},
		{name: "Female", idNumber: "111111199901011121", want: "女"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateGender(tt.idNumber); got != tt.want {
				t.Errorf("calculateGender() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_calculateZodiac(t *testing.T) {
	tests := []struct {
		name     string
		idNumber string
		want     string
	}{
		{name: "鼠年", idNumber: "111111198401011111", want: "鼠"},
		{name: "牛年", idNumber: "111111198502021111", want: "牛"},
		{name: "虎年", idNumber: "111111198603031111", want: "虎"},
		{name: "兔年", idNumber: "111111198704041111", want: "兔"},
		{name: "龙年", idNumber: "111111198805051111", want: "龙"},
		{name: "蛇年", idNumber: "111111198906061111", want: "蛇"},
		{name: "马年", idNumber: "111111199007071111", want: "马"},
		{name: "羊年", idNumber: "111111199108081111", want: "羊"},
		{name: "猴年", idNumber: "111111199209091111", want: "猴"},
		{name: "鸡年", idNumber: "111111199310101111", want: "鸡"},
		{name: "狗年", idNumber: "111111199411111111", want: "狗"},
		{name: "猪年", idNumber: "111111199512121111", want: "猪"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateZodiac(tt.idNumber); got != tt.want {
				t.Errorf("calculateZodiac() = %v, want %v", got, tt.want)
			}
		})
	}
}
