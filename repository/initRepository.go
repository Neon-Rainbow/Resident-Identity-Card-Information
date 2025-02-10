package repository

import (
	"encoding/csv"
	"id-information/domain"
	"log"
	"os"
)

var (
	ProvinceRepository map[string]*domain.Province
	CityRepository     map[string]*domain.City
	AreaRepository     map[string]*domain.Area
)

// InitRepository 初始化数据库
func InitRepository() {
	var err error
	// 读取省份数据
	ProvinceRepository, err = readDataFromCsv[*domain.Province]("./data/provinces.csv")
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
		return
	}
	// 读取城市数据
	CityRepository, err = readDataFromCsv[*domain.City]("./data/cities.csv")
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
	}
	// 读取区县数据
	AreaRepository, err = readDataFromCsv[*domain.Area]("./data/areas.csv")
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
	}
	return
}

// readDataFromCsv 从 CSV 文件中读取数据
//
// 参数:
//   - filepath: 文件路径
//
// 返回值:
//   - map[string]T: 数据映射
//   - error: 错误信息
func readDataFromCsv[T any](filepath string) (map[string]T, error) {
	// 打开文件
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 创建 CSV reader
	reader := csv.NewReader(file)

	// 跳过表头
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	// 创建结果 map
	result := make(map[string]T)

	// 读取数据
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		// 根据不同类型创建对应的结构体
		var item T
		switch any(item).(type) {
		case *domain.Province:
			province := &domain.Province{
				Code: record[0],
				Name: record[1],
			}
			result[province.Code] = any(province).(T)
		case *domain.City:
			city := &domain.City{
				Code:         record[0],
				Name:         record[1],
				ProvinceCode: record[2],
			}
			result[city.Code] = any(city).(T)
		case *domain.Area:
			area := &domain.Area{
				Code:         record[0],
				Name:         record[1],
				CityCode:     record[2],
				ProvinceCode: record[3],
			}
			result[area.Code] = any(area).(T)
		}
	}

	return result, nil
}
