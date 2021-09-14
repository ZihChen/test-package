package main

import (
	"test-package-02/app/global"
	"test-package-02/internal/database"
	"test-package-02/internal/entry"
)

func main()  {
	global.Start()

	database.NewConnect()

	entry.Run()
}