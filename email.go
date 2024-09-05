package main

import (
	"errors"
	"fmt"
	"math/rand"
)




//This function mocks the behaviour of the email service
//I have made the logic such that it will fail more rather succeed
func emailMockService(name,email string)error{
	randomNumber:=rand.Intn(5)
	if randomNumber<4{
		emailMessage:=getEmailMessage(name,email)
	fmt.Print(emailMessage)
	return nil
	}
	
	return errors.New("email service failed try again")
	
}

func getEmailMessage(name,email string)string{
	companyEmail:="digivatelabs.com@gmail.com"
	return fmt.Sprintf(`
*----------------------------------------------------------------------------*
 From : %s																	 
*----------------------------------------------------------------------------*
 To : %s 																	 
*----------------------------------------------------------------------------*
 Hey %s,																	 
 																			 
 You are registered successfully Enjoy our blazingly fast services			 
 																			 
 Regards,																	
 Digivate Labs																

*----------------------------------------------------------------------------*`,companyEmail,email,name)
}
