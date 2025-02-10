package repository

import (
	"id-information/domain"
)

// GetAddress 根据身份证号码获取地址
//
// 参数:
//   - idNumber: 身份证号码
//
// 返回值:
//   - *domain.Address: 地址信息
//   - *domain.ApiError: 错误信息
func GetAddress(idNumber string) (*domain.Address, *domain.ApiError) {
	area := AreaRepository[idNumber[:6]]
	if area == nil {
		// 只需要在这里判断是否存在即可,如果这里存在,那么后面的city和province一定存在
		return nil, domain.ErrInvalidIDNumber
	}
	city := CityRepository[idNumber[:4]]
	province := ProvinceRepository[idNumber[:2]]
	return &domain.Address{
		Province: province.Name,
		City:     city.Name,
		Area:     area.Name,
	}, nil
}
