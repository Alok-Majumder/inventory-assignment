# inventory-assignment

Running the Application locally

Prerequisites:
 1. Start the postgres DB using a docker image : For local test not doing any volume mount. So if you restart the container the data will be lost

 ``` 
  docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres
 
 ``` 




The default postgres user and postgres database are created


 2. To connect to the container, in order to do some work run the following

 ``` 
  docker exec -it some-postgres bash

  su postgres

  psql

  SELECT 1;

``` 

     OR

Via pgadmin tool. Details you can find https://www.pgadmin.org/


Set the enviroment variables:

export DB_USER=postgres
export DB_PASSWORD=mysecretpassword
export DB_HOST=127.0.0.1
export DB_PORT=5432
export DB_NAME=postgres


Create INVENTORY TABLE:

CREATE TABLE INVENTORY (
    ART_ID varchar(50) PRIMARY KEY ,
    NAME varchar(150) NOT NULL,
    STOCK NUMERIC(5,2)
)



