drop database booking_db;
CREATE database booking_db;

USE booking_db;


CREATE TABLE Booking
(
  BookingID				varchar (5) NOT NULL,
  PassengerID			varchar (5)     NOT NULL,
DriverID			varchar (5)      NOT NULL,
PickUp			varchar (6)     NOT NULL,
DropOff		varchar (6)     NOT NULL,
BookingDateTime		datetime		NOT NULL ,
BookingStatus     	varchar(15)  	NOT NULL,
  CONSTRAINT PK_Booking PRIMARY KEY NONCLUSTERED (BookingID)

)
