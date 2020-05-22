# OnesignalCreateNotification

```go
onesignal.NewClient()
var app onesignal.AppCreate
app.AppID="YourOnesignalAppID"

//Notification for all users in your app
app.CreateNotification("YourMessage","YourTitle")

//Notification for user or users in your app
var id []string
id=append(id,"UserID")
app.CreateNotification("YourMessage","YourTitle",id)

```
