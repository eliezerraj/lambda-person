package services

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/google/go-cmp/cmp"

	"lambda-person/internal/repository"
	"lambda-person/internal/core/domain"
	"lambda-person/internal/adapter/notification"

)

var (
	tableName 		= "person_tenant"
	eventSource		=	"lambda-person"
	eventBusName	=	"event-bus-person"
	personRepository	*repository.PersonRepository
	personNotification *notification.PersonNotification

	person = domain.NewPerson("001","","Mr Luigi","F")
	adress01 = domain.NewAddress("ADDRESS-101","ADDRESS-101","St One",1,"zip-101")
	adress02 = domain.NewAddress("ADDRESS-102","ADDRESS-102","St Two",2,"zip-102")
	listAdresses = []domain.Address{*adress01, *adress02}
	//personAddress = domain.NewPersonAddress(*person, listAdresses)

	person2 = domain.NewPerson("002","","Mr Cookie","M")
	adress201 = domain.NewAddress("ADDRESS-201","ADDRESS-201","St Three",1,"zip-201")
	adress202 = domain.NewAddress("ADDRESS-202","ADDRESS-202","St Four",2,"zip-202")
	adress203= domain.NewAddress("ADDRESS-203","ADDRESS-203","St Five",2,"zip-203")
	listAdresses2 = []domain.Address{*adress201, *adress202, *adress203}
	//personAddress2 = domain.NewPersonAddress(*person2, listAdresses2)

	person99 = domain.NewPerson("999","","Mr Delete ME","M")
	adress99 = domain.NewAddress("ADDRESS-99","ADDRESS-99","St nini nine",1,"zip-99")
	listAdresses99 = []domain.Address{*adress99}
	personAddress99 = domain.NewPersonAddress(*person99, listAdresses99)
)

/*func TestSum(t *testing.T) {
    total := Sum(5, 5)
    if total != 10 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}*/

func TestAddPerson(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestAddPerson Create Repository DynanoDB")
	}

	personNotification, err = notification.NewPersonNotification(eventSource,eventBusName)
	if err != nil {
		t.Errorf("Error -TestAddPerson Access EventBridge %v : ", err)
	}

	personService	:= NewPersonService(*personRepository, *personNotification)
	
	// ============================================ //
	personAdd := domain.NewPerson("001","","Mr Luigi","F")

    result, err := personService.AddPerson(*personAdd)
	if err != nil {
		t.Errorf("Error -TestAddPerson DynanoDB %v err : %v", tableName, err)
	}
	//println(cmp.Equal(person, result))

	// setting keys
	personAdd.ID = "PERSON-" + personAdd.ID
	personAdd.SK = personAdd.ID
	if (cmp.Equal(personAdd, result)) {
		t.Logf("Success on TestAddPerson!!! result : %v ", result)
	} else {
		t.Errorf("Error TestAddPerson input : %v" , *personAdd)
	}

	// ============================================ //
	personAdd2 := domain.NewPerson("002","","Mr Cookie","M")
    result, err = personService.AddPerson(*personAdd2)
	if err != nil {
		t.Errorf("Error -TestAddPerson Access DynanoDB %v err: %v ", tableName, err)
	}
	//println(cmp.Equal(person, result))

	// setting keys
	personAdd2.ID = "PERSON-" + personAdd2.ID
	personAdd2.SK = personAdd2.ID
	if (cmp.Equal(personAdd2, result)) {
		t.Logf("Success on TestAddPerson!!! result : %v ", result)
	} else {
		t.Errorf("Error TestAddPerson input : %v" , *personAdd2)
	}
}

func TestUpdatePerson(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestUpdatePerson Create Repository DynanoDB")
	}

	personNotification, err = notification.NewPersonNotification(eventSource,eventBusName)
	if err != nil {
		t.Errorf("Error -TestUpdatePerson Access EventBridge %v ", err)
	}

	personService	:= NewPersonService(*personRepository, *personNotification)
	
	personUpdate := domain.NewPerson("099","","Mr Luigi","F")
    result, err := personService.AddPerson(*personUpdate)
	if err != nil {
		t.Errorf("Error -TestAddPerson DynanoDB %v err : %v", tableName, err)
	}

	personUpdate = domain.NewPerson("099","","Mr Luigi - UPDATED","M")
    result, err = personService.UpdatePerson(*personUpdate)
	if err != nil {
		t.Errorf("Error -TestUpdatePerson Access DynanoDB %v err: %v", tableName, err)
	}

	personUpdate.ID = "PERSON-" + personUpdate.ID
	personUpdate.SK = personUpdate.ID
	if (cmp.Equal(personUpdate, result)) {
		t.Logf("Success on TestUpdatePerson!!! result : %v ", result)
	} else {
		t.Errorf("Error TestUpdatePerson input : %v" , personUpdate)
	}
}

