//package main

import (
    "fmt"
    "log"
)

func main() {

//uncomment to test without using Payload option
    // ARGS["first_name"] = "Bryan"
    // ARGS["last_name"] = "Smith"

    value, ok := ARGS["first_name"]
    

    if !ok { 
        log.Println("first_name is required")
        return 
    }
	
	firstName, ok := value.(string)

    if !ok { 
        log.Println("first_name must be a string")
        return 
    }

    value, ok = ARGS["last_name"]
	
    if !ok { 
        log.Println("last_name is required")
        return 
    }
    
	lastName , ok:= value.(string)
    
    if !ok { 
        log.Println("last_name must be a string")
        return 
    }

	var preferred = ""

	if len(firstName) > 0 && len(lastName) > 0 {

		if (len(firstName)+len(lastName))%2 == 0 {
			preferred = "cats"
		} else {
			preferred = "dogs"
		}
	}
    
    result := map[string]string{ "preferred": preferred }
    resultBytes, err := json.Marshal(result)
    
    if err != nil {
        log.Println("failed to serialize result to json")
        return
    }
    
    fmt.Println(string(resultBytes))
}pe
