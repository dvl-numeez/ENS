package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
)

//This function mocks the behaviour of the email service
//I have made the logic such that it will fail more rather succeed
func emailMockService(name,email string)error{
	randomNumber:=rand.Intn(20)
	if randomNumber<7{
		return errors.New("email service failed try again")
	}
	getEmailMessage(os.Stdout,name,email)
	return nil
	
	
	
}

func getEmailMessage(writer io.Writer,name,email string){
	companyEmail:="digivatelabs.com@gmail.com"
	fmt.Fprintf(writer,`
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
