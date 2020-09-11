package callapi

import (
	"log"
)

/* func splitinputs(v interface{}) (interface{}, error) {

} */

func (Cfg *C) getinputargs(methodname string, index int, current_methodinput methodinput) ([]interface{}, int) {
	var myargs []interface{}
	var inputvars map[string]interface{}
	var length int

	if methodname == "method1" {
		inputvars = Cfg.Cfg.Method1.InputVars
		if len(inputvars) != 0 {
			myargs, _ = getslice(inputvars, index, current_methodinput)
		}

	} else if methodname == "method2" {
		inputvars = Cfg.Cfg.Method2.InputVars
		if len(inputvars) != 0 {
			myargs, length = getslice(inputvars, index, current_methodinput)
		}

	} else if methodname == "finalmethod" {
		inputvars = Cfg.Cfg.Finalmethod.InputVars
		if len(inputvars) != 0 {
			myargs, _ = getslice(inputvars, index, current_methodinput)
		}

	} else {
		log.Fatal("no method name provided so I cannot continue to get myargs.")
	}

	if myargs != nil {
		return myargs, length
	}
	return nil, 0
}
