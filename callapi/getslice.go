package callapi

import (
	"strings"
	"time"
)

func getslice(inputvars map[string]interface{}, index int, current_methodinput methodinput) ([]interface{}, int) {
	var myargs []interface{}
	var length int
	for _, y := range inputvars {
		s, ok := y.(string)
		if ok == true {
			if strings.Contains(s, "method1") {
				x := strings.Split(s, ".")
				if len(x) != 0 {
					if strings.Contains(x[1], "array") {
						var my_methodinputer methodinputer
						my_methodinputer = &current_methodinput
						k, i := my_methodinputer.getarray(x[len(x)-1], "method1", index)
						if k != nil {
							length = i
							myargs = append(myargs, k)
						}
					} else {
						var my_methodinputer methodinputer
						my_methodinputer = &current_methodinput
						k, i := my_methodinputer.getvalue(x[len(x)-1], "method1", index)
						if k != nil {
							length = i
							myargs = append(myargs, k)
						}
					}
				}
			} else if strings.Contains(s, "method2") {
				x := strings.Split(s, ".")
				if len(x) != 0 {
					if strings.Contains(x[1], "array") {
						var my_methodinputer methodinputer
						my_methodinputer = &current_methodinput
						k, i := my_methodinputer.getarray(x[len(x)-1], "method2", index)
						length = i
						myargs = append(myargs, k)
					} else {
						var my_methodinputer methodinputer
						my_methodinputer = &current_methodinput
						k, i := my_methodinputer.getvalue(x[len(x)-1], "method2", index)
						length = i
						myargs = append(myargs, k)
					}
				}
			} else if strings.Contains(s, "bool") {
				x := strings.Split(s, ".")
				//fmt.Printf("x is: %v\n", x)
				if len(x) != 0 {
					if x[len(x)-1] == "true" || x[len(x)-1] == "1" {
						k := true
						myargs = append(myargs, k)
					} else {
						k := false
						myargs = append(myargs, k)
					}
				}
			} else if strings.Contains(s, "datetime") {
				x := strings.Split(s, ".")
				if len(x) != 0 {
					const layout = "2006-01-02T15:04:05"
					k, _ := time.Parse(layout, x[len(x)-1])
					myargs = append(myargs, k)
				}
			} else {
				myargs = append(myargs, s)
			}
		} else {
			myargs = append(myargs, y)
		}
	}
	return myargs, length
}
