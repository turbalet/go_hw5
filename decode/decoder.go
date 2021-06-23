package decode

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
)


type User struct {
	ID      IntVal   `json:"id" xml:"id"`
	Address Address `json:"address" xml:"address"`
	Age     IntVal     `json:"age" xml:"age"`
}

type IntVal struct {
	Val int64
}

type Address struct {
	CityID IntVal  `json:"city_id" xml:"city_id"`
	Street string `json:"street" xml:"street"`
}

type Users struct {
	Users   []User `xml:"users"`
}

func (i *IntVal) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &i.Val); err != nil {
		var str string
		err = json.Unmarshal(b, &str)
		res, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		i.Val = int64(res)
	}
	return nil
}

func (i *IntVal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if err := d.DecodeElement(&i, &start); err != nil {
		var s string
		if err := d.DecodeElement(&s, &start); err != nil {
			return err
		}
		res, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		i.Val = int64(res)
	}
	return nil
}