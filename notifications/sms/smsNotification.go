package sms

import (
	"github.com/ravikondu/notification-app/entities/user"
	userDto "github.com/ravikondu/notification-app/dto/user"
	"fmt"

)

type SmsNotifcation struct {}

func (e *SmsNotifcation) SendNotification(user *user.User) (response *userDto.NotificationResponse)  {
	// send sms notification and update response
	fmt.Println("sms notification sent")
	return response
}