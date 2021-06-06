package notificationType

type notificationType int

const (
	EMAIL notificationType = iota
	CALL
	SMS
)

var notificationtypeName = [...]string{
	"email",
	"call",
	"sms",
}

func (nt notificationType) Name() string{
	return notificationtypeName[nt]
}




