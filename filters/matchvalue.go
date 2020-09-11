package filters

import (
	"fmt"

	"github.com/bjin01/autoapi/getyaml"
	"github.com/bjin01/go-xmlrpc"
)

type r struct {
	Intmap       []map[string]int
	Stringmap    []map[string]string
	Datemap      []map[string]interface{}
	Innerlistmap [][]map[string]interface{}
}

type getValer interface {
	filterfields(fieldname string, matchval interface{}, u xmlrpc.Value, searchfields []string)
	getvalue(v xmlrpc.Value, searchfields []string)
	getvalue2(v xmlrpc.Value, fieldname string) map[string]interface{}
	getval3(b xmlrpc.Value) interface{}
}

type sortSearchfielder interface {
	sortOutputFields(methodname string, cfg getyaml.Config)
}

type mySearchFields struct {
	Searchlist []string
}

func (c *mySearchFields) sortOutputFields(methodname string, cfg getyaml.Config) {

	switch methodname {
	case "method1":
		if len(cfg.Method1.Outvariables) != 0 {
			for h := 0; h < len(cfg.Method1.Outvariables); h++ {

				c.Searchlist = append(c.Searchlist, cfg.Method1.Outvariables[h])
			}
		}

	case "method2":
		if len(cfg.Method2.Outvariables) != 0 {
			for h := 0; h < len(cfg.Method2.Outvariables); h++ {

				c.Searchlist = append(c.Searchlist, cfg.Method2.Outvariables[h])
			}
		}

	case "finalmethod":
		if len(cfg.Finalmethod.Outvariables) != 0 {
			for h := 0; h < len(cfg.Finalmethod.Outvariables); h++ {

				c.Searchlist = append(c.Searchlist, cfg.Finalmethod.Outvariables[h])
			}
		}

	}

}

func matchvalues(u xmlrpc.Value, matchvalue interface{}) bool {
	if matchvalue != "" {
		z := u.Kind()
		y := u
		//fmt.Printf("matchvalue is: %v\n", matchvalue)
		if z == 0 {
			return false
		}

		switch f := z; f {

		case 3:
			if matchvalue == y.Bool() {
				return true

			}

		case 6: //this is a int type
			if y.Int() != 0 {
				if matchvalue == y.Int() {
					return true
				}
			}

		case 7: //this is a string type
			//fmt.Printf("the original val is: %v\n", y.Text())
			if y.Text() != "" {
				if matchvalue == y.Text() {
					return true
				}
			}
		}

	} else {
		return false
	}
	return false
}

func (r *r) getvalue(v xmlrpc.Value, searchfields []string) {
	if len(v.Members()) != 0 {
		var innerlist []map[string]interface{}

		for _, y := range v.Members() {

			fmt.Printf("\t%v: \t", y.Name())
			//for all defined filters which is a map
			if len(searchfields) != 0 {
				for h := 0; h < len(searchfields); h++ {
					var innerlistmap map[string]interface{}
					if y.Name() == searchfields[h] {
						innerlistmap = r.getvalue2(y.Value(), y.Name())
					}
					if innerlistmap != nil {
						innerlist = append(innerlist, innerlistmap)
					}
				}

			}

		}
		if len(innerlist) != 0 {
			r.Innerlistmap = append(r.Innerlistmap, innerlist)
		}
	}
}

func (r *r) getvalue2(v xmlrpc.Value, fieldname string) map[string]interface{} {
	z := v.Kind()
	y := v

	if z == 0 {
		return nil
	}

	switch f := z; f {
	case 1:
		//fmt.Printf("This is an array: %v\n", y.Values())
		if len(y.Values()) != 0 {
			var templist []interface{}
			for _, b := range y.Values() {
				val := r.getval3(b)
				templist = append(templist, val)
			}
			innerlistmap := map[string]interface{}{
				fieldname: templist,
			}
			return innerlistmap

		}

	case 2:
		fmt.Printf("\t%v\n", y.Bytes())
	case 3:
		fmt.Printf("\t%v\n", y.Bool())
	case 4:
		fmt.Printf("\t%s\n", y.Time())
		datetimemap := map[string]interface{}{
			fieldname: y.Time(),
		}
		r.Datemap = append(r.Datemap, datetimemap)
		//datelist = append(datelist, y.Time())
	case 5:
		fmt.Printf("%v\n", y.Double())
	case 6: //this is a int type

		intmap := map[string]int{
			fieldname: y.Int(),
		}
		r.Intmap = append(r.Intmap, intmap)
		fmt.Printf("\t%v\n", y.Int())
	case 7: //this is a string type

		stringmap := map[string]string{
			fieldname: y.Text(),
		}
		r.Stringmap = append(r.Stringmap, stringmap)
		//strlist = append(strlist, y.Text())
		fmt.Printf("\t%v\n", y.Text())
	case 8: //this is a member type

		fmt.Printf("This is a Value type 8\n")
	}
	return nil
}

