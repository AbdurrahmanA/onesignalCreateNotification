# OnesignalCreateNotification

```go
	onesignal.NewClient()
	var app onesignal.AppCreate
	app.AppID = "YourOnesignalAppID"

	//Notification for all users in your app
	err := app.CreateNotification("YourMessage", "YourTitle")
	if err != nil {
		fmt.Println(err.Error())
	}
	//Notification for user or users in your app
	var id []string
	id = append(id, "UserID")
	err = app.CreateNotification("YourMessage", "YourTitle", id)
	if err != nil {
		fmt.Println(err.Error())
	}

```