func TestListPerson(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		t.Errorf("Error Create Repository DynanoDB")
	}

	personNotification, err = notification.NewPersonNotification(eventSource,eventBusName)
	if err != nil {
		t.Errorf("Error -TestListPerson Access EventBridge : %v ", err)
	}

	personService	:= NewPersonService(*personRepository, *personNotification)

    result, err := personService.ListPerson()
	if err != nil {
		t.Errorf("Error - TestListPerson Access DynanoDB %v err: %v ", tableName, err)
	}

	t.Logf("Success TestListPerson !!! result : %v", result)
}

func TestGetPerson(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		t.Errorf("Error -TestGetPerson  Create Repository DynanoDB")
	}

	personNotification, err = notification.NewPersonNotification(eventSource,eventBusName)
	if err != nil {
		t.Errorf("Error -TestGetPerson Access EventBridge %v ", err)
	}

	personService	:= NewPersonService(*personRepository, *personNotification)

	//================
	personGet := domain.NewPerson("001","","Mr Luigi","F")
    result, err := personService.GetPerson(personGet.ID)
	if err != nil {
		t.Errorf("Error - TestGetPerson Access DynanoDB %v err: %v:", tableName, err)
	}

	// setting keys
	personGet.ID = "PERSON-" + personGet.ID
	personGet.SK = personGet.ID
	if (cmp.Equal(personGet, result)) {
		t.Logf("Success -TestGetPerson !!! result : %v :", result)
	} else {
		t.Errorf("Error TestGetPerson input : %v ", personGet.ID)
	}

	personGet2 := domain.NewPerson("002","","Mr Cookie","M")
	result, err = personService.GetPerson(personGet2.ID)
	if err != nil {
		t.Errorf("Error Access DynanoDB %v err: %v ", tableName, err)
	}

	// setting keys
	personGet2.ID = "PERSON-" + personGet2.ID
	personGet2.SK = personGet2.ID
	if (cmp.Equal(personGet2, result)) {
		t.Logf("Success -TestGetPerson !!! result : %v :", result)
	} else {
		t.Errorf("Error TestGetPerson input : %v ", personGet2.ID)
	}

}

func TestAddPersonDelete(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestAddPersonDelete Create Repository DynanoDB")
	}

	personNotification, err = notification.NewPersonNotification(eventSource,eventBusName)
	if err != nil {
		t.Errorf("Error -TestAddPersonDelete Access EventBridge %v ", err)
	}

	personService	:= NewPersonService(*personRepository, *personNotification)
	
	personDel := domain.NewPerson("999","","Mr Delete ME","M")
    result, err := personService.AddPerson(*personDel)
	if err != nil {
		t.Errorf("Error -TestAddPersonDelete Access DynanoDB %v err: %v", tableName, err)
	}

	personDel.ID = "PERSON-" + personDel.ID
	personDel.SK = personDel.ID
	if (cmp.Equal(personDel, result)) {
		t.Logf("Success on TestAddPersonDelete!!! result : %v ", result)
	} else {
		t.Errorf("Error TestAddPersonDelete input : %v" , personDel)
	}
}

func TestDeletePerson(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestDeletePerson Create Repository DynanoDB")
	}

	personNotification, err = notification.NewPersonNotification(eventSource,eventBusName)
	if err != nil {
		t.Errorf("Error -TestDeletePerson Access EventBridge %v ", err)
	}

	personService	:= NewPersonService(*personRepository, *personNotification)
	
    err = personService.DeletePerson(person99.ID, person99.ID)
	if err != nil {
		t.Errorf("Error - TestDeletePerson DynanoDB %v err: %v", tableName, err)
	}

	t.Logf("Success TestDeletePerson input : %s %s", person99.ID, person99.SK)
}

/*func TestListPersonAddress(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		t.Errorf("Error Create Repository DynanoDB")
	}

	personNotification, err = notification.NewPersonNotification(eventSource,eventBusName)
	if err != nil {
		t.Errorf("Error -TestDeletePerson Access EventBridge %v ", err)
	}

	personService	:= NewPersonService(*personRepository, *personNotification)
	
    result, err := personService.ListPersonAddress()
	if err != nil {
		t.Errorf("Error Access DynanoDB %v err: %v", tableName, err)
	}
	
	t.Logf("Success !!! result : %v", result)
}*/

