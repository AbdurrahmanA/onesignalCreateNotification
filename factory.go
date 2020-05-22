package onesignal

//INotification IPush
type INotification interface {
	getMessage() string
	getTitle() string
	getID() [][]string
}

//Notification Notification
type Notification struct {
	Message string
	Title   string
	ID      [][]string
}

func (a Notification) getMessage() string {
	return a.Message
}
func (a Notification) getTitle() string {
	return a.Title
}
func (a Notification) getID() [][]string {
	return a.ID
}

//ForAllUsers ForAllUsers
type ForAllUsers struct {
	Notification
}

//ForUser ForUser
type ForUser struct {
	Notification
}

func (a ForUser) newPushForUsers() INotification {
	return &ForUser{
		Notification: Notification{
			Message: a.Message,
			Title:   a.Title,
			ID:      a.ID,
		},
	}
}

func (a ForAllUsers) newPushForAllUsers() INotification {
	var all [][]string
	var s []string
	s = append(s, "All")
	all = append(all, s)
	return &ForAllUsers{
		Notification: Notification{
			Message: a.Message,
			Title:   a.Title,
			ID:      all,
		},
	}
}

//NotificationFactory NotificationFactory
func NotificationFactory(a INotification) Notification {
	return Notification{
		Message: a.getMessage(),
		Title:   a.getTitle(),
		ID:      a.getID(),
	}
}

var app AppCreate

//AppCreate AppCreate
type AppCreate struct {
	AppID string
}

func createNotification(msg string, title string, id ...[]string) error {
	if len(id) == 0 {
		forAllUsersnoti := ForAllUsers{Notification{Message: msg, Title: title}}
		adapted := NotificationForAllUsers{}
		notiAdapter := NotificationForAllUsersAdapter{
			users: adapted,
		}
		iNotification := forAllUsersnoti.newPushForAllUsers()
		notification := NotificationFactory(iNotification)
		err := Push(notiAdapter, notification, app)
		return err
	}
	forUsernoti := ForUser{Notification{ID: id, Title: title, Message: msg}}

	iNotification := forUsernoti.newPushForUsers()
	notification := NotificationFactory(iNotification)
	notificationType := NotificationForUser{}
	err := Push(notificationType, notification, app)
	return err

}
