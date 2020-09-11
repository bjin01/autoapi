package callapi

import (
	"fmt"
	"log"

	"github.com/bjin01/autoapi/getyaml"

	"github.com/bjin01/go-xmlrpc"
)

type Caller interface {
	Callapi(method string, result xmlrpc.Value, result2 xmlrpc.Value) (xmlrpc.Value, error)
	runapi(methodname string, myargs []interface{}) (xmlrpc.Value, error)
	getinputargs(methodname string, index int) ([]interface{}, int)
}

type methodinputer interface {
	getvalue(field string, methodname string, index int) (interface{}, int)
	getarray(field string, methodname string, index int) (interface{}, int)
}

type methodinput struct {
	apicallname     string
	outputvariables []string
	result          xmlrpc.Value
	result2         xmlrpc.Value
}

type C struct {
	Cfg *getyaml.Config
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

func (Cfg *C) runapi(methodname string, myargs []interface{}) (xmlrpc.Value, error) {
	client := xmlrpc.NewClient(Cfg.Cfg.Server.ApiUrl)

	f, err := client.Call("auth.login", Cfg.Cfg.Server.Username, Cfg.Cfg.Server.Password)
	check(err)

	var method string

	if methodname == "method1" {
		method = Cfg.Cfg.Method1.Methodname
	} else if methodname == "method2" {
		method = Cfg.Cfg.Method2.Methodname
	} else if methodname == "finalmethod" {
		method = Cfg.Cfg.Finalmethod.Methodname
	} else {
		log.Fatal("error no methodname found")
	}

	switch {
	case len(myargs) == 0:
		fmt.Printf("Calling: %v\n", method)
		u, err := client.Call(method, f.Text())
		return u, err

	case len(myargs) == 1:
		fmt.Printf("Calling: %v\n", method)
		u, err := client.Call(method, f.Text(), myargs[0])
		return u, err

	case len(myargs) == 2:
		fmt.Printf("Calling: %v\n", method)
		u, err := client.Call(method, f.Text(), myargs[0], myargs[1])
		return u, err

	case len(myargs) == 3:
		fmt.Printf("Calling: %v\n", method)
		u, err := client.Call(method, f.Text(), myargs[0], myargs[1], myargs[2])
		return u, err

	case len(myargs) == 4:
		fmt.Printf("Calling: %v\n", method)
		u, err := client.Call(method, f.Text(), myargs[0], myargs[1], myargs[2], myargs[3])
		return u, err

	case len(myargs) == 5:
		fmt.Printf("Calling: %v\n", method)
		u, err := client.Call(method, f.Text(), myargs[0], myargs[1], myargs[2], myargs[3], myargs[4])
		return u, err
	}

	_, err = client.Call("auth.logout", f.Text())
	check(err)
	return nil, nil

}

func (Cfg *C) Callapi(method string, result xmlrpc.Value, result2 xmlrpc.Value) (xmlrpc.Value, error) {

	var current_methodinput methodinput
	if method == "method1" {
		current_methodinput = methodinput{
			apicallname:     Cfg.Cfg.Method1.Methodname,
			outputvariables: Cfg.Cfg.Method1.Outvariables,
		}
	}

	if method == "method2" {
		current_methodinput = methodinput{
			apicallname:     Cfg.Cfg.Method2.Methodname,
			outputvariables: Cfg.Cfg.Method2.Outvariables,
		}
		if result != nil {
			current_methodinput = methodinput{
				apicallname:     Cfg.Cfg.Method2.Methodname,
				outputvariables: Cfg.Cfg.Method2.Outvariables,
				result:          result,
			}
		}

	}

	if method == "finalmethod" {
		current_methodinput = methodinput{
			apicallname:     Cfg.Cfg.Finalmethod.Methodname,
			outputvariables: Cfg.Cfg.Finalmethod.Outvariables,
		}
		if result != nil {
			current_methodinput = methodinput{
				apicallname:     Cfg.Cfg.Method2.Methodname,
				outputvariables: Cfg.Cfg.Method2.Outvariables,
				result:          result,
			}
		}
		if result2 != nil && result != nil {
			current_methodinput = methodinput{
				apicallname:     Cfg.Cfg.Method2.Methodname,
				outputvariables: Cfg.Cfg.Method2.Outvariables,
				result:          result,
				result2:         result2,
			}
		}
	}

	if current_methodinput.apicallname != "" && len(current_methodinput.outputvariables) != 0 {
		if current_methodinput.result != nil {
			fmt.Printf("lets see the current_methodinput.result: %v\n", current_methodinput.result.Values())
		}

		myargs, n := Cfg.getinputargs(method, 0, current_methodinput)
		if n != 0 {
			for i := 0; i < n; i++ {
				myargs, _ := Cfg.getinputargs(method, i, current_methodinput)
				u, err := Cfg.runapi(method, myargs)
				checkprint(err, myargs)

				fmt.Println(u.Values())
			}
		} else {
			u, err := Cfg.runapi(method, myargs)
			checkprint(err, myargs)
			return u, err
		}

	}

	return nil, nil //errors.New("no methodname or output variables found.")
}
