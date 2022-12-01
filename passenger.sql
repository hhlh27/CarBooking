drop database account_db;
CREATE database passenger_db;

USE passenger_db;



/* Table: Passenger */
CREATE TABLE Passenger
(
  PassengerID 				 varchar (5) NOT NULL,
  FirstName				varchar(50) 	NOT NULL,
	LastName				varchar(50) 	NOT NULL,
	Mobile  	    	  int  	NOT NULL,
  Email		    	varchar(50)  	NOT NULL,
  
  AccPassword	    varchar(255)  	NOT NULL,
  CONSTRAINT PK_Passenger PRIMARY KEY NONCLUSTERED (PassengerID)
)
;





INSERT INTO Passenger (PassengerID, FirstName, LastName, Mobile,Email, AccPassword) VALUES ("P0001", "Jake", "Lee", 99991111,'jakelee@gmail.com','password123');
