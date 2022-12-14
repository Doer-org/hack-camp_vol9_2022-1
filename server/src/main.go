package main

import (
	"github.com/Doer-org/hack-camp_vol9_2022-1/infrastructure/database"
	"github.com/Doer-org/hack-camp_vol9_2022-1/presentation/router"
)

func main() {
	db, err := database.NewDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := router.NewRouter()
	r.InitRoomRouter(db)
	r.InitMemberRouter(db)
	r.InitChatRouter(db)
	r.InitHealthRouter()

	r.Serve()
}
