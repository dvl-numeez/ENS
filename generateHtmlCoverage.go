package main

import (
	"log"
	"os/exec"
)


func GenerateHtmlCoverageReport(){
	coverage:=exec.Command("go","test","./...","-coverprofile=cover.out")
	coverageHtml:=exec.Command("go","tool","cover","-html=cover.out")

	if err:=coverage.Run();err!=nil{
		log.Fatal(err)
	}
	if err:=coverageHtml.Run();err!=nil{
		log.Fatal(err)
	}

}