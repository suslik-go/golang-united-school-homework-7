package coverage

import (
	"os"
	//"errors"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}



//Len returns the length of input value
func TestLen(t *testing.T) {
	var persons People
	
	if persons.Len() != len(persons) {
		t.Errorf("Failed Len function")
	}
}

//Less reports when it is necessary to sort element i before element j
// time year int, month Month, day, hour, min, sec, nsec int, loc *Location
func TestLess(t *testing.T) {
	var persons People
	persons = append(persons, Person{ "Adam" , "Brown" , time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)})
	persons = append(persons, Person{ "Adam" , "Atkinson" , time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)})
	persons = append(persons, Person{ "Billy" , "Brown" , time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)})
	persons = append(persons, Person{ "Billy" , "Brown" , time.Date(1999, time.January, 1, 0, 0, 0, 0, time.UTC)})

	if persons.Less(1,0) == false {
		t.Errorf("Failed Less function")
	}

	if persons.Less(0,2) == false {
		t.Errorf("Failes Less function")
	}

	if persons.Less(3,0) {
		t.Errorf("Failed Less function")
	}

}

//Swap swaps the elements with indexes i and j.
func TestSwap(t *testing.T) {
	var persons People
	persons = append(persons, Person{ "Adam" , "Brown" , time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)})
	persons = append(persons, Person{ "Carol" , "Atkinson" , time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC)})
	
	
	persons.Swap(0,1)

	// if persons[0].firstName != "Carol" || persons[0].lastName != "Atkinson" || persons[0].birthDay.Equal(time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC)) {
	// 	t.Errorf("Failed Swap function")
	// }

}

/*New creates a matrix from a string.*/
func TestNew(t *testing.T) {

	test1 := `78	12 45
	4 45 78 
	78 2 35`

	_, err := New(test1)

	if err != nil {
		t.Errorf("Failed New function")
	}


	
}


/*Rows gets the matrix represented in rows.*/
func TestRows(t *testing.T) {
	matrix := Matrix {
		rows: 3,
		cols: 3,
		data: []int{
			12, 54,78,
			78, 12, 56,
			78, 89, 45,
		},
	}

	if len(matrix.Rows()) != 3 {
		t.Errorf("Failed Rows function")
	}

}


/*Cols gets the matrix represented in columns*/
func TestCols(t *testing.T) {
	matrix := Matrix {
		rows: 3,
		cols: 3,
		data: []int{
			12, 54,78,
			78, 12, 56,
			78, 89, 45,
		},
	}

	if len(matrix.Cols()) != 3 {
		t.Errorf("Failed Rows function")
	}
}


/*Set sets the value of the matrix at point row, col.*/
func TestSet(t *testing.T) {
	var matrix Matrix

	matrix.Set(2,5,10)

	if matrix.data[2 + 5 *2] != 10 {
		t.Errorf("Failed Set function")
	}
}