package filters

import "fmt"

type printfilter interface {
	printfiltered(s []string)
}

func (r *r) printfiltered(s []string) {
	var length int
	if len(r.Stringmap) != 0 {

		if length < len(r.Stringmap) {
			length = len(r.Stringmap)
		}

	}

	if len(r.Intmap) != 0 {

		if length < len(r.Intmap) {
			length = len(r.Intmap)
		}

	}

	if len(r.Datemap) != 0 {

		if length < len(r.Datemap) {
			length = len(r.Datemap)
		}

	}

	if len(r.Innerlistmap) != 0 {
		if length < len(r.Innerlistmap) {
			length = len(r.Innerlistmap)
		}

	}

	for index := 0; index < length; index++ {
		fmt.Printf("\nnew data set: ----------------\n")
		for i := 0; i < len(s); i++ {
			r.printstring(index, s[i])
			r.printinteger(index, s[i])
			r.printdatetime(index, s[i])
			r.printinnerlist(index, s[i])
		}
	}

}

func (r *r) printstring(i int, s string) {
	if len(r.Stringmap) != 0 {
		for c, d := range r.Stringmap[i] {
			if c == s {
				fmt.Printf("%v:", c)
				fmt.Printf("\t%v\n", d)
			}
		}

	}
}

func (r *r) printinteger(i int, s string) {
	if len(r.Intmap) != 0 {
		for a, b := range r.Intmap[i] {
			if a == s {
				fmt.Printf("%v:", a)
				fmt.Printf("\t%v\n", b)
			}

		}
	}
}

func (r *r) printdatetime(i int, s string) {
	if len(r.Datemap) != 0 {
		for a, b := range r.Datemap[i] {
			if a == s {
				fmt.Printf("%v:", a)
				fmt.Printf("\t%v\n", b)
			}

		}
	}
}

func (r *r) printinnerlist(i int, s string) {
	if len(r.Innerlistmap) != 0 {
		for _, d := range r.Innerlistmap[i] {
			for a, b := range d {
				if a == s {
					fmt.Println("inner list of api return:")
					fmt.Printf("\t%v:", a)
					fmt.Printf("\t%v\n", b)
				}

			}
		}

	}
}
