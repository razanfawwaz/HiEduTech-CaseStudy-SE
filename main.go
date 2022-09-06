package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func connect() (*sql.DB, error) {
	host := "localhost"
	port := "5432"
	user := "user"
	password := "postgres"
	dbname := "postgres"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Database failed to connect : %s", err)
	}

	return db, nil
}

func dumpDb() {
	timeStart := time.Now()
	//app := "sqlite3 /Northwind_small.sqlite .dump > db.sql"
	dbpath := "./Northwind_small.sqlite"
	cmd := exec.Command("sqlite3", dbpath, ".dump")
	data, err := cmd.Output()

	f, err := os.Create("data2.sql")
	if err != nil {
		log.Fatal(err)
	}
	t := strings.Replace(string(data), "BLOB", "bytea", -1)
	t2 := strings.Replace(t, "PRAGMA foreign_keys=OFF;", "", -1)
	t3 := strings.Replace(t2, "[", "", -1)
	t4 := strings.Replace(t3, "]", "", -1)
	t5 := strings.Replace(t4, `"`, "", -1)
	t6 := strings.Replace(t5, "IF NOT EXISTS", "", -1)
	t7 := strings.Replace(t6, "CREATE TABLE  Order \n(", `CREATE TABLE "Order" (`, -1)
	t8 := strings.Replace(t7, "INSERT INTO Order VALUES", `INSERT INTO "Order" VALUES`, -1)
	t9 := strings.Replace(t8, "DOUBLE", "DOUBLE PRECISION", -1)

	defer f.Close()
	_, err = f.WriteString(t9)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Export data success!")
	}

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	timeElapsed := time.Since(timeStart)
	fmt.Printf("Total time for export data: %s\n", timeElapsed)
}

func importData() {
	timeStart := time.Now()

	db, err := connect()
	query, _ := os.ReadFile("./data2.sql")
	_, err = db.Query(string(query))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("\nInput Data Success!")
	}
	timeElapsed := time.Since(timeStart)
	fmt.Printf("Total time for input data: %s\n\n", timeElapsed)
}

func countData() {
	timeStart := time.Now()
	var count int
	var tables []string = []string{"Category", "Customer", "CustomerCustomerDemo", "CustomerDemographic", "Employee", "EmployeeTerritory", "Order", "OrderDetail", "Product", "Region", "Shipper", "Supplier", "Territory"}
	db, err := connect()
	for i := 0; i < len(tables); i++ {
		db.QueryRow("SELECT COUNT(*) FROM " + tables[i]).Scan(&count)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Total data from table", tables[i], ":", count)
	}
	timeElapsed := time.Since(timeStart)
	fmt.Printf("Total time for data count: %s\n", timeElapsed)

}

func main() {
	timeStart := time.Now()
	dumpDb()
	importData()
	countData()
	timeElapsed := time.Since(timeStart)
	fmt.Printf("\nTotal time elapsed: %s", timeElapsed)
}
