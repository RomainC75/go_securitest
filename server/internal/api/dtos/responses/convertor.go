package dto_res

import (
	"database/sql"
	"fmt"
	"reflect"
	db "server/db/sqlc"
	shared_dto "shared/dto"
	"shared/types"
	shared_utils "shared/utils"
)

const dtotag = "dto"

type IAdress struct {
	Number int
	Street string
}

type ITest struct {
	Name   string
	Adress IAdress
}

func convertStruct(inter interface{}) any {
	// vStruct := ITest{
	// 	Name: "my name",
	// 	Adress: IAdress{
	// 		Number: 12,
	// 		Street: "Sauffroy",
	// 	},
	// }

	// shared_utils.PrettyDisplay("++++", inter)

	// v := reflect.ValueOf(inter)
	// fmt.Println("v :  ", v.Type())
	// for i := 0; i < v.NumField(); i++ {
	// 	fmt.Println("->", v.Type().Field(i))
	// }

	// var v float64 = 3.14

	// analyse
	// tp := reflect.TypeOf(inter)
	v := reflect.ValueOf(inter).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch field.Kind() {
		case reflect.Struct:
			convertStruct(field)
		case reflect.Ptr:
			if !field.IsNil() {
				convertStruct(field.Elem())
			}
		case reflect.Interface:
			if !field.IsNil() {
				convertStruct(field.Elem())
			}
		default:
			if field.Type() == reflect.TypeOf(types.SNullString{}) {
				field.Set(reflect.ValueOf(convertSNullStringToNullString(field.Interface().(types.SNullString))))
			} else if field.Type() == reflect.TypeOf(types.SNullInt32{}) {
				field.Set(reflect.ValueOf(convertSNullInt32ToNullInt32(field.Interface().(types.SNullInt32))))
			}
		}
	}

	return inter
}

func convertSNullStringToNullString(s types.SNullString) sql.NullString {
	return sql.NullString{
		String: s.String,
		Valid:  s.Valid,
	}
}

func convertSNullInt32ToNullInt32(s types.SNullInt32) sql.NullInt32 {
	return sql.NullInt32{
		Int32: s.Int32,
		Valid: s.Valid,
	}
}

func ToFullPortTestScenarioReq(inter db.ListScansByUserRow) shared_dto.FullPortTestScenarioReq {
	// fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(inter)
	t := reflect.TypeOf(inter)
	fmt.Println("typeof : ", t)
	fmt.Println("value of : ", v)
	shared_utils.PrettyDisplay("valueOf", v)
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		t = reflect.TypeOf(v.Field(i))
		fmt.Println("t : ", t)
		shared_utils.PrettyDisplay("fieldInfo : ", fieldInfo)
		// fmt.Println("--", reflect.TypeOf(v.Field(i)), reflect.ValueOf(v.Field(i)))
		// reflect.Struct
		// tag := fieldInfo.Tag
		// name := tag.Get(dtotag)
		// if name == "" {
		// 	name = strings.ToLower(fieldInfo.Name)
		// }
		// fields[name] = v.Field(i)
		// }
	}
	return shared_dto.FullPortTestScenarioReq{}
}
