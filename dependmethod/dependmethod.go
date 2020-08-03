package dependmethod

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/bjin01/autoapi/finalmethod"
	"github.com/bjin01/autoapi/getyaml"
	"github.com/bjin01/autoapi/listmethod2"
	"github.com/bjin01/autoapi/printfinalresult"
	"github.com/bjin01/autoapi/printresult"
	"github.com/bjin01/autoapi/printresult2"
	"github.com/bjin01/autoapi/sort"
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

func Dependmethod(inputsfinal map[string]interface{}, inputslist2 map[string]interface{},
	resultsmethod1 *printresult.PrintResults, cfg *getyaml.Config, SortedListmethod2Outvars []string, SortedFinalmethodOutvars []string) {

	if inputsfinal != nil && inputslist2 != nil {

		for k, v := range inputsfinal {
			for x, y := range inputslist2 {
				if strings.Contains(k, x) && reflect.ValueOf(v).String() == reflect.ValueOf(y).String() {
					fmt.Printf("great, we found matching %v: %v\n", k, y)
					_, n := listmethod2.Getfromlistmethod1(splitinputarg(v), resultsmethod1, 0)
					for i := 0; i < n; i++ {
						z, _ := listmethod2.Getfromlistmethod1(splitinputarg(v), resultsmethod1, i)
						cfg.Listmethod2.InputVars[x] = z
						cfg.Finalmethod.InputVars[k] = z
						fmt.Printf("The return from Listmethod2 is: %v, %v\n", cfg.Listmethod2.InputVars, i)
						fmt.Printf("The return from Finalmethod is: %v, %v\n", cfg.Finalmethod.InputVars, i)
						//need to set cfg.Finalmethod.Options.Meth2dependmeth1 to false to avoid forking into
						//dependmethd again.
						cfg.Finalmethod.Options.Meth2dependmeth1 = false
						result2 := new(listmethod2.Result)
						resultfinal := new(finalmethod.Result)
						if cfg.Listmethod2.Methodname != "" {
							fmt.Printf("cfg.Listmethod2.InputVars: %v\n", cfg.Listmethod2.InputVars)
							//fmt.Printf("%v\n", cfg.Listmethod2.InputVars)
							listmethod2.Listmethod2(cfg.Server.ApiUrl, cfg.Server.Username, cfg.Server.Password,
								cfg.Listmethod2.Methodname, cfg.Listmethod2.InputVars, SortedListmethod2Outvars, result2, resultsmethod1)
						}
						resultsmethod2 := printresult2.Printresult(result2, sort.SortSlice(cfg.Listmethod2.Outvariables))
						if cfg.Finalmethod.Methodname != "" {
							finalmethod.Finalmethod(cfg.Server.ApiUrl, cfg.Server.Username, cfg.Server.Password,
								cfg.Finalmethod.Methodname, cfg.Finalmethod.InputVars, SortedFinalmethodOutvars,
								resultfinal, resultsmethod1, resultsmethod2)
						}
						printfinalresult.Printresult(resultfinal, sort.SortSlice(cfg.Finalmethod.Outvariables))

					}

				}
			}
		}
	}

}

func splitinputarg(v interface{}) string {
	s, ok := v.(string)
	//fmt.Printf("\t%v\n", s)

	if ok == true {
		if strings.Contains(s, "listmethod1") {
			x := strings.Split(s, ".")
			return x[len(x)-1]
		}
	}
	return ""
}
