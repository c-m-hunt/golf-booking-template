# Golf Booking Template

Produces K8s cron for month from template

Known issue: Needs adjusting for BST/DST

## Running
Add a `.env` file and adjust the parameters
```
MONTH=Apr,2022
DAY=Friday
TIME=08:12
PLAYERS=PLAYER1,PLAYER2
COURSE=SouthEast
```

Run
```
go run main.go
```

Output files are now in `./out`.

You can then send these to K8s with
```
kubectl apply -f out
```