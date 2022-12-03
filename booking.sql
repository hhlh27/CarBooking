drop database booking_db;
CREATE database booking_db;

USE booking_db;


CREATE TABLE Booking
(
  BookingID				int NOT NULL,
  PassengerID			int      NOT NULL,
DriverID			int      NOT NULL,
PickUp			int      NOT NULL,
DropOff		int      NOT NULL,
BookingDateTime		datetime		NOT NULL ,
BookingStatus     	varchar(50)  	NOT NULL,
  CONSTRAINT PK_Booking PRIMARY KEY NONCLUSTERED (BookingID)

)