func (r *r) getval3(b xmlrpc.Value) interface{} {
	z := b.Kind()
	y := b

	if z == 0 {
		return 0
	}

	switch f := z; f {
	case 1:
		fmt.Printf("This is an array: %v\n", y.Values())
		if len(y.Values()) != 0 {
			var templist []interface{}
			for _, b := range y.Values() {
				val := r.getval3(b)
				templist = append(templist, val)
			}
			return templist
		}

	case 2:
		fmt.Printf("\t%v\n", y.Bytes())
	case 3:
		fmt.Printf("\t%v\n", y.Bool())
	case 4:

		return y.Time()
	case 5:
		fmt.Printf("%v\n", y.Double())
	case 6: //this is a int type
		fmt.Printf("\t%v\n", y.Int())
		return y.Int()

	case 7: //this is a string type
		fmt.Printf("\t%v\n", y.Text())
		return y.Text()
	case 8: //this is a member type

		fmt.Printf("This is a Value type 8\n")
	}
	return 0
}

func (r *r) filterfields(fieldname string, matchval interface{}, u xmlrpc.Value, searchfields []string) {

	/* fmt.Printf("u.Values(): %v\n", u.Values())
	fmt.Printf("u.Values() length: %v\n", len(u.Values())) */
	found := false
	z := 0
	//if u is not empty
	if len(u.Values()) != 0 {
		for a, b := range u.Values() {
			//fmt.Printf("u.Members(): %v\n", b.Members())
			//if b's value's members is not empty
			if len(b.Members()) != 0 {
				for _, y := range b.Members() {
					//fmt.Printf("a is: %v, \tfieldname is %v\n", a, y.Name())
					//for all defined filters which is a map

					if y.Name() == fieldname {
						if y.Value().Kind() == 0 {
							return
						}
						if y.Value().Kind() == 1 {
							//fmt.Printf("found %v, \t%v\n", y.Name(), len(y.Value().Values()))
							if len(y.Value().Values()) != 0 {
								for _, i := range y.Value().Values() {
									//fmt.Printf("\tin array value kind is: %v\n", i.Kind())
									found = matchvalues(i, matchval)

								}
							}
							//r.filterfields(cfg, y.Value(), searchfields)
						} else {
							//fmt.Printf("\tvalue kind is: %v\n", y.Value().Kind())
							found = matchvalues(y.Value(), matchval)

						}

					}

				}

			}
			if found == true {
				z++
				fmt.Printf("match found, total: %v, in slice number %v\n", z, a)
				r.getvalue(b, searchfields)
			}

		}

	}

}

func (r *r) print() {
	if len(r.Intmap) != 0 {
		for a, b := range r.Intmap {
			fmt.Println(a, b)
		}
	}

	if len(r.Stringmap) != 0 {
		for a, b := range r.Stringmap {
			fmt.Println(a, b)
		}
	}

	if len(r.Datemap) != 0 {
		for a, b := range r.Datemap {
			fmt.Println(a, b)
		}
	}

	if len(r.Innerlistmap) != 0 {
		for a, b := range r.Innerlistmap {
			fmt.Println(a, b)
		}
	}
}

func ApplyFilter(cfg getyaml.Config, u xmlrpc.Value, methodname string) {
	var R r
	var searchlist mySearchFields
	var SortedSearchFields sortSearchfielder = &searchlist
	var result getValer = &R
	var printfiltered printfilter = &R

	if methodname == "method1" {
		if cfg.Method1.Filters != nil && cfg.Method1.Outvariables != nil {
			SortedSearchFields.sortOutputFields("method1", cfg)
			fmt.Printf("\nsorted searchfields: %v\n", searchlist.Searchlist)
			fmt.Printf("filters: ")
			for a, b := range cfg.Method1.Filters {
				fmt.Printf("\t%v: \t%v\n", a, b)
				result.filterfields(a, b, u, searchlist.Searchlist)
			}
			fmt.Printf("output applyfilter:\n")

			printfiltered.printfiltered(searchlist.Searchlist)

		}
	}

	if methodname == "method2" {
		if cfg.Method2.Filters != nil && cfg.Method2.Outvariables != nil {
			SortedSearchFields.sortOutputFields("method2", cfg)
			fmt.Printf("\nsorted searchfields: %v\n", searchlist.Searchlist)
			fmt.Printf("filters: ")
			for a, b := range cfg.Method1.Filters {
				fmt.Printf("\t%v: \t%v\n", a, b)
				result.filterfields(a, b, u, searchlist.Searchlist)
			}
			fmt.Printf("output applyfilter:\n")

			printfiltered.printfiltered(searchlist.Searchlist)

		}
	}

	if methodname == "finalmethod" {
		if cfg.Finalmethod.Filters != nil && cfg.Finalmethod.Outvariables != nil {
			SortedSearchFields.sortOutputFields("finalmethod", cfg)
			fmt.Printf("\nsorted searchfields: %v\n", searchlist.Searchlist)
			fmt.Printf("filters: ")
			for a, b := range cfg.Method1.Filters {
				fmt.Printf("\t%v: \t%v\n", a, b)
				result.filterfields(a, b, u, searchlist.Searchlist)
			}
			fmt.Printf("output applyfilter:\n")

			printfiltered.printfiltered(searchlist.Searchlist)

		}
	}

}
