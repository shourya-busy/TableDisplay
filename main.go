package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)



func MarshalOrdered(data []map[string]interface{}, keys []string) ([]byte, error) {
	var jsonData []byte
	jsonData = append(jsonData, '[')

	for i, lead := range data {
		if i > 0 {
			jsonData = append(jsonData, ',')
		}

		jsonData = append(jsonData, '{')

	
		for j, key := range keys {
			if j > 0 {
				jsonData = append(jsonData, ',')
			}
			jsonData = append(jsonData, '"')
			jsonData = append(jsonData, key...)
			jsonData = append(jsonData, '"', ':')

			val, err := json.Marshal(lead[key])
			if err != nil {
				return nil, err
			}
			jsonData = append(jsonData, val...)
		}

		jsonData = append(jsonData, '}')
	}

	jsonData = append(jsonData, ']')

	return jsonData, nil
}


func main() {
	
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/leads",func(ctx *gin.Context) {

		dialect := goqu.Dialect("postgres")

		pgDb, err := sql.Open("postgres", "user=shouryagautam password=shourya dbname=postgres sslmode=disable")
		if err != nil {
		panic(err.Error())
		}
		goqudb := dialect.DB(pgDb)

		rows, err := goqudb.Query(`WITH RECURSIVE employee_hierarchy AS (
			SELECT employee_id, employee_name, manager_id
			FROM employees
			WHERE manager_id IS NULL
		
			UNION ALL
			SELECT e.employee_id, e.employee_name, e.manager_id
			FROM employees e
			INNER JOIN employee_hierarchy eh ON e.manager_id = eh.employee_id
		)
		SELECT manager_id as manager,employee_id  FROM employee_hierarchy`)

		if err != nil {
			fmt.Println(err.Error())
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		col, err := rows.Columns()
		if err != nil {
			fmt.Println(err.Error())
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		var leads []map[string]interface{}

		// Iterate over rows
		for rows.Next() {
			// Map to hold the current row's data
			row := make(map[string]interface{})

			// Slice to hold values of current row
			values := make([]interface{}, len(col))
			
			// Assign each pointer to the corresponding value in the values slice
			for i := range values {
				values[i] = new(interface{})
			}

			// Scan the current row into the pointers
			if err := rows.Scan(values...); err != nil {
				fmt.Println(err.Error())
				continue
			}

			// Populate the row map with column name and value
			for i, colName := range col {
				row[colName] = *(values[i].(*interface{}))
			}

			// Append the row map to the result
			leads = append(leads, row)
		}

		jsonData,err := MarshalOrdered(leads, col)

		if err != nil {
			fmt.Println(err.Error())
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

    	ctx.Data(200, "application/json", jsonData)
	})

	router.Run(":8080")
}


