package server

func Build() {
	database := migrateDB()

	defer func() {
		sqlDB, _ := database.DB()
		sqlDB.Close()
	}()
}
