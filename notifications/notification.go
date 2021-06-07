package notifications

import (
	"fmt"
	"github.com/ravikondu/notification-app/constants/notificationType"
	userDto "github.com/ravikondu/notification-app/dto/user"
	userEntity "github.com/ravikondu/notification-app/entities/user"
	userDao "github.com/ravikondu/notification-app/dao/user"
	"github.com/ravikondu/notification-app/notifications/call"
	"github.com/ravikondu/notification-app/notifications/email"
	"github.com/ravikondu/notification-app/notifications/sms"
	"runtime/debug"
	"sync"
)

type Notifier interface {
	SendNotification(user *userEntity.User) *userDto.NotificationResponse
}

func GetNotificationHandlers(preferredNotificationTypes []string) []Notifier {
	fmt.Println("preferred notification types ", preferredNotificationTypes)
	notificationHandlers := make([]Notifier, 0)
	for _, preferredNotificationType := range preferredNotificationTypes {
		switch preferredNotificationType {
		case notificationType.EMAIL.Name():
			notificationHandlers = append(notificationHandlers, &email.EmailNotifcation{})
		case notificationType.CALL.Name():
			notificationHandlers = append(notificationHandlers, &call.PhoneCallNotifcation{})
		case notificationType.SMS.Name():
			notificationHandlers = append(notificationHandlers, &sms.SmsNotifcation{})
		default:
			fmt.Println("invalid notification type received: ", preferredNotificationType)
		}
	}
	return notificationHandlers
}

func SendNotificationToUsers() {
	users, err := userDao.GetAllUsersList()
	if err != nil {
		fmt.Println("Error in fetching users list")
		return
	}
	wg := sync.WaitGroup{}
	const maxRoutines = 10 // This can vary as per requirement
	channel := make(chan int, maxRoutines)
	for _, user := range users {
		wg.Add(1)
		channel <- 1
		go func(user userEntity.User) {
			defer func() {
				r := recover()
				if r != nil {
					fmt.Println("panic occured, stack trace for the same: ", string(debug.Stack()))
				}
				wg.Done()
			}()
			notificationHandlers := GetNotificationHandlers(user.PreferredNotificationType)
			for _, notificationHandler := range notificationHandlers {
				response := notificationHandler.SendNotification(&user)
				fmt.Println("Response for notification ", response)
			}
			<- channel
		}(user)

	}
	close(channel)
	wg.Wait()
}
