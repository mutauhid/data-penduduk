package models

import "time"

type Province struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Regency struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	ProvinceID string    `json:"province_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type District struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	RegencyID string    `json:"regency_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type People struct {
	ID         string    `json:"id"`
	NIK        string    `json:"nik"`
	Name       string    `json:"name"`
	Gender     string    `json:"gender"`
	DOB        time.Time `json:"dob"`
	POB        string    `json:"pob"`
	ProvinceID string    `json:"province_id"`
	RegencyID  string    `json:"regency_id"`
	DistrictID string    `json:"district_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
