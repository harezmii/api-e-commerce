package db

import (
	"fmt"
)

var Client = NewClient()

func PrismaConnection() {
	prismaConnectionError := Client.Prisma.Connect()
	if prismaConnectionError != nil {
		fmt.Printf("Prisma Connection Error")
	}

}
func PrismaDisConnection() {
	prismaDisConnectionError := Client.Prisma.Disconnect()
	if prismaDisConnectionError != nil {
		fmt.Printf("Prisma DisConnection Error")
	}
}
