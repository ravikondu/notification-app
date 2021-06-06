package email

import (
	"fmt"
	userDto "github.com/ravikondu/notification-app/dto/user"
	"github.com/ravikondu/notification-app/entities/user"
)

type EmailNotifcation struct{}

func (e *EmailNotifcation) SendNotification(user *user.User) (response *userDto.NotificationResponse) {
	// send email notification and update response
	fmt.Println("email notification sent")
	return response
}
