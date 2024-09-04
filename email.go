package main

import "fmt"





func emailMockService(name,email string)error{
	emailMessage:=getEmailMessage(name,email)
	fmt.Print(emailMessage)
	return nil

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
