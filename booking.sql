
CREATE database booking_db;

USE booking_db;


CREATE TABLE Booking
(
  BookingID				varchar (20) NOT NULL,
  PassengerID			varchar (9)     NOT NULL,
DriverID			varchar (9)      NOT NULL,
PickUp			varchar (6)     NOT NULL,
DropOff		varchar (6)     NOT NULL,
BookingDateTime		varchar(20) 		NOT NULL ,
BookingStatus     	varchar(15)  	NOT NULL,
  CONSTRAINT PK_Booking PRIMARY KEY NONCLUSTERED (BookingID)

)
