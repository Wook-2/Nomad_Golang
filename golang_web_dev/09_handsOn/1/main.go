package main

import (
	"log"
	"os"
	"text/template"
)

// Create a data structure to pass to a template which
// 1. contains information about California hotels including Name, Address, City, Zip, Region
// 2. Region can be: Southern, Central, Northern
// 3. can hold unlimited number of hotels

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
}

type region struct {
	Region string
	Hotels []hotel
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	h1 := hotel{
		Name:    "hotelA",
		Address: "123-456",
		City:    "CityA",
		Zip:     "ABC-123",
	}
	h2 := hotel{
		Name:    "hotelB",
		Address: "987-654",
		City:    "CityB",
		Zip:     "DEF-456",
	}

	hotels := []hotel{
		h1,
		h2,
	}

	r1 := region{
		Region: "Region A",
		Hotels: hotels,
	}
	err := tpl.Execute(os.Stdout, r1)

	if err != nil {
		log.Fatalln(err)
	}
}
