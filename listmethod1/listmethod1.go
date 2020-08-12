package listmethod1

import (
	"fmt"
	"log"
	"sort"

	"github.com/bjin01/go-xmlrpc"
)

type Result struct {
	Intmap    []map[string]int
	Stringmap []map[string]string
	Datemap   []map[string]interface{}
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

func Listmethod1(url string, user string, password string, listmethod string,
	inputmaps map[string]interface{}, searchfields []string, result *Result) {
	fmt.Printf("Calling %v...\n", listmethod)
	var myargs []interface{}

	client := xmlrpc.NewClient(url)

	f, err := client.Call("auth.login", user, password)
	check(err)

	intlist := []int{}
	strlist := []string{}
	var datelist []interface{}

	if len(inputmaps) != 0 {
		keys := []string{}
		for key, _ := range inputmaps {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, v := range keys {
			myargs = append(myargs, inputmaps[v])
		}
		switch {
		case len(myargs) == 0:
			u, err := client.Call(listmethod, f.Text())
			checkprint(err, myargs)
			if u != nil {
				GetVal(u, searchfields, result, datelist, intlist, strlist)
			}
		case len(myargs) == 1:
			u, err := client.Call(listmethod, f.Text(), myargs[0])
			checkprint(err, myargs)
			if u != nil {
				GetVal(u, searchfields, result, datelist, intlist, strlist)
			}
		case len(myargs) == 2:
			u, err := client.Call(listmethod, f.Text(), myargs[0], myargs[1])
			checkprint(err, myargs)
			if u != nil {
				GetVal(u, searchfields, result, datelist, intlist, strlist)
			}
		case len(myargs) == 3:
			u, err := client.Call(listmethod, f.Text(), myargs[0], myargs[1], myargs[2])
			checkprint(err, myargs)
			if u != nil {
				GetVal(u, searchfields, result, datelist, intlist, strlist)
			}
		case len(myargs) == 4:
			u, err := client.Call(listmethod, f.Text(), myargs[0], myargs[1], myargs[2], myargs[3])
			checkprint(err, myargs)
			if u != nil {
				GetVal(u, searchfields, result, datelist, intlist, strlist)
			}
		case len(myargs) == 5:
			u, err := client.Call(listmethod, f.Text(), myargs[0], myargs[1], myargs[2], myargs[3], myargs[4])
			checkprint(err, myargs)
			if u != nil {
				GetVal(u, searchfields, result, datelist, intlist, strlist)
			}
		}
	} else {
		u, err := client.Call(listmethod, f.Text())
		checkprint(err, myargs)
		if u != nil {
			GetVal(u, searchfields, result, datelist, intlist, strlist)
		}
		//fmt.Printf("result %v", u)
	}
	_, err = client.Call("auth.logout", f.Text())
	check(err)
	//fmt.Printf("intlist: %v\n", intlist)

}

func GetVal(v xmlrpc.Value, searchfields []string, result *Result,
	datelist []interface{}, intlist []int, strlist []string) {
	/* fmt.Printf("lets see here %v %v\n", len(v.Values()), len(v.Members())) */
	sort.Strings(searchfields)
	if len(v.Values()) == 0 && len(v.Members()) != 0 {
		x := v.Members()
		for _, v := range x {
			if len(searchfields) != 0 {
				for h := 0; h < len(searchfields); h++ {
					if v.Name() == searchfields[h] {

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

			for _, v := range x {
				if len(searchfields) != 0 {

					for h := 0; h < len(searchfields); h++ {

						/* 	newfieldname := strings.Split(searchfields[h], "_")
						   	fmt.Printf("searchfields newfield: %v\n", newfieldname[len(newfieldname)-1]) */
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
