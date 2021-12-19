# Data Load Service
Couldn't compelete the full fledge functionality
1. This Service will load the data from File to Postgres DB
2. It can be run from command line argument as below. Provided that you have input json file ready


```go
go run main.go [FILENAME]
```
- `FILENAME` - The name of the file to load

## Pending Work
1. Most the functionnality is incomplete 
2. Only Basic coding standard has been established as per the go standard and best practices. 


## Folder Description


    ├── inventory-assignment  #This folder contains the API Service for Sell and Get Product Details
    ├── inventory-load       #This folder contains datalaod service to load data from file to PostgresDB
