package models

import "time"

type Province struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Regency struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Province  Province  `json:"province,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type District struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Regency   Regency   `json:"regency,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type People struct {
	ID         string    `json:"id,omitempty"`
	NIK        string    `json:"nik,omitempty"`
	Name       string    `json:"name,omitempty"`
	Gender     string    `json:"gender,omitempty"`
	DOB        time.Time `json:"dob,omitempty"`
	POB        string    `json:"pob,omitempty"`
	ProvinceID string    `json:"province_id,omitempty"`
	Province   Province  `json:"province,omitempty"`
	RegencyID  string    `json:"regency_id,omitempty"`
	Regency    Regency   `json:"regency,omitempty"`
	DistrictID string    `json:"district_id,omitempty"`
	District   District  `json:"district,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}
