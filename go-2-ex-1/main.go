package main

import "fmt"

type FullName struct {
	firstName    string 
    lastName     string

	// TODO: add fields
}

// TODO: declare a structure for birth date
type BirthDate struct {
	dayOfBirth byte
	mothOfBirth byte
	yearOfBirth int16
}

type Profile struct {
	// TODO: embed full name and birth date information
	firstName			string 
    lastName			string
	dayOfBirth			byte
	mothOfBirth			byte
	yearOfBirth 		int16
	NumberOfSiblings	byte
	ZodiacSign       	rune
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		firstName: "Max",
		lastName: "Mustermann",
		dayOfBirth: 15,
		mothOfBirth: 8,
		yearOfBirth: 1995,
		NumberOfSiblings: 2,   // TODO: adjust
		ZodiacSign:       '\u2648', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings++;
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
