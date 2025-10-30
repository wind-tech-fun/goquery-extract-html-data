package common

import (
	"fmt"
	"encoding/json"
    "os"
)

/**
*
* Generate json file from `[]map[string]string` structure data source
*
*/
func GenJsonFile(data []map[string]string, file_name string) {
    jsonData, err := json.MarshalIndent(data, "", "  ")

    if err != nil {
        fmt.Println("Error marshaling data to JSON:", err)
        return
    }

    err = os.WriteFile(file_name, jsonData, 0644)
    if err != nil {
        fmt.Println("Error writting JSON to file:", err)
        return
    }
    fmt.Println("json file '" + file_name + "' generated successfully")
}
