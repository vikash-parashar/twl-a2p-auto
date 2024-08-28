// Step 6: Add a Phone Number to the Messaging Service

package a2p

import (
	"fmt"
	"os"

	messaging "github.com/twilio/twilio-go/rest/messaging/v1"
)

type PhoneNumberData struct {
	PhoneNumberSid string
}

// Step 7.1: Fetch Available Phone Numbers By Country Code
// TODO: Add a function to list phone numbers by country code

// Step 7.2: Purchase a Phone Number only if you don't have one
// TODO: Add a function to purchase a phone number from twilio

// Step 7.3: Add a Phone Number to the Messaging Service , once you have a phone number, you can associate it with the messaging service
func (s *A2PService) AddPhoneNumberToMessagingService(serviceSid string, params *messaging.CreatePhoneNumberParams) (*messaging.MessagingV1PhoneNumber, error) {
	// Add the phone number to the Messaging Service
	resp, err := s.client.MessagingV1.CreatePhoneNumber(serviceSid, params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		if resp.Sid != nil {
			fmt.Println("Phone number added successfully with SID:", *resp.Sid)
		} else {
			fmt.Println("Failed to retrieve SID from response.")
		}
	}
	return resp, nil
}
