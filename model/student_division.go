package model

type StudentDivision int

const (
	StudentNotStudent StudentDivision = iota
	StudentElementaryOrLess
	StudentJuniorHighSchool
	StudentHighSchool
	StudentUniversity
)
