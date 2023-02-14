package services

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"lambda-person/internal/repository"
	"lambda-person/internal/core/domain"

)

var (
	tableName = "person"
	personRepository	*repository.PersonRepository
	person = domain.NewPerson("9999","Person Nine","M")
)

func TestSum(t *testing.T) {
    total := Sum(5, 5)
    if total != 10 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}

func TestAddPerson(t *testing.T) {

	t.Setenv("AWS_REGION", "us-east-2")
	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		t.Errorf("Error Create Repository DynanoDB")
	}

	personService	:= NewPersonService(*personRepository)
	
    result, err := personService.AddPerson(*person)
	if err != nil {
		t.Errorf("Error Access DynanoDB %v ", tableName)
	}
	//println(cmp.Equal(person, result))

	if (cmp.Equal(person, result)) {
		t.Logf("Success !!!")
	} else {
		t.Errorf("Error AddPerson")
	}
}

func TestListPerson(t *testing.T) {

	t.Setenv("AWS_REGION", "us-east-2")
	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		t.Errorf("Error Create Repository DynanoDB")
	}

	personService := NewPersonService(*personRepository)
    result, err := personService.ListPerson()
	if err != nil {
		t.Errorf("Error Access DynanoDB %v ", tableName)
	}

	t.Logf("Success !!! result : %v", result)

}

func TestGetPerson(t *testing.T) {

	t.Setenv("AWS_REGION", "us-east-2")
	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		t.Errorf("Error Create Repository DynanoDB")
	}

	personService	:= NewPersonService(*personRepository)
	
    result, err := personService.GetPerson(person.ID)
	if err != nil {
		t.Errorf("Error Access DynanoDB %v ", tableName)
	}

	if (cmp.Equal(person, result)) {
		t.Logf("Success !!!")
	} else {
		t.Errorf("Error AddPerson")
	}
}

func TestDeletePerson(t *testing.T) {

	t.Setenv("AWS_REGION", "us-east-2")
	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		t.Errorf("Error Create Repository DynanoDB")
	}

	personService	:= NewPersonService(*personRepository)
	
    err = personService.DeletePerson(person.ID)
	if err != nil {
		t.Errorf("Error Access DynanoDB %v ", tableName)
	}

	t.Logf("Success !!!")
}