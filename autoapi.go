package main

import (
	"fmt"
	"log"

	"github.com/bjin01/autoapi/dependmethod"
	"github.com/bjin01/autoapi/finalmethod"
	"github.com/bjin01/autoapi/getyaml"
	"github.com/bjin01/autoapi/listmethod1"
	"github.com/bjin01/autoapi/listmethod2"
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

	result := new(listmethod1.Result)
	result2 := new(listmethod2.Result)
	resultfinal := new(finalmethod.Result)
	//resultdependmethod := new(dependmethod.Result)
	var SortedListmethod1Outvars, SortedListmethod2Outvars, SortedFinalmethodOutvars []string
	if len(cfg.Listmethod1.Outvariables) != 0 {
		SortedListmethod1Outvars = sort.SortSlice(cfg.Listmethod1.Outvariables)
	}

	if len(cfg.Listmethod2.Outvariables) != 0 {
		SortedListmethod2Outvars = sort.SortSlice(cfg.Listmethod2.Outvariables)
	}

	if len(cfg.Finalmethod.Outvariables) != 0 {
		SortedFinalmethodOutvars = sort.SortSlice(cfg.Finalmethod.Outvariables)
	}
	//fmt.Printf("%v\n", cfg.Listmethod1.InputVars)
	if cfg.Listmethod1.Methodname != "" {

		listmethod1.Listmethod1(cfg.Server.ApiUrl, cfg.Server.Username, cfg.Server.Password,
			cfg.Listmethod1.Methodname, cfg.Listmethod1.InputVars, SortedListmethod1Outvars, result)

	}

	//fmt.Printf("new sorted slice: %q\n", SortedListmethod1Outvars)
	resultsmethod1 := printresult.Printresult(result, sort.SortSlice(cfg.Listmethod1.Outvariables))

	if cfg.Listmethod2.Methodname != "" {
		fmt.Printf("cfg.Listmethod2.InputVars: %v\n", cfg.Listmethod2.InputVars)
		//fmt.Printf("%v\n", cfg.Listmethod2.InputVars)
		listmethod2.Listmethod2(cfg.Server.ApiUrl, cfg.Server.Username, cfg.Server.Password,
			cfg.Listmethod2.Methodname, cfg.Listmethod2.InputVars, SortedListmethod2Outvars, result2, resultsmethod1)
	}

	resultsmethod2 := printresult2.Printresult(result2, sort.SortSlice(cfg.Listmethod2.Outvariables))

	if cfg.Finalmethod.Methodname != "" {
		//fmt.Printf("hallo %v\n", cfg.Finalmethod.InputVars)
		if cfg.Finalmethod.Options.Meth2dependmeth1 == true {
			dependmethod.Dependmethod(cfg.Finalmethod.InputVars, cfg.Listmethod2.InputVars, resultsmethod1, cfg,
				SortedListmethod2Outvars, SortedFinalmethodOutvars)
		} else {
			finalmethod.Finalmethod(cfg.Server.ApiUrl, cfg.Server.Username, cfg.Server.Password,
				cfg.Finalmethod.Methodname, cfg.Finalmethod.InputVars, SortedFinalmethodOutvars,
				resultfinal, resultsmethod1, resultsmethod2)
		}

	}

	printfinalresult.Printresult(resultfinal, sort.SortSlice(cfg.Finalmethod.Outvariables))
	//fmt.Printf("Final Job Output: %v\n", resultfinalmethod)
}
