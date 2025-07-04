package main

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"
)


func main() {
	// // tmpl := template.New("example")
	// // tmpl, err := template.New("example").Parse("Hello, {{.name}}! How are you today?")
	// tmpl := template.Must(template.New("example").Parse("Hello, {{.name}}! How are you today?"))
	// // if err != nil {
	// // 	panic(err)
	// // }
	// data := map[string]any{
	// 	"name": "Alice",
	// }
	// // err = tmpl.ExecuteTemplate(os.Stdout, "example", data)
	// // if err != nil {
	// // 	panic(err)
	// // }
	// err := tmpl.Execute(os.Stdout, data)
	// if err != nil {
	// 	panic(err)
	// }

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	name = strings.TrimSpace(name)
	templates := map[string]string{
		"Wellcome": "Hello, {{.name}}! Welcome to our program.",
		"notification": "Hello, {{.name}}! You have a new notification! {{.notification}}",
		"error": "Hello, {{.name}}! An error occurred: {{.error}}",
	}
	parsedTemplates := make(map[string]*template.Template)
	for name, tmpl := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for {
		fmt.Println("\nMenu:")
		fmt.Println("Name:", name)
		fmt.Println("1. Wellcome")
		fmt.Println("2. Notification")
		fmt.Println("3. Error")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")
		option, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		option = strings.TrimSpace(option)
		switch option {
		case "1":
			data := map[string]any{"name": name}
			err := parsedTemplates["Wellcome"].Execute(os.Stdout, data)
			if err != nil {
				panic(err)
			}
		case "2":
			fmt.Print("Enter notification message: ")
			notification, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			notification = strings.TrimSpace(notification)
			data := map[string]any{"name": name, "notification": notification}
			err = parsedTemplates["notification"].Execute(os.Stdout, data)
			if err != nil {
				panic(err)
			}
		case "3":
			fmt.Print("Enter error message: ")			
			errorMessage, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			errorMessage = strings.TrimSpace(errorMessage)
			data := map[string]any{"name": name, "error": errorMessage}
			err = parsedTemplates["error"].Execute(os.Stdout, data)
			if err != nil {
				panic(err)
			}
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")	
			}	
	}
}