func TestAddPersonAddress(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	
	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		t.Errorf("Error Create Repository DynanoDB")
	}

	personNotification, err = notification.NewPersonNotification(eventSource,eventBusName)
	if err != nil {
		t.Errorf("Error -TestAddPerson Access EventBridge %v ", err)
	}

	personService	:= NewPersonService(*personRepository, *personNotification)
	
	person_address := domain.NewPerson("001","","Mr Luigi","F")
	personAddress1 := domain.PersonAddress{*person_address, listAdresses}
	
    result, err := personService.AddPersonAddress(personAddress1)
	if err != nil {
		t.Errorf("Error - TestAddPersonAddress Access DynanoDB %v err: %v", tableName, err)
	}

	// setting keys
	person_address.ID = "PERSON-" + person_address.ID
	personAddress1 = domain.PersonAddress{*person_address, listAdresses}
	if (cmp.Equal(&personAddress1, result)) {
		t.Logf("Success TestAddPersonAddress result : %v ", result)
	} else {
		t.Errorf("Error NOT EQUAL TestAddPersonAddress input : %v || result : %v ",&personAddress1, result)
	}

	person_address2 := domain.NewPerson("002","","Mr Cookie","M")
	personAddress2 := domain.PersonAddress{*person_address2, listAdresses2}

	result, err = personService.AddPersonAddress(personAddress2)
	if err != nil {
		t.Errorf("Error - TestAddPersonAddress Access DynanoDB %v err: ¨%v ", tableName, err)
	}

	// setting keys
	person_address2.ID = "PERSON-" + person_address2.ID
	personAddress2 = domain.PersonAddress{*person_address2, listAdresses2}
	if (cmp.Equal(&personAddress2, result)) {
		t.Logf("Success TestAddPersonAddress result : %v ", result)
	} else {
		t.Errorf("Error NOT EQUAL TestAddPersonAddress input : %v || result : %v ",&personAddress2, result)
	}
}

func TestQueryPersonAddress(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestQueryPersonAddressCreate Repository DynanoDB")
	}

	personNotification, err = notification.NewPersonNotification(eventSource,eventBusName)
	if err != nil {
		t.Errorf("Error -TestAddPerson Access EventBridge %v ", err)
	}

	personService	:= NewPersonService(*personRepository, *personNotification)

	id := person.ID
    result, err := personService.QueryPersonAddress(id)
	if err != nil {
		t.Errorf("Error - TestQueryPersonAddress Access DynanoDB %v err: %v ", tableName, err)
	}
	
	if (result != nil) {
		t.Logf("Success TestQueryPersonAddress id: %v || result : %v ", id ,result)
	} else {
		t.Errorf("Error TestQueryPersonAddress input %s ", id)
	}

	id = person2.ID
    result, err = personService.QueryPersonAddress(id)
	if err != nil {
		t.Errorf("Error - TestQueryPersonAddress Access DynanoDB %v err:%v ", tableName, err)
	}
	if (result != nil) {
		t.Logf("Success TestQueryPersonAddress id: %v || result : %v ", id ,result)
	} else {
		t.Errorf("Error TestQueryPersonAddress input %s ", id)
	}
}

func TestQueryPersonAddressNotFound(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestQueryPersonAddressNotFound Repository DynanoDB")
	}

	personNotification, err = notification.NewPersonNotification(eventSource,eventBusName)
	if err != nil {
		t.Errorf("Error -TestAddPerson Access EventBridge %v ", err)
	}

	personService	:= NewPersonService(*personRepository, *personNotification)

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
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestAddPersonAddressNotFound Create Repository DynanoDB")
	}

	personNotification, err = notification.NewPersonNotification(eventSource,eventBusName)
	if err != nil {
		t.Errorf("Error -TestAddPerson Access EventBridge %v ", err)
	}

	personService	:= NewPersonService(*personRepository, *personNotification)

	personAddress := domain.PersonAddress{*person99, listAdresses99}
	
    _, err = personService.AddPersonAddress(personAddress)
	if err != nil {
		t.Logf("Success Item not Found result : %v ", err)
	} else {
		t.Errorf("Error - TestAddPersonAddressNotFound Item Found %v ", err)
	}
}