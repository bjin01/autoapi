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
	}

	if u != nil {
		u2, err := mycall.Callapi("method2", u, nil)

		if u2 != nil {
			fmt.Println("finalmethod: ", u2)
			checkprint(err, nil)
			result := callapi.R{
				U: u2,
			}
			print_u = &result
			err := print_u.Printmethod2(&c)
			check(err)

		}

		if u2 != nil {
			u3, err := mycall.Callapi("finalmethod", u, u2)

			if u3 != nil {
				checkprint(err, nil)
				result := callapi.R{
					U: u3,
				}
				print_u = &result

				err := print_u.Printfinalmethod(&c)
				check(err)

			}
		}
	}
}
