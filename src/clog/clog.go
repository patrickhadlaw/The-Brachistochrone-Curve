//
//	Custom log
//	file: clog.go
//  package: clog
//	description: custom log which writes log data to [LOG].log file
//
//	by: Patrick Stanislaw Hadlaw
//

package clog

import (
	"io/ioutil"
	"strings"
	"os"
	"bufio"
	"time"
	"fmt"
)

var log []string

func Println(e string){ // No formatting just writes line to log file
	s := e
	log = append(log, s)
	f, _ := ioutil.ReadFile("[LOG].log")
	
	err := ioutil.WriteFile("[LOG].log", []byte(string(f[:]) + "\r\n" + s), 0644)
	fmt.Println(e)
	if(err != nil){
		fmt.Println("yklog.Println: failed to log: " + e)
	}
}

func System(e string){ // System error format
	s := "\t[SYSTEM]" + e
	log = append(log, s)
	f, _ := ioutil.ReadFile("[LOG].log")
	
	err := ioutil.WriteFile("[LOG].log", []byte(string(f[:]) + "\r\n" + s), 0644)
	fmt.Println("[SYSTEM]" + e)
	if(err != nil){
		fmt.Println("yklog.Println: failed to log: " + e)
	}
}

func Error(e string){ // General error format
	s := "/t[ERROR]" + e
	log = append(log, s)
	f, _ := ioutil.ReadFile("[LOG].log")
	
	err := ioutil.WriteFile("[LOG].log", []byte(string(f[:]) + "\r\n" + s), 0644)
	fmt.Println("[ERROR]" + e)
	if(err != nil){
		fmt.Println("yklog.Error: failed to log: " + e)
	}
}

func Fatal(e string){ // Fatal error format: writes exception log in log/
	log = append(log, string("\t[FATAL]" + e))
	s := ""
	for _, i := range log{
		s = s + i + "\r\n"
	}
	t := time.Now()
	s = "_____.___. ____   ____  __.\r\n\\__  |   |/  _ \\ |    |/ _|\r\n /   |   |>  _ </\\      <  \r\n \\____   /  <_\\ \\/    |  \\ \r\n / ______\\_____\\ \\____|__ \\\r\n \\/             \\/       \\/\r\n\r\n" + s
	
	err := ioutil.WriteFile("log/[EXCEPTION LOG][" + strings.Replace(t.String(), ":", " ", -1) + "]" + ".log", []byte(s), 0644)
	fmt.Println("[FATAL]" + e + err.Error())
	if(err != nil){
		fmt.Println("yklog.Fatal: failed to write log")
	}
	bufio.NewReader(os.Stdin).ReadString('\n')
	os.Exit(1)
}