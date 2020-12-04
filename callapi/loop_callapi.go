package callapi

import (
	"fmt"

	"github.com/bjin01/go-xmlrpc"
)

func Call_api_2(Cfg *C, methodname string, index int, current_methodinput methodinput, u xmlrpc.Value) {

	//var mycall Caller = Cfg
	var print_u Printvalue

	if methodname == "method2" {
		if current_methodinput.result != nil && u != nil {
			if Cfg.Cfg.Finalmethod.Methodname != "" {
				method := "finalmethod"
				current_methodinput = methodinput{
					apicallname:     Cfg.Cfg.Finalmethod.Methodname,
					outputvariables: Cfg.Cfg.Finalmethod.Outvariables,
					result:          current_methodinput.result,
					result2:         u,
				}
				var myargs []interface{}
				myargs, _ = Cfg.getinputargs(method, index, current_methodinput)
				u, err := Cfg.runapi(method, myargs)
				//check(err)
				if err != nil {
					fmt.Println(err)
				}
				if u != nil {
					result := R{
						U: u,
					}
					print_u = &result

					err := print_u.Printfinalmethod(Cfg)
					//check(err)
					if err != nil {
						fmt.Println(err)
					}
				}
			} else {
				return
			}
		}
	}
}
