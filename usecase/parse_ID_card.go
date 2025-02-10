package usecase

import (
	"fmt"
	"id-information/domain"
	"id-information/repository"
	"math"
	"strconv"
	"time"
)

// ParseIDNumber 函数用于解析身份证号码并返回相关信息
func ParseIDNumber(idNumber string) (*domain.IDInformation, *domain.ApiError) {
	var info domain.IDInformation
	// 判断身份证号码是否合法
	if !isValidIDNumber(idNumber) {
		return nil, domain.ErrInvalidIDNumber
	}
	address, apiError := repository.GetAddress(idNumber)
	if apiError != nil {
		return nil, apiError
	}
	info.Address = *address

	// 计算年龄
	info.Age = calculateAge(idNumber)

	// 计算性别
	info.Gender = calculateGender(idNumber)

	// 计算生日
	info.Birthday = fmt.Sprintf("%s-%s-%s", idNumber[6:10], idNumber[10:12], idNumber[12:14])

	// 判断是否成年
	if info.Age < 18 {
		info.IsAdult = false
	} else {
		info.IsAdult = true
	}

	info.Zodiac = calculateZodiac(idNumber)
	info.Constellation = calculateConstellation(idNumber)

	return &info, nil
}

// isValidIDNumber 函数用于验证身份证号码是否合法
//
// 参数:
//   - idNumber: 身份证号码
//
// 返回值:
//   - bool: 如果身份证号码合法，返回true；否则返回false
func isValidIDNumber(idNumber string) bool {
	// 身份证号码长度应为18位
	if len(idNumber) != 18 {
		return false
	}
	// 身份证号码的前17位应为数字
	for _, char := range idNumber[:17] {
		if char < '0' || char > '9' {
			return false
		}
	}
	// 身份证号码的第18位应为数字或大写字母X
	if (idNumber[17] < '0' || idNumber[17] > '9') && idNumber[17] != 'X' {
		return false
	}
	// 验证身份证号码的校验位
	if !check(idNumber) {
		return false
	}
	return true
}

// check 函数用于验证身份证号的校验位
func check(id string) bool {
	var sum, num int

	// 遍历前17位字符
	for index := 0; index < 17; index++ {
		// 计算权重并累加
		weight := int(math.Pow(2, float64(17-index))) % 11
		sum += weight * int(id[index]-'0')
	}

	// 计算校验位
	num = (12 - (sum % 11)) % 11

	// 检查校验位
	if num < 10 {
		return num == int(id[17]-'0')
	} else {
		return id[17] == 'X'
	}
}

// calculateAge 函数用于计算年龄
func calculateAge(idNumber string) int {
	// 从身份证号中提取出生年月日
	birthYear := idNumber[6:10]
	birthMonth := idNumber[10:12]
	birthDay := idNumber[12:14]

	// 将字符串转换为整数
	year, _ := strconv.Atoi(birthYear)
	month, _ := strconv.Atoi(birthMonth)
	day, _ := strconv.Atoi(birthDay)

	// 获取当前时间
	now := time.Now()
	age := now.Year() - year

	// 如果还没到生日，年龄减1
	if now.Month() < time.Month(month) || (now.Month() == time.Month(month) && now.Day() < day) {
		age--
	}

	return age
}

func calculateGender(idNumber string) string {
	// 从身份证号中提取性别码
	genderCode := idNumber[16]
	// 判断性别码的奇偶性
	if genderCode%2 == 0 {
		return "女"
	} else {
		return "男"
	}
}

func calculateZodiac(idNumber string) string {
	// 从身份证号中提取年份码
	yearCode := idNumber[6:10]
	// 将年份码转换为整数
	year, _ := strconv.Atoi(yearCode)
	// 计算生肖
	zodiacIndex := (year - 4) % 12
	// 定义生肖数组
	zodiacs := []string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}
	// 返回生肖
	return zodiacs[zodiacIndex]
}

// calculateConstellation 函数用于计算星座
func calculateConstellation(idNumber string) string {
	// 从身份证号中提取月份和日期
	month, _ := strconv.Atoi(idNumber[10:12])
	day, _ := strconv.Atoi(idNumber[12:14])

	switch {
	case (month == 3 && day >= 21) || (month == 4 && day <= 19):
		return "白羊座"
	case (month == 4 && day >= 20) || (month == 5 && day <= 20):
		return "金牛座"
	case (month == 5 && day >= 21) || (month == 6 && day <= 21):
		return "双子座"
	case (month == 6 && day >= 22) || (month == 7 && day <= 22):
		return "巨蟹座"
	case (month == 7 && day >= 23) || (month == 8 && day <= 22):
		return "狮子座"
	case (month == 8 && day >= 23) || (month == 9 && day <= 22):
		return "处女座"
	case (month == 9 && day >= 23) || (month == 10 && day <= 23):
		return "天秤座"
	case (month == 10 && day >= 24) || (month == 11 && day <= 22):
		return "天蝎座"
	case (month == 11 && day >= 23) || (month == 12 && day <= 21):
		return "射手座"
	case (month == 12 && day >= 22) || (month == 1 && day <= 19):
		return "摩羯座"
	case (month == 1 && day >= 20) || (month == 2 && day <= 18):
		return "水瓶座"
	default: // (month == 2 && day >= 19) || (month == 3 && day <= 20)
		return "双鱼座"
	}
}
