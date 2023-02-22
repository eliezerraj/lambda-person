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

	person99 = domain.NewPerson("PERSON-999","PERSON-999","Mr Delete ME","M")
	adress99 = domain.NewAddress("ADDRESS-99","ADDRESS-99","St nini nine",1,"zip-99")
	listAdresses99 = []domain.Address{*adress99}
	personAddress99 = domain.NewPersonAddress(*person99, listAdresses99)
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
		t.Errorf("Error - TestAddPerson Create Repository DynanoDB")
	}

	personService	:= NewPersonService(*personRepository)
	
    result, err := personService.AddPerson(*person)
	if err != nil {
		t.Errorf("Error -TestAddPerson Access DynanoDB %v ", tableName)
	}
	//println(cmp.Equal(person, result))

	if (cmp.Equal(person, result)) {
		t.Logf("Success on TestAddPerson!!! result : %v ", result)
	} else {
		t.Errorf("Error TestAddPerson input : %v" , *person)
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
		t.Errorf("Error - TestListPerson Access DynanoDB %v ", tableName)
	}

	t.Logf("Success TestListPerson !!! result : %v", result)
}

func TestGetPerson(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		t.Errorf("Error -TestGetPerson  Create Repository DynanoDB")
	}

	personService	:= NewPersonService(*personRepository)
	
    result, err := personService.GetPerson(person.ID)
	if err != nil {
		t.Errorf("Error - TestGetPerson Access DynanoDB %v :", tableName)
	}

	if (cmp.Equal(person, result)) {
		t.Logf("Success -TestGetPerson !!! result : %v :", result)
	} else {
		t.Errorf("Error TestGetPerson input : %v : ", person.ID)
	}

	result, err = personService.GetPerson(person2.ID)
	if err != nil {
		t.Errorf("Error Access DynanoDB %v ", tableName)
	}

	if (cmp.Equal(person2, result)) {
		t.Logf("Success -TestGetPerson !!! result : %v :", result)
	} else {
		t.Errorf("Error TestGetPerson input : %v : ", person2.ID)
	}

}

func TestAddPersonDelete(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestAddPersonDelete Create Repository DynanoDB")
	}

	personService	:= NewPersonService(*personRepository)
	
    result, err := personService.AddPerson(*person99)
	if err != nil {
		t.Errorf("Error -TestAddPersonDelete Access DynanoDB %v ", tableName)
	}

	if (cmp.Equal(person99, result)) {
		t.Logf("Success on TestAddPersonDelete!!! result : %v ", result)
	} else {
		t.Errorf("Error TestAddPersonDelete input : %v" , person99)
	}
}

func TestDeletePerson(t *testing.T) {

	t.Setenv("AWS_REGION", "us-east-2")
	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestDeletePerson Create Repository DynanoDB")
	}

	personService	:= NewPersonService(*personRepository)
	
    err = personService.DeletePerson(person99.ID, person99.SK)
	if err != nil {
		t.Errorf("Error - TestDeletePerson DynanoDB %v ", tableName)
	}

	t.Logf("Success TestDeletePerson input : %s %s", person99.ID, person99.SK)
}

func TestAddPersonAddress(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		t.Errorf("Error Create Repository DynanoDB")
	}

	personService	:= NewPersonService(*personRepository)
	personAddress1 := domain.PersonAddress{*person, listAdresses}
	
    result, err := personService.AddPersonAddress(personAddress1)
	if err != nil {
		t.Errorf("Error - TestAddPersonAddress Access DynanoDB %v ", tableName)
	}

	if (cmp.Equal(&personAddress1, result)) {
		t.Logf("Success TestAddPersonAddress result : %v ", result)
	} else {
		t.Errorf("Error NOT EQUAL TestAddPersonAddress input : %v ",&personAddress1)
	}

	personAddress2 := domain.PersonAddress{*person2, listAdresses2}

	result, err = personService.AddPersonAddress(personAddress2)
	if err != nil {
		t.Errorf("Error - TestAddPersonAddress Access DynanoDB %v ", tableName)
	}

	if (cmp.Equal(&personAddress2, result)) {
		t.Logf("Success TestAddPersonAddress result : %v ", result)
	} else {
		t.Errorf("Error NOT EQUAL TestAddPersonAddress input : %v ",&personAddress2)
	}
}

/*func TestListPersonAddress(t *testing.T) {
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
}*/

func TestQueryPersonAddress(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestQueryPersonAddressCreate Repository DynanoDB")
	}

	personService	:= NewPersonService(*personRepository)
	id := person.ID
    result, err := personService.QueryPersonAddress(id)
	if err != nil {
		t.Errorf("Error - TestQueryPersonAddress Access DynanoDB %v ", tableName)
	}
	
	if (result != nil) {
		t.Logf("Success TestQueryPersonAddress result : %v ", result)
	} else {
		t.Errorf("Error TestQueryPersonAddress input %s ", id)
	}

	id = person2.ID
    result, err = personService.QueryPersonAddress(id)
	if err != nil {
		t.Errorf("Error - TestQueryPersonAddress Access DynanoDB %v ", tableName)
	}
	if (result != nil) {
		t.Logf("Success TestQueryPersonAddress result : %v ", result)
	} else {
		t.Errorf("Error TestQueryPersonAddress input %s ", id)
	}
}

func TestQueryPersonAddressNotFound(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestQueryPersonAddressNotFound Repository DynanoDB")
	}

	personService	:= NewPersonService(*personRepository)
	id := "PERSON-9999"
    result, err := personService.QueryPersonAddress(id)
	if err != nil {
		t.Errorf("Error - TestQueryPersonAddressNotFound Access DynanoDB %v ", tableName)
	}

	if (result != nil) {
		t.Logf("Success TestQueryPersonAddressNotFound result : %v ", result)
	} else {
		t.Errorf("Error TestQueryPersonAddressNotFound input %s ", id)
	}
}

func TestAddPersonAddressNotFound(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestAddPersonAddressNotFound Create Repository DynanoDB")
	}

	personService	:= NewPersonService(*personRepository)
	personAddress := domain.PersonAddress{*person99, listAdresses99}
	
    _, err = personService.AddPersonAddress(personAddress)
	if err != nil {
		t.Logf("Success Item not Found result : %v ", err)
	} else {
		t.Errorf("Error - TestAddPersonAddressNotFound Item Found %v ", err)
	}

}