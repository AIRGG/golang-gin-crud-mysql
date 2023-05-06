package utils

import (
	"database/sql"
	"log"
)

func UtilParsingDxxxB(rows *sql.Rows, err error) []map[string]interface{} {
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
			// b, ok := val.([]byte)
			// if ok {
			// 	row[col] = string(b)
			// } else {
			// 	row[col] = val
			// }
			row[col] = val
		}
		result = append(result, row)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
