package main

import (
	"fmt"
	"log"

	"github.com/bjin01/autoapi/dependmethod"
	"github.com/bjin01/autoapi/finalmethod"
	"github.com/bjin01/autoapi/getyaml"
	"github.com/bjin01/autoapi/method1"
	"github.com/bjin01/autoapi/method2"
	"github.com/bjin01/autoapi/printfinalresult"
	"github.com/bjin01/autoapi/printresult"
	"github.com/bjin01/autoapi/printresult2"
	"github.com/bjin01/autoapi/sort"
)

const (
	SUMAURL string = "http://bjsuma.bo2go.home/rpc/api"
)

type Login struct {
	Username string
	Passwd   string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
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

	result := new(method1.Result)
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

		method1.Method1(cfg.Server.ApiUrl, cfg.Server.Username, cfg.Server.Password,
			cfg.Method1.Methodname, cfg.Method1.InputVars, Sortedmethod1Outvars, result)

	}

	//fmt.Printf("new sorted slice: %q\n", Sortedmethod1Outvars)
	resultsmethod1 := printresult.Printresult(result, sort.SortSlice(cfg.Method1.Outvariables))

	if cfg.Method2.Methodname != "" {
		fmt.Printf("cfg.Method2.InputVars: %v\n", cfg.Method2.InputVars)
		//fmt.Printf("%v\n", cfg.Method2.InputVars)
		method2.Method2(cfg.Server.ApiUrl, cfg.Server.Username, cfg.Server.Password,
			cfg.Method2.Methodname, cfg.Method2.InputVars, Sortedmethod2Outvars, result2, resultsmethod1)
	}

	resultsmethod2 := printresult2.Printresult(result2, sort.SortSlice(cfg.Method2.Outvariables))

	if cfg.Finalmethod.Methodname != "" {
		//fmt.Printf("hallo %v\n", cfg.Finalmethod.InputVars)
		if cfg.Finalmethod.Options.Meth2dependmeth1 == true {
			dependmethod.Dependmethod(cfg.Finalmethod.InputVars, cfg.Method2.InputVars, resultsmethod1, cfg,
				Sortedmethod2Outvars, SortedFinalmethodOutvars)
		} else {
			finalmethod.Finalmethod(cfg.Server.ApiUrl, cfg.Server.Username, cfg.Server.Password,
				cfg.Finalmethod.Methodname, cfg.Finalmethod.InputVars, SortedFinalmethodOutvars,
				resultfinal, resultsmethod1, resultsmethod2)
		}

	}

	printfinalresult.Printresult(resultfinal, sort.SortSlice(cfg.Finalmethod.Outvariables))
	//fmt.Printf("Final Job Output: %v\n", resultfinalmethod)
}
