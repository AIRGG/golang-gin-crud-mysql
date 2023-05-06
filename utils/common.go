package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func UtilParsingDB(rows *sql.Rows, err error) []map[string]interface{} {
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	var result []map[string]interface{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			log.Fatal(err)
		}

		row := make(map[string]interface{})
		for i, col := range columns {
			val := values[i]
			// vx := type val
			// fmt.Println(vx)
			fmt.Printf(" typenya %T", val)
			switch v := val.(type) {
			case int64:
				row[col] = int(val.(int64))
			case float64:
				row[col] = val.(float64)
			case []byte:
				// If the value is a []byte slice, attempt to parse it as a number or JSON
				// object, depending on its contents.
				if len(v) > 0 && (v[0] == '{' || v[0] == '[') {
					var obj interface{}
					err := json.Unmarshal(v, &obj)
					if err == nil {
						row[col] = obj
						break
					}
				}
				num, err := strconv.Atoi(string(v))
				if err == nil {
					row[col] = num
					break
				}
				row[col] = string(v)
				b := val.([]uint8)
				row[col] = string(b)
			default:
				row[col] = v
			}
		}
		result = append(result, row)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
