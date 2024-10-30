package main

import (
	"fmt"
	"strings"
)

func main(){
    var conferenceName = "Go Conference"
    const conferenceTickets = 50
    var remainingTickets = 50 
    var bookings []string

    fmt.Printf("Welcome to %v conference booking application\n", conferenceName);
    fmt.Println("We have total of" ,conferenceTickets,"tickets and",remainingTickets,"are still available.");
    fmt.Println("Get your ticket here to attend");
    

    for{
        var firstName string
        var lastName string
        var email string
        var userTickets int 
    
        fmt.Println("Enter your first name: ")
        fmt.Scan(&firstName)
    
        fmt.Println("Enter your last name: ")
        fmt.Scan(&lastName)
    
        fmt.Println("Enter your email address: ")
        fmt.Scan(&email)
    
        fmt.Println("Enter your user tickets: ")
        fmt.Scan(&userTickets)

        isValidName := len(firstName) >= 2 && len(lastName) >= 2
        isValidEmail := strings.Contains(email,"@")
        isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

      
        if isValidName && isValidEmail && isValidTicketNumber {
            remainingTickets = remainingTickets - userTickets
            bookings = append(bookings,firstName + " " + lastName)
        
            fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
            fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
    
            firstNames := []string{}
            for _, booking := range bookings {
                var names = strings.Fields(booking)
                firstNames = append(firstNames, names[0])
            }
            fmt.Printf("These are all our bookings: %v\n",firstNames)
    
            if remainingTickets == 0 {
                fmt.Printf("Our conference is booked out. Come back next year.")
                break
            }
        }else{       
            if !isValidName{
                fmt.Println("first name or last name you entered is too short")
            }
            if !isValidEmail{
                fmt.Println("email address you entered doesn't contain @ sign")

            }
            if !isValidTicketNumber{
                fmt.Println("number of tickets you entered is invalid")
            }
            fmt.Printf("Your Input is Invalid\n")
        }
        
       
    }
   
}