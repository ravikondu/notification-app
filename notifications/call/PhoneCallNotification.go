package call

import (
	"github.com/ravikondu/notification-app/entities/user"
	userDto "github.com/ravikondu/notification-app/dto/user"
	"fmt"
)

type PhoneCallNotifcation struct {}

func (e *PhoneCallNotifcation) SendNotification(user *user.User) (response *userDto.NotificationResponse) {
	// send phone call notification and update response
	fmt.Println("call notification sent")
	return response
}