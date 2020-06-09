package onesignal

//INotification IPush
type iNotification interface {
	getMessage() string
	getTitle() string
	getID() [][]string
}

//Notification Notification
type notification struct {
	Message string
	Title   string
	ID      [][]string
}

func (a notification) getMessage() string {
	return a.Message
}
func (a notification) getTitle() string {
	return a.Title
}
func (a notification) getID() [][]string {
	return a.ID
}

//ForAllUsers ForAllUsers
type forAllUsers struct {
	notification
}

//ForUser ForUser
type forUser struct {
	notification
}

func (a forUser) newPushForUsers() iNotification {
	return &forUser{
		notification: notification{
			Message: a.Message,
			Title:   a.Title,
			ID:      a.ID,
		},
	}
}

func (a forAllUsers) newPushForAllUsers() iNotification {
	var all [][]string
	var s []string
	s = append(s, "All")
	all = append(all, s)
	return &forAllUsers{
		notification: notification{
			Message: a.Message,
			Title:   a.Title,
			ID:      all,
		},
	}
}

//NotificationFactory NotificationFactory
func notificationFactory(a iNotification) notification {
	return notification{
		Message: a.getMessage(),
		Title:   a.getTitle(),
		ID:      a.getID(),
	}
}

//AppCreate AppCreate
type AppCreate struct {
	AppID string
}

//CreateNotification CreateNotification
func (app AppCreate) CreateNotification(msg string, title string, id ...[]string) error {
	if len(id) == 0 {
		forAllUsersnoti := forAllUsers{notification{Message: msg, Title: title}}
		adapted := notificationForAllUsers{}
		notiAdapter := notificationForAllUsersAdapter{
			users: adapted,
		}
		iNotification := forAllUsersnoti.newPushForAllUsers()
		notification := notificationFactory(iNotification)
		err := push(notiAdapter, notification, app)
		return err
	}
	forUsernoti := forUser{notification{ID: id, Title: title, Message: msg}}

	iNotification := forUsernoti.newPushForUsers()
	notification := notificationFactory(iNotification)
	notificationType := notificationForUser{}
	err := push(notificationType, notification, app)
	return err

}
