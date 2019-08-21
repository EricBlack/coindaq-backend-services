package model

import "bx.com/user-service/bxgo"

type Country struct {
	Id   int64
	CountryNB string 	`xorm:"country_nb text notnull"`
	Mark string			`xorm:"mark text notnull"`
	EName string		`xorm:"ename text notnull"`
	CName string 		`xorm:"cname text notnull"`
}

func (c Country) TableName() string {
	return "country"
}

func QueryCountryById(id int64) (Country ,error) {
	country := Country{}
	_, err := bxgo.OrmEngin.Id(id).Get(&country)

	return country, err
}

func QueryCountryByMark(mark string) (Country, error) {
	country := Country{}
	_, err := bxgo.OrmEngin.Where("mark=? ", mark).Get(&country)

	return country, err
}

func QueryAllCountry() ([]Country, error) {
	countryList := []Country{}
	err := bxgo.OrmEngin.Find(&countryList)

	return countryList, err
}
