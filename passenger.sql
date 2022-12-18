CREATE database passenger_db;
USE passenger_db;

/* Table: Passenger */
CREATE TABLE Passenger
(
PassengerID varchar (9) NOT NULL,
FirstName varchar(50) 	NOT NULL,
LastName varchar(50) 	NOT NULL,
Mobile  varchar(8)  	NOT NULL,
Email	varchar(50)  	NOT NULL,
AccPassword varchar(255)  	NOT NULL,
CONSTRAINT PK_Passenger PRIMARY KEY NONCLUSTERED (PassengerID)
)
;

INSERT INTO Passenger (PassengerID, FirstName, LastName, Mobile,Email, AccPassword) VALUES ("S1234567S", "Jake", "Lee", "99991111",'jakelee@gmail.com','password123');
