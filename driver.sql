CREATE database driver_db;

USE driver_db;
/* Table: Driver */
CREATE TABLE Driver
(
DriverID 			varchar (9)     NOT NULL,
FirstName			varchar(50) 	NOT NULL,
LastName			varchar(50) 	NOT NULL,
CarNo				varchar(8) 	    NOT NULL,
Mobile  	    	varchar(8)  	NOT NULL,
Email		    	varchar(50)  	NOT NULL,
AccPassword		    varchar(255)  	NOT NULL,
CONSTRAINT PK_Driver PRIMARY KEY NONCLUSTERED (DriverID)
);

INSERT INTO Driver (DriverID, FirstName, LastName,IdenNo,CarNo, Mobile,Email, AccPassword) VALUES ('S1234567A', "Tom", "Lim",'SDF1234A', "92221111",'tomlee@gmail.com','password123');

