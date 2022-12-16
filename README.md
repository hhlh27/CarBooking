# CarBooking
CarBooking is a ride-sharing platform using microservice architecture. The platform has 2 primary group of users, namely the passengers and drivers. During creation of passenger account, first name, last name, mobile number, and email address are required. Subsequently, users can update any information in their account, but they are not able to delete their accounts for audit purposes.For driver account creation, first name, last name, mobile number, email address, identification number and car license number are required. Drivers can update all information except their identification number. Similarly, a driver account cannot be deleted.

A passenger can request for a trip with the postal codes of the pick-up and drop-off location. The platform will assign an available driver, who is not driving a passenger, to the trip. This driver will then be able to initiate a start trip or end trip. The passenger can retrieve all trips he/she has taken before in reverse chronological order.

Passenger can
-	Sign up
-	Update info
-	Request trip(status- pending)
-	Retrieve trips 

Driver can
-	Sign up
-	Update info
-	initiate a start trip - accept assignment(status confirmed) , 
- end trip -destination reached(status completed)

Design considerations
1. Each service is responsible for a single part of the functionality. Each microservice will handle only one functionality for a certain user only(either driver or passenger). For example, one microservice for driver account management and one microservice for passenger account management. Having a single concern makes the microservice easier to maintain and scale.

2.	Having different databases for each microservice. Each database will store the data needed for the microservice , and used solely for that microservice. For example, passenger account management microservice has its own passenger Database and driver account management has its own driver database.

3.	Designing services to be loosely coupled, have high cohesion, and cover a single bounded context. A loosely coupled service depends minimally on other services. Having high cohesion requires that the design of the service should follow the single responsibility principle. It should perform only one main function and do it well. Each microservice is designed so that it has minimal dependencies on other microservice. This ensures that a microservice covers a single-bounded context, achieving a Domain-Driven Design (DDD).
Architecture diagram
![image](https://user-images.githubusercontent.com/104514493/208100257-796a7371-08e2-4273-8e41-03676cfd1468.png)

Setting up and running your microservices



Passenger user
Passenger account management
To create new passenger account, choose option 1 
Enter your identification number, first name, mobile number, email and password. There will be a message displayed if account is created successfully.

To update your account, enter option 2. Enter your identification number. The system will then prompt for your updated details. There will be a message displayed if account is updated successfully.

Passenger trip booking management
To book a car ride, choose option 1. Enter your pickup postal code and drop off postal code. A booking will be created with status pending(waiting for a driver to accept the assignment). If a driver has accepted the assignment, the status will changed to confirmed
To view all past bookings, choose option 2. The system will display all booking records in reverse chronological order

Driver User
Driver account management
To create new driver account, choose option 1 
Enter your identification number, first name, car number, mobile number, email and password. There will be a message displayed if account is created successfully.

To update your account, enter option 2. Enter your identification number. The system will then prompt for your updated details. There will be a message displayed if account is updated successfully.

Driver trip booking management
To accept an assignment, choose option 1. The booking status for the booking record will change to ‘confirmed’. The driver can then proceed to pick up the passenger at the pickup location.
To end the trip, choose option 2. The booking status for the booking record will change to ‘completed’. The passenger can then alight at the destination and the trip has been completed successfully.
