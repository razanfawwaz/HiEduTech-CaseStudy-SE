# Study Case (Software Engineering Data Specialist)

## Introduction
This is a study case project repository for the Software Engineering Data Specialist from [HiEduTech (GovTechEdu)](https://hiedutech.softr.app/). Thank you for [HiEduTech (GovTechEdu)](https://hiedutech.softr.app/) team for giving me this opportunity to learn and improve my skills. 

## Project Description
This project owned by `Muhammad Razan Fawwaz` Undergraduate informatics student at Universitas Syiah Kuala. Using Go Language and Postgres Library for Golang.

## Software Requirements
- Go 
- Postgres 
- Docker 
- Sqlite3 

## How To Use
1. Clone this repository
2. Install Postgres using Docker
3. Run `go run main.go` on your terminal
4. Program will return `success` if the data is successfully exported and imported to Postgres database.

## How Program Works
1. Program will convert the data from `Northwind_small.sqlite` to `data.sql`.
2. After successfully converted. Program will import the data to Postgres database using the `data.sql` file.
3. If the data is successfully imported, the program will return `success`.
4. Program will print out the `amount of data` in each table.
5. Program also will print out the `amount of time` that used to export, import, and count the data.
6. Program exit

