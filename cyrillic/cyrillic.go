package cyrillic

import (
	"reflect"
	"unicode"
)

func RemoveKirill(s string) string {
	var res string
	for _, r := range s {
		if !unicode.Is(unicode.Cyrillic, r) {
			res += string(r)
		}
	}
	return res
}

func ListFields(a interface{}) {
	v := reflect.ValueOf(a).Elem()
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Kind() == reflect.String {
			v.Field(i).SetString(RemoveKirill(v.Field(i).Interface().(string)))
			//fmt.Println(v.Field(i).String())
		} else if v.Field(i).Kind() == reflect.Ptr && v.Field(i).Elem().Kind() == reflect.String {
			val := v.Field(i).Elem().String()
			//s := RemoveKirill(v.Field(i).Elem().Interface().(string))
			s := RemoveKirill(val)
			v.Field(i).Elem().SetString(s)
		} else if v.Field(i).Kind() == reflect.Struct {
			ListFields(v.Field(i).Addr().Interface())
		}
	}

}
