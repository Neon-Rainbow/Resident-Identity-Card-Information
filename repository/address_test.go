package repository

import (
	"id-information/domain"
	"testing"
)

func TestGetAddress(t *testing.T) {
	// 初始化测试数据
	ProvinceRepository = map[string]*domain.Province{
		"11": {Name: "北京市"},
		"32": {Name: "江苏省"},
	}

	CityRepository = map[string]*domain.City{
		"1101": {Name: "北京市"},
		"3201": {Name: "南京市"},
	}

	AreaRepository = map[string]*domain.Area{
		"110101": {Name: "东城区"},
		"320102": {Name: "玄武区"},
	}

	tests := []struct {
		name     string
		idNumber string
		want     *domain.Address
		wantErr  *domain.ApiError
	}{
		{
			name:     "正常地址解析-北京",
			idNumber: "110101199003078888",
			want: &domain.Address{
				Province: "北京市",
				City:     "北京市",
				Area:     "东城区",
			},
			wantErr: nil,
		},
		{
			name:     "正常地址解析-南京",
			idNumber: "320102199003078888",
			want: &domain.Address{
				Province: "江苏省",
				City:     "南京市",
				Area:     "玄武区",
			},
			wantErr: nil,
		},
		{
			name:     "非法身份证号-地址不存在",
			idNumber: "888888199003078888",
			want:     nil,
			wantErr:  domain.ErrInvalidIDNumber,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAddress(tt.idNumber)
			if err != tt.wantErr {
				t.Errorf("GetAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				if got.Province != tt.want.Province {
					t.Errorf("GetAddress() Province = %v, want %v", got.Province, tt.want.Province)
				}
				if got.City != tt.want.City {
					t.Errorf("GetAddress() City = %v, want %v", got.City, tt.want.City)
				}
				if got.Area != tt.want.Area {
					t.Errorf("GetAddress() Area = %v, want %v", got.Area, tt.want.Area)
				}
			}
		})
	}
}
