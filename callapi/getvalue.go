package callapi

import (
	"fmt"
	"log"

	"github.com/bjin01/go-xmlrpc"
)

type Printvalue interface {
	Printmethod1(cfg *C) error
	Printmethod2(cfg *C) error
	Printfinalmethod(cfg *C) error
}

type R struct {
	U xmlrpc.Value
}

type Result struct {
	Intmap      map[string]int
	Stringmap   map[string]string
	Datetimemap map[string]interface{}
	Boolmap     map[string]bool
	Innerlist   map[string]interface{}
}

func (r *R) Printfinalmethod(cfg *C) error {

	return nil
}

func (r *R) Printmethod2(cfg *C) error {

	return nil
}

func (r *R) Printmethod1(cfg *C) error {
	log.Printf("method1 output: \n")
	fmt.Printf("%v\n", r.U.Values())
	return nil
}

func (methodvalues *methodinput) getvalue(field string, methodname string, index int) (interface{}, int) {
	var val_length int
	var getval_value interface{}
	//fmt.Printf("\nmethodname: %v, field: %v: has %v results, index is %v\n", methodname, field, len(methodvalues.result.Values()), index)
	if methodname == "method1" {
		val_length = len(methodvalues.result.Values())
		getval_value = GetVal(methodvalues.result, field, index)
		fmt.Printf("\n%v, %v: %v, index: %v\n", methodname, field, getval_value, index)
	}

	if methodname == "method2" {
		fmt.Println("getvalue method2 is here")
		val_length = len(methodvalues.result2.Values())
		getval_value = GetVal(methodvalues.result2, field, index)
		fmt.Printf("\n%v, %v: %v, index: %v\n", methodname, field, getval_value, index)
	}

	/*if methodname == "finalmethod" {
		val_length = len(methodvalues.result2.Values())
		getval_value = GetVal(methodvalues.result2, field, index)
		fmt.Printf("\n%v, %v: %v, index: %v\n", methodname, field, getval_value, index)
	}*/

	return getval_value, val_length
}

func (methodvalues *methodinput) getarray(field string, methodname string, index int) (interface{}, int) {
	var val_length int
	var temp_val interface{}
	var getval_value []interface{}
	if methodname == "method1" {
		val_length = len(methodvalues.result.Values())
		temp_val = GetVal(methodvalues.result, field, index)
		getval_value = append(getval_value, temp_val)
		fmt.Printf("\n%v, %v: %v, index: %v\n", methodname, field, getval_value, index)
	}

	if methodname == "method2" {
		fmt.Println("getarray method2 is here")
		val_length = len(methodvalues.result2.Values())
		if val_length != 0 {
			for i := 0; i < val_length; i++ {
				temp_val = GetVal(methodvalues.result2, field, i)
				getval_value = append(getval_value, temp_val)
			}
		} else {
			temp_val = GetVal(methodvalues.result2, field, index)
			getval_value = append(getval_value, temp_val)
		}
		fmt.Printf("\n%v, %v: %v, index: %v\n", methodname, field, getval_value, index)
	}

	/*if methodname == "finalmethod" {
		val_length = len(methodvalues.result2.Values())
		getval_value = GetVal(methodvalues.result2, field, index)
		fmt.Printf("\n%v, %v: %v, index: %v\n", methodname, field, getval_value, index)
	}*/

	return getval_value, val_length
}

func GetVal(v xmlrpc.Value, searchfield string, index int) interface{} {
	fmt.Printf("lets see here %v %v\n", len(v.Values()), len(v.Members()))
	var return_val interface{}
	if len(v.Values()) == 0 && len(v.Members()) != 0 {
		x := v.Members()
		for _, v := range x {
			fmt.Printf("will print the member value here. %v\n", v.Value())
		}

	}

	if len(v.Values()) == 0 && len(v.Members()) == 0 {
		if searchfield != "" {
			fmt.Printf("will print the simiple value here.")

		}
	}

	if len(v.Values()) != 0 {

		x := v.Values()[index].Members()
		if len(x) == 0 {
			return_val = GetVal3(v.Values()[index], searchfield)
			return return_val
		}

		for _, v := range x {
			if v.Name() == searchfield {
				return_val = GetVal3(v.Value(), searchfield)
				return return_val
			}
		}
	}
	return return_val
}

func GetVal3(v xmlrpc.Value, searchfield string) interface{} {
	z := v.Kind()
	y := v
	var return_val interface{}

	switch f := z; f {
	case 1:
		GetMembers(y.Members(), searchfield)
	case 2:
		fmt.Printf("\t%v\n", y.Bytes())
	case 3:
		fmt.Printf("\t%v\n", y.Bool())
	case 4:
		//fmt.Printf("\t%s\n", y.Time())
		return y.Time
	case 5:
		fmt.Printf("%v\n", y.Double())
	case 6: //this is a int type
		return y.Int()

	case 7: //this is a string type
		return y.Text()

	case 8: //this is a member type

		return_val = GetVal(y, searchfield, 0)
	default:

		return_val = GetVal(y, searchfield, 0)
	}
	return return_val
}

func GetMembers(x []xmlrpc.Member, searchfield string) {
	//fmt.Printf("func GetMembers: %v\n", x)

	for _, y := range x {
		GetVal3(y.Value(), searchfield)

	}
}
