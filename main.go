package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

const conferenceTickets = 50
var conferenceName = "Go conference"
var remainingTickets = 50
var bookings = make([]UserData,0)

type UserData struct{
	firstName string
	lastName string
	email string
	numberOfTickets int

}

func main(){

	

	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {
	

	firstName,lastName, email, userTickets :=  getUserInput()	
	isValidName,isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName,email, userTickets,remainingTickets)

    

	if isValidName && isValidEmail && isValidTicketNumber {

		

		bookTickets(userTickets,firstName,lastName,email)
		go sendTicket(userTickets,firstName,lastName,email)
		firstNames :=getFirstNames()
		fmt.Printf("These are the first names of all the bookings: %v\n",firstNames)
	
		


		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out.Come back next year.")
			break

		}
	}else {

		if !isValidEmail{
			fmt.Println("input email address is wrong")
		}

		if !isValidName{
			fmt.Println("input name is too short")
		}

		if !isValidTicketNumber{
			fmt.Println("input ticket number is wrong")
		}

		
	}

	}

	
}


func greetUsers(){
	fmt.Printf("Welcome to %v \n",conferenceName)
	fmt.Println("We have total of ",conferenceTickets," tickets and ",remainingTickets,"are still available.")
	fmt.Println("Get your tickets here to attend")
	fmt.Printf("Conference Tickets is %T, Remaining Tickets is %T, Conference Name is %T\n",conferenceTickets,remainingTickets,conferenceName)
}


func getFirstNames() []string{
	firstNames := []string{}
		for _,booking := range bookings{
		
		firstNames = append(firstNames, booking.firstName)
		}

		

		return firstNames

}




func getUserInput()(string,string,string,int){

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter firstname:")
	fmt.Scan(&firstName)
	fmt.Println("Enter lastname:")
	fmt.Scan(&lastName)
	fmt.Println("Enter email:")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets booked by user:")
	fmt.Scan(&userTickets)

   return firstName,lastName,email,userTickets

}

func bookTickets(userTickets int,firstName string,lastName string,email string){
		remainingTickets=remainingTickets - userTickets


		var userData = UserData{
			firstName: firstName,
			lastName: lastName,
			email: email,
			numberOfTickets: userTickets,
		}
		
		bookings=append(bookings, userData)
		fmt.Printf("List of bookings is %v \n",bookings)



		fmt.Printf("Whole slice is: %v\n",bookings)
		fmt.Printf("The first value is: %v\n",bookings[0])
		fmt.Printf("Slice type: %T\n",bookings)
		fmt.Printf("Slice length is: %v\n",len(bookings))



		fmt.Printf("THANK YOU %v %v for booking %v tickets. You will receive a confirmation mail at %v\n", firstName , lastName ,userTickets, email)
		fmt.Printf("%v tickets remaining for %v \n",remainingTickets,conferenceName)

}

func sendTicket(userTickets int, firstName string, lastName string,email string){

	time.Sleep(10*time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("######")
	fmt.Printf("Sending ticket:\n %v \n to email address %v \n",ticket,email)
	fmt.Println("######")
}