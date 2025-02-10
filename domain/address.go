package domain

type Province struct {
	Code string
	Name string
}

type City struct {
	Code         string
	Name         string
	ProvinceCode string
}

type Area struct {
	Code         string
	Name         string
	CityCode     string
	ProvinceCode string
}
