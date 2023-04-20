package server

func Migrate() {
	database := migrateDB()

	defer func() {
		sqlDB, _ := database.DB()
		sqlDB.Close()
	}()
}
