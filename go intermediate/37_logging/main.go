package main

import (
	"log"
	"os"
)

var( 
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)


func main() {
	log.Println("This is a log message");
	
	log.SetPrefix("INFO: ")
	log.Println("This is a info message");

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("This is a log message with date, time and short file name");

	infoLogger.Println("This is an info message using custom logger");
	warningLogger.Println("This is a warning message using custom logger");
	errorLogger.Println("This is an error message using custom logger");

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	debugLogger := log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger.Println("This is a debug message written to a file");

	infoLogger1 := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger1 := log.New(file, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger1 := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	infoLogger1.Println("This is an info message using custom logger");
	warningLogger1.Println("This is a warning message using custom logger");
	errorLogger1.Println("This is an error message using custom logger");
}