package printresult

import (
	"fmt"

	"github.com/bjin01/autoapi/method1"
)

type PrintResults struct {
	Intmap      map[string][]int
	Stringmap   map[string][]string
	Datetimemap map[string][]interface{}
}

func Printresult(result *method1.Result, s []string) *PrintResults {

	keylist := []string{}

	intmap := make(map[string][]int)
	stringmap := make(map[string][]string)
	datetimemap := make(map[string][]interface{})

	for _, v := range result.Intmap {
		for h, _ := range v {
			//fmt.Printf("got it: %v\n", h)
			keylist = append(keylist, h)

		}
	}
	for _, a := range keylist {
		intlist := []int{}
		for _, v := range result.Intmap {
			for h, i := range v {
				//fmt.Printf("%v: \t%v\n", h, i)
				if h == a {
					intlist = append(intlist, i)
					intmap[h] = intlist
				}
			}
		}
	}

	for _, v := range result.Stringmap {
		for h, _ := range v {
			//fmt.Printf("got it: %v\n", h)
			keylist = append(keylist, h)
		}
	}

	for _, a := range keylist {
		strlist := []string{}
		for _, v := range result.Stringmap {
			for h, i := range v {

				if h == a {
					//fmt.Printf("%v: \t%v\n", h, i)
					strlist = append(strlist, i)
					stringmap[a] = strlist
				}

			}
		}

	}

	for _, v := range result.Datemap {
		for h, _ := range v {
			//fmt.Printf("got it: %v\n", h)
			keylist = append(keylist, h)

		}
	}

	for _, a := range keylist {
		datetimelist := []interface{}{}
		for _, v := range result.Datemap {
			for h, i := range v {
				if h == a {
					//fmt.Printf("%v: \t%v\n", h, i)
					datetimelist = append(datetimelist, i)
					datetimemap[h] = datetimelist
				}
			}
		}
	}

	var Preresults = &PrintResults{
		intmap,
		stringmap,
		datetimemap,
	}

	printsingle(Preresults, s)
	return Preresults

}

func printsingle(result *PrintResults, s []string) {

	var length int
	if len(result.Stringmap) != 0 {
		for _, b := range result.Stringmap {
			if length < len(b) {
				length = len(b)
			}

		}
	}

	if len(result.Intmap) != 0 {
		for _, b := range result.Intmap {
			if length < len(b) {
				length = len(b)
			}

		}
	}

	if len(result.Datetimemap) != 0 {
		for _, b := range result.Datetimemap {
			if length < len(b) {
				length = len(b)
			}

		}
	}

	for index := 0; index < length; index++ {
		fmt.Printf("\nnew data set: ----------------\n")
		for i := 0; i < len(s); i++ {
			printstring(result, index, s[i])
			printinteger(result, index, s[i])
			printdatetime(result, index, s[i])
		}
	}

}

func printstring(result *PrintResults, i int, s string) {
	if len(result.Stringmap) != 0 {
		for a, b := range result.Stringmap {
			if a == s {
				fmt.Printf("%v:", a)
				fmt.Printf("\t%v\n", b[i])
			}
		}
	}
}

func printinteger(result *PrintResults, i int, s string) {
	if len(result.Intmap) != 0 {
		for a, b := range result.Intmap {
			if a == s {
				fmt.Printf("%v:", a)
				fmt.Printf("\t%v\n", b[i])
			}

		}
	}
}

func printdatetime(result *PrintResults, i int, s string) {
	if len(result.Datetimemap) != 0 {
		for a, b := range result.Datetimemap {
			if a == s {
				fmt.Printf("%v:", a)
				fmt.Printf("\t%v\n", b[i])
			}

		}
	}
}
