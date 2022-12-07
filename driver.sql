CREATE database driver_db;

USE driver_db;

CREATE TABLE Driver
(
DriverID 				varchar (5) NOT NULL,
FirstName				varchar(50) 	NOT NULL,
LastName				varchar(50) 	NOT NULL,
IdenNo				varchar(9) 	NOT NULL,
CarNo				varchar(8) 	NOT NULL,
Mobile  	    	  varchar(8)  	NOT NULL,
Email		    	varchar(50)  	NOT NULL,
AccPassword		    varchar(255)  	NOT NULL,
CONSTRAINT PK_Driver PRIMARY KEY NONCLUSTERED (DriverID)
);

INSERT INTO Driver (DriverID, FirstName, LastName,IdenNo,CarNo, Mobile,Email, AccPassword) VALUES ("D0001", "Tom", "Lim",'S1234567A','SDF1234A', "92221111",'tomlee@gmail.com','password123');

