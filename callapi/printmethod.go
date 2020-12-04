package callapi

import (
	"fmt"

	"github.com/bjin01/go-xmlrpc"
)

type PrintValuer interface {
	Printvalue(cfg *C, methodname string)
	Printvalue2(fields []string)
}

func (v *R) PrintValue(cfg *C, methodname string) {
	if methodname != "" && methodname == "method1" {
		outvars := cfg.Cfg.Method1.Outvariables
		v.Printvalue2(outvars)
		fmt.Println("----------------------------")
	}

	if methodname != "" && methodname == "method2" {
		outvars := cfg.Cfg.Method2.Outvariables
		v.Printvalue2(outvars)
		fmt.Println("----------------------------")
	}

	if methodname != "" && methodname == "finalmethod" {
		outvars := cfg.Cfg.Finalmethod.Outvariables
		v.Printvalue2(outvars)
		fmt.Println("----------------------------")
	}

}

func (v *R) Printvalue2(fields []string) {
	if len(v.U.Values()) == 0 && len(v.U.Members()) != 0 {
		x := v.U.Members()
		for _, v := range x {
			for _, searchfield := range fields {
				if v.Name() == searchfield {
					fmt.Printf("%v: ", searchfield)
					printval3(v.Value(), searchfield)
				}
			}
		}
	}

	if len(v.U.Values()) == 0 && len(v.U.Members()) == 0 {
		if len(fields) != 0 {
			for _, searchfield := range fields {
				fmt.Printf("%v: ", searchfield)
				printval3(v.U, searchfield)

			}
		}
	}

	if len(v.U.Values()) != 0 {
		for k := 0; k < len(v.U.Values()); k++ {
			x := v.U.Values()[k].Members()
			if len(x) == 0 {
				if len(fields) != 0 {
					for _, searchfield := range fields {
						fmt.Printf("%v: ", searchfield)
						printval3(v.U.Values()[k], searchfield)
					}
				}
			} else {
				if len(fields) != 0 {
					for _, searchfield := range fields {
						for _, i := range x {
							if i.Name() == searchfield {
								fmt.Printf("%v: ", searchfield)
								printval3(i.Value(), searchfield)
							}
						}
					}
				}
			}

		}

	}

}

func printval3(v xmlrpc.Value, searchfield string) {
	z := v.Kind()
	y := v

	switch f := z; f {
	case 1:
		fmt.Printf("case 1: %v\n", y.Kind())
	case 2:
		fmt.Printf("\t%v\n", y.Bytes())
	case 3:
		fmt.Printf("\t%v\n", y.Bool())
	case 4:
		fmt.Printf("\t%s\n", y.Time())
	case 5:
		fmt.Printf("%v\n", y.Double())
	case 6: //this is a int type
		fmt.Printf("%v\n", y.Int())
	case 7: //this is a string type
		fmt.Printf("%v\n", y.Text())
	case 8: //this is a member type
		fmt.Printf("case 8: %v\n", y.Kind())
	default:
		fmt.Printf("case default: %v\n", y.Kind())
	}
}
