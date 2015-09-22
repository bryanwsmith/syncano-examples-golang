//package main

import (
    "fmt"
    "log"
)

func main() {

//uncomment to test without using Payload option
    data := map[string]interface{}{ "first_name" : "Bryan", "last_name": "Smith" }
    payload := map[string]interface{}{"payload" : data }
    ARGS["POST"] = payload


    var params map[string]interface{}

    ok := true

    params, ok = ARGS["GET"].(map[string]interface{})

    if !ok || len(params) == 0 {

        params, ok = ARGS["POST"].(map[string]interface{})

        //log.Println(params)

        if ok && len(params) > 0 {
            params, ok = params["payload"].(map[string]interface{})
        }

        if !ok || len(params) == 0 { 
            var msg = "parameters first_name, last_name must be posted as x-www-form-urlencoded payload=json, example:  payload='{\"first_name\": \"Bryan\", \"last_name\": \"Smith\"}'"
            msg += " or sent in the query string /?first_name=value&last_name=value"
         
            log.Println(msg)
         
            return 
        }
    }

    value, ok := params["first_name"]

    if !ok { 
        log.Println("first_name is required")
        return 
    }
	
	firstName, ok := value.(string)

    if !ok { 
        log.Println("first_name must be a string")
        return 
    }

    value, ok = params["last_name"]
	
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

		if (len(firstName)+len(lastName)) %2 == 0 {
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
}
