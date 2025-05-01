package main

type University struct {
	ID                uint   `gorm:"primaryKey" json:"id"`
	Name              string `json:"name"`
	Category          string `json:"category"`
	Website           string `json:"website"`
	DetailURL         string `json:"detail_url"`
	YearEstablished   string `json:"year_established"`
	PermanentCampus   string `json:"permanent_campus"`
	ViceChancellor    string `json:"vice_chancellor"`
	ProViceChancellor string `json:"pro_vice_chancellor"`
	Treasurer         string `json:"treasurer"`
	Registrar         string `json:"registrar"`
	Contact           string `json:"contact"`
	Email             string `json:"email"`
	Telephone         string `json:"telephone"`
	Fax               string `json:"fax"`
}
