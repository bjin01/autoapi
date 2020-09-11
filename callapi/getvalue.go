package callapi

import (
	"fmt"
	"log"

	"github.com/bjin01/go-xmlrpc"
)

type Printvalue interface {
	Printmethod1(cfg *C) error
	Printmethod2(cfg *C) error
	Printfinalmethod(cfg *C) error
}

type R struct {
	U xmlrpc.Value
}

type Result struct {
	Intmap      map[string]int
	Stringmap   map[string]string
	Datetimemap map[string]interface{}
	Boolmap     map[string]bool
	Innerlist   map[string]interface{}
}

func (r *R) Printfinalmethod(cfg *C) error {

	return nil
}

func (r *R) Printmethod2(cfg *C) error {

	return nil
}

func (r *R) Printmethod1(cfg *C) error {
	log.Printf("method1 output: \n")
	fmt.Printf("%v\n", r.U.Values())
	return nil
}

func (methodvalues *methodinput) getvalue(field string, methodname string, index int) (interface{}, int) {
	var val interface{}

	return val, 0
}

func (methodvalues *methodinput) getarray(field string, methodname string, index int) (interface{}, int) {
	var val interface{}

	return val, 0
}
