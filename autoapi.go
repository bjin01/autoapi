package main

import (
	"fmt"
	"log"

	"github.com/bjin01/autoapi/callapi"
	"github.com/bjin01/autoapi/getyaml"
)

const (
	SUMAURL string = "http://bjsuma.bo2go.home/rpc/api"
)

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

func main() {
	cfgPath, err := getyaml.ParseFlags()
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := getyaml.NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	c := callapi.C{
		Cfg: cfg,
	}
	var mycall callapi.Caller = &c
	var print_u callapi.Printvalue
	u, err := mycall.Callapi("method1", nil, nil)
	check(err)
	if u != nil {
		result := callapi.R{
			U: u,
		}
		print_u = &result

		err := print_u.Printmethod1(&c)
		check(err)
		/* log.Printf("method1 output: \n")
		fmt.Println(u.Values()) */
	}

	u2, err := mycall.Callapi("method2", u, nil)
	fmt.Println("finalmethod: ", u2)
	check(err)
	if u2 != nil {
		result := callapi.R{
			U: u2,
		}
		print_u = &result

		err := print_u.Printmethod2(&c)
		check(err)
	}

	u3, err := mycall.Callapi("finalmethod", u, u2)
	check(err)
	if u3 != nil {
		result := callapi.R{
			U: u3,
		}
		print_u = &result

		err := print_u.Printfinalmethod(&c)
		check(err)
	}

	/* result := new(method1.Result)
	result2 := new(method2.Result)
	resultfinal := new(finalmethod.Result)
	//resultdependmethod := new(dependmethod.Result)
	var Sortedmethod1Outvars, Sortedmethod2Outvars, SortedFinalmethodOutvars []string
	if len(cfg.Method1.Outvariables) != 0 {
		Sortedmethod1Outvars = sort.SortSlice(cfg.Method1.Outvariables)
	}

	if len(cfg.Method2.Outvariables) != 0 {
		Sortedmethod2Outvars = sort.SortSlice(cfg.Method2.Outvariables)
	}

	if len(cfg.Finalmethod.Outvariables) != 0 {
		SortedFinalmethodOutvars = sort.SortSlice(cfg.Finalmethod.Outvariables)
	}
	//fmt.Printf("%v\n", cfg.Method1.InputVars)
	if cfg.Method1.Methodname != "" {

		method1.Method1(cfg, cfg.Server.ApiUrl, cfg.Server.Username, cfg.Server.Password, cfg.Method1.Methodname, cfg.Method1.InputVars, Sortedmethod1Outvars, result)

	}

	//fmt.Printf("new sorted slice: %q\n", Sortedmethod1Outvars)
	resultsmethod1 := printresult.Printresult(result, sort.SortSlice(cfg.Method1.Outvariables))

	if cfg.Method2.Methodname != "" {
		fmt.Printf("cfg.Method2.InputVars: %v\n", cfg.Method2.InputVars)
		//fmt.Printf("%v\n", cfg.Method2.InputVars)
		method2.Method2(cfg, cfg.Server.ApiUrl, cfg.Server.Username, cfg.Server.Password,
			cfg.Method2.Methodname, cfg.Method2.InputVars, Sortedmethod2Outvars, result2, resultsmethod1)
	}

	resultsmethod2 := printresult2.Printresult(result2, sort.SortSlice(cfg.Method2.Outvariables))

	if cfg.Finalmethod.Methodname != "" {
		//fmt.Printf("hallo %v\n", cfg.Finalmethod.InputVars)
		if cfg.Finalmethod.Options.Meth2dependmeth1 == true {
			dependmethod.Dependmethod(cfg.Finalmethod.InputVars, cfg.Method2.InputVars, resultsmethod1, cfg,
				Sortedmethod2Outvars, SortedFinalmethodOutvars)
		} else {
			finalmethod.Finalmethod(cfg, cfg.Server.ApiUrl, cfg.Server.Username, cfg.Server.Password,
				cfg.Finalmethod.Methodname, cfg.Finalmethod.InputVars, SortedFinalmethodOutvars,
				resultfinal, resultsmethod1, resultsmethod2)
		}

	}

	printfinalresult.Printresult(resultfinal, sort.SortSlice(cfg.Finalmethod.Outvariables)) */
	//fmt.Printf("Final Job Output: %v\n", resultfinalmethod)
}
