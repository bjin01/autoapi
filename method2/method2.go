package method2

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/bjin01/autoapi/filters"
	"github.com/bjin01/autoapi/getyaml"
	"github.com/bjin01/autoapi/printresult"
	"github.com/bjin01/go-xmlrpc"
)

type Result struct {
	Intmap    []map[string]int
	Stringmap []map[string]string
	Datemap   []map[string]interface{}
}

type ResultinMap struct {
	IdMap map[interface{}]interface{}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkprint(err error, myargs []interface{}) {
	if err != nil {
		fmt.Printf("\n---uups---\n")
		log.Println(err)
		if len(myargs) != 0 {
			for _, v := range myargs {
				s := fmt.Sprintf("%v", v)
				fmt.Printf("\t%s\n", s)
			}
		}

	}
}

func Method2(cfg *getyaml.Config, url string, user string, password string, method string,
	inputmaps map[string]interface{}, searchfields []string, result *Result, resultsmethod1 *printresult.PrintResults) {
	fmt.Printf("\nCalling %v...\n", method)

	var inputmapvalslice []interface{}
	client := xmlrpc.NewClient(url)

	f, err := client.Call("auth.login", user, password)
	check(err)

	if len(inputmaps) != 0 {
		keys := []string{}
		for key, _ := range inputmaps {
			keys = append(keys, key)
		}

		sort.Strings(keys)

		for _, v := range keys {

			inputmapvalslice = append(inputmapvalslice, inputmaps[v])

		}
	}

	myargs, loopnum := createinputargs(inputmapvalslice, resultsmethod1, 0)

	if loopnum > 0 {
		for i := 0; i < loopnum; i++ {

			myargs, _ := createinputargs(inputmapvalslice, resultsmethod1, i)
			//fmt.Printf("my args are: %v\n", myargs)

			callapi(cfg, client, method, f.Text(), searchfields, myargs, result)

		}
	} else {
		callapi(cfg, client, method, f.Text(), searchfields, myargs, result)
	}

	_, err = client.Call("auth.logout", f.Text())
	check(err)

}

func callapi(cfg *getyaml.Config, client xmlrpc.Client, method string, sessionkey string,
	searchfields []string, myargs []interface{}, result *Result) {
	intlist := []int{}
	strlist := []string{}
	var datelist []interface{}

	switch {
	case len(myargs) == 0:
		u, err := client.Call(method, sessionkey)
		checkprint(err, myargs)
		if u != nil {
			//adding filter feature here
			if cfg.Method2.Filters != nil {
				filters.ApplyFilter(*cfg, u, "method2")
			} else {
				GetVal(u, searchfields, result, datelist, intlist, strlist)
			}

		}
	case len(myargs) == 1:
		u, err := client.Call(method, sessionkey, myargs[0])
		checkprint(err, myargs)
		if u != nil {
			//adding filter feature here
			if cfg.Method2.Filters != nil {
				filters.ApplyFilter(*cfg, u, "method2")
			} else {
				GetVal(u, searchfields, result, datelist, intlist, strlist)
			}

		}
	case len(myargs) == 2:
		u, err := client.Call(method, sessionkey, myargs[0], myargs[1])
		checkprint(err, myargs)
		if u != nil {
			//adding filter feature here
			if cfg.Method2.Filters != nil {
				filters.ApplyFilter(*cfg, u, "method2")
			} else {
				GetVal(u, searchfields, result, datelist, intlist, strlist)
			}

		}
	case len(myargs) == 3:
		u, err := client.Call(method, sessionkey, myargs[0], myargs[1], myargs[2])
		checkprint(err, myargs)
		if u != nil {
			//adding filter feature here
			if cfg.Method2.Filters != nil {
				filters.ApplyFilter(*cfg, u, "method2")
			} else {
				GetVal(u, searchfields, result, datelist, intlist, strlist)
			}

		}
	case len(myargs) == 4:
		u, err := client.Call(method, sessionkey, myargs[0], myargs[1], myargs[2], myargs[3])
		checkprint(err, myargs)
		if u != nil {
			//adding filter feature here
			if cfg.Method2.Filters != nil {
				filters.ApplyFilter(*cfg, u, "method2")
			} else {
				GetVal(u, searchfields, result, datelist, intlist, strlist)
			}

		}
	case len(myargs) == 5:
		u, err := client.Call(method, sessionkey, myargs[0], myargs[1], myargs[2], myargs[3], myargs[4])
		checkprint(err, myargs)
		if u != nil {
			//adding filter feature here
			if cfg.Method2.Filters != nil {
				filters.ApplyFilter(*cfg, u, "method2")
			} else {
				GetVal(u, searchfields, result, datelist, intlist, strlist)
			}

		}

	}
}

func createinputargs(inputmapvalslice []interface{}, resultsmethod1 *printresult.PrintResults, h int) ([]interface{}, int) {
	myargs := make([]interface{}, 0)

	var loopnum int

	for i := 0; i < len(inputmapvalslice); i++ {
		s, num := splitinputvar(inputmapvalslice[i], resultsmethod1, h)
		myargs = append(myargs, s)
		if num > 0 {
			loopnum = num
		}

	}

	return myargs, loopnum

}

func splitinputvar(v interface{}, resultsmethod1 *printresult.PrintResults, h int) (interface{}, int) {
	s, ok := v.(string)
	//fmt.Printf("\t%v\n", s)

	if ok == true {
		if strings.Contains(s, "method1") {
			x := strings.Split(s, ".")

			if len(x) != 0 {

				k, num := Getfrommethod1(x[len(x)-1], resultsmethod1, h)

				return k, num
			}
		} else if strings.Contains(s, "bool") {
			x :=    strings.Split(s, ".")
			//fmt.Printf("x is: %v\n", x)
			if len(x) != 0 {
				if x[len(x)-1] == "true" || x[len(x)-1] == "1" {

					k := true

					return k, 0
				}
				if x[len(x)-1] == "false" || x[len(x)-1] == "0" {
					k := false
					return k, 0
				}

			}
		} else if strings.Contains(s, "datetime") {
			x := strings.Split(s, ".")
			//fmt.Printf("x is: %v\n", x)
			if len(x) != 0 {

				const layout = "2006-01-02T15:04:05"
				k, _ := time.Parse(layout, x[len(x)-1])

				return k, 0

			}
		} else {
			return v, 0
		}
	}
	return v, 0
}

func Getfrommethod1(s string, resmethod1 *printresult.PrintResults, h int) (interface{}, int) {
	var x interface{}
	var loopnum int

	for k, v := range resmethod1.Intmap {
		if k == s {

			x = v[h]
			loopnum = len(v)
			return x, loopnum

		}

	}

	for k, v := range resmethod1.Stringmap {
		if k == s {

			x = v[h]
			loopnum = len(v)
			return x, loopnum

		}

	}

	for k, v := range resmethod1.Datetimemap {
		if k == s {

			x = v[h]
			loopnum = len(v)
			return x, loopnum

		}

	}
	return x, 0
}

func GetVal(v xmlrpc.Value, searchfields []string, result *Result,
	datelist []interface{}, intlist []int, strlist []string) {
	/* fmt.Printf("lets see here %v %v\n", len(v.Values()), len(v.Members())) */

	if len(v.Values()) == 0 && len(v.Members()) != 0 {
		x := v.Members()
		for _, v := range x {
			if len(searchfields) != 0 {
				for h := 0; h < len(searchfields); h++ {
					if v.Name() == searchfields[h] {
						fmt.Printf("\t%v:", searchfields[h])
						if len(v.Value().Values()) != 0 {
							for i := 0; i < len(v.Value().Values()); i++ {
								GetVal3(v.Value().Values()[i], searchfields, v.Name(), result, datelist, intlist, strlist)
							}
						} else {
							GetVal3(v.Value(), searchfields, v.Name(), result, datelist, intlist, strlist)
						}

					}
				}
			} else {
				GetVal3(v.Value(), searchfields, "nil", result, datelist, intlist, strlist)
			}
		}
	}

	if len(v.Values()) == 0 && len(v.Members()) == 0 {
		if len(searchfields) != 0 {
			for h := 0; h < len(searchfields); h++ {

				GetVal3(v, searchfields, searchfields[h], result, datelist, intlist, strlist)

			}
		}
	}

	if len(v.Values()) != 0 {
		//fmt.Printf("in values\n")
		//fmt.Printf("%v\n", searchfields)
		for k := 0; k < len(v.Values()); k++ {

			x := v.Values()[k].Members()
			if len(x) == 0 {
				GetVal3(v.Values()[k], searchfields, searchfields[0], result, datelist, intlist, strlist)
				continue
			}
			//fmt.Printf("---------------------\n")
			for _, v := range x {
				if len(searchfields) != 0 {

					for h := 0; h < len(searchfields); h++ {
						//for g, h := range searchfields {
						if v.Name() == searchfields[h] {
							//fmt.Printf("%d - %v:", h, searchfields[h])
							if len(v.Value().Values()) != 0 {
								for i := 0; i < len(v.Value().Values()); i++ {
									GetVal3(v.Value().Values()[i], searchfields, v.Name(), result, datelist, intlist, strlist)
								}
							} else {
								//fmt.Printf("else: %v\n", v.Value())
								GetVal3(v.Value(), searchfields, v.Name(), result, datelist, intlist, strlist)
							}

						}
					}
				} else {
					GetVal3(v.Value(), searchfields, "nil", result, datelist, intlist, strlist)
				}
			}
		}

	}

}

func GetVal3(v xmlrpc.Value, searchfields []string, matchfield string,
	result *Result, datelist []interface{},
	intlist []int, strlist []string) {
	z := v.Kind()

	// if z or v.Kind() is return invalid or 0 than we want to exit this func otherwise endless loop happens.
	if z == 0 {
		return
	}
	y := v

	switch f := z; f {
	case 1:
		GetMembers(y.Members(), searchfields, result, datelist, intlist, strlist)
	case 2:
		fmt.Printf("\t%v\n", y.Bytes())
	case 3:
		fmt.Printf("\t%v\n", y.Bool())
	case 4:
		//fmt.Printf("\t%s\n", y.Time())
		datetimemap := map[string]interface{}{
			matchfield: y.Time(),
		}
		result.Datemap = append(result.Datemap, datetimemap)
		datelist = append(datelist, y.Time())
	case 5:
		fmt.Printf("%v\n", y.Double())
	case 6: //this is a int type

		intmap := map[string]int{
			matchfield: y.Int(),
		}
		result.Intmap = append(result.Intmap, intmap)

	case 7: //this is a string type

		stringmap := map[string]string{
			matchfield: y.Text(),
		}
		result.Stringmap = append(result.Stringmap, stringmap)
		strlist = append(strlist, y.Text())

	case 8: //this is a member type

		GetVal(y, searchfields, result, datelist, intlist, strlist)
	default:

		GetVal(y, searchfields, result, datelist, intlist, strlist)
	}
}

func GetMembers(x []xmlrpc.Member, searchfields []string, result *Result, datelist []interface{},
	intlist []int, strlist []string) {
	//fmt.Printf("func GetMembers: %v\n", x)

	for _, y := range x {
		if len(searchfields) != 0 {
			for h := 0; h < len(searchfields); h++ {
				if y.Name() == searchfields[h] {

					if len(y.Value().Values()) != 0 {
						for i := 0; i < len(y.Value().Values()); i++ {
							GetVal3(y.Value().Values()[i], searchfields, y.Name(), result, datelist, intlist, strlist)
						}
					} else {
						GetVal3(y.Value(), searchfields, y.Name(), result, datelist, intlist, strlist)
					}
				}
			}
		} else {
			GetVal3(y.Value(), searchfields, "nil", result, datelist, intlist, strlist)
		}
	}
}
