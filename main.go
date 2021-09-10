package main

import (
	"test-package-02/app/global/helper"
	"test-package-02/internal/database"
	"test-package-02/internal/entry"
)

func main()  {
	database.NewConnect()
	helper.PrintLog("rrrr")
	entry.Run()
}