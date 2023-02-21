package services

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"lambda-person/internal/repository"
	"lambda-person/internal/core/domain"

)

var (
	tableName = "person_tenant"
	personRepository	*repository.PersonRepository
	person = domain.NewPerson("PERSON-001","PERSON-001","Mr Luigi","F")
	adress01 = domain.NewAddress("ADDRESS-101","ADDRESS-101","St One",1,"zip-101")
	adress02 = domain.NewAddress("ADDRESS-102","ADDRESS-102","St Two",2,"zip-102")
	listAdresses = []domain.Address{*adress01, *adress02}
	personAddress = domain.NewPersonAddress(*person, listAdresses)

	person2 = domain.NewPerson("PERSON-002","PERSON-002","Mr Cookie","M")
	adress201 = domain.NewAddress("ADDRESS-201","ADDRESS-201","St Three",1,"zip-201")
	adress202 = domain.NewAddress("ADDRESS-202","ADDRESS-202","St Four",2,"zip-202")
	adress203= domain.NewAddress("ADDRESS-203","ADDRESS-203","St Five",2,"zip-203")
	listAdresses2 = []domain.Address{*adress201, *adress202, *adress203}
	personAddress2 = domain.NewPersonAddress(*person2, listAdresses2)

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
		t.Logf("Success 01!!!")
	} else {
		t.Errorf("Error GetPerson 01")
	}

	result, err = personService.GetPerson(person2.ID)
	if err != nil {
		t.Errorf("Error Access DynanoDB %v ", tableName)
	}

	if (cmp.Equal(person2, result)) {
		t.Logf("Success 02!!!")
	} else {
		t.Errorf("Error GetPerson 02")
	}

}
/*
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
}*/

func TestAddPersonAddress(t *testing.T) {

	t.Setenv("AWS_REGION", "us-east-2")
	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		t.Errorf("Error Create Repository DynanoDB")
	}

	personService	:= NewPersonService(*personRepository)
	
    result, err := personService.AddPersonAddress(*person, listAdresses)
	if err != nil {
		t.Errorf("Error Access DynanoDB %v ", tableName)
	}
	//println(cmp.Equal(person, result))

	if (cmp.Equal(personAddress, result)) {
		t.Logf("Success 01 !!!")
	} else {
		t.Errorf("Error 01 TestAddPersonAddress")
	}

	result, err = personService.AddPersonAddress(*person2, listAdresses2)
	if err != nil {
		t.Errorf("Error Access DynanoDB %v ", tableName)
	}
	//println(cmp.Equal(person, result))

	if (cmp.Equal(personAddress2, result)) {
		t.Logf("Success 02 !!!")
	} else {
		t.Errorf("Error 02 TestAddPersonAddress")
	}
}

func TestListPersonAddress(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		t.Errorf("Error Create Repository DynanoDB")
	}

	personService	:= NewPersonService(*personRepository)
	
    result, err := personService.ListPersonAddress()
	if err != nil {
		t.Errorf("Error Access DynanoDB %v ", tableName)
	}
	
	t.Logf("Success !!! result : %v", result)
}

func TestQueryPersonAddress(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		t.Errorf("Error Create Repository DynanoDB")
	}

	personService	:= NewPersonService(*personRepository)
	
    result, err := personService.QueryPersonAddress(*person)
	if err != nil {
		t.Errorf("Error Access DynanoDB %v ", tableName)
	}
	
	t.Logf("Success !!! result : %v", result)
}