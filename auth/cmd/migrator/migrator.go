package migrator

import (
	"context"
	"fmt"
	"github.com/oreshkanet/aTES/auth/internal/config"
	"github.com/oreshkanet/aTES/auth/internal/repository"
	"github.com/oreshkanet/aTES/packages/pkg/database"
	"log"
)

func main() {
	// TODO: Запуск приложения для миграции БД

	conf := config.Load()

	ctx := context.Background()

	// Создаём подключение к БД
	dbURL := fmt.Sprintf(
		"sqlserver://%s:%s@%s?database=%s",
		conf.MsSqlUser, conf.MsSqlPwd,
		conf.MsSqlHost, conf.MsSqlDb,
	)
	db, err := database.NewDBMsSQL(ctx, dbURL)
	if err != nil {
		log.Fatalf("%s", err.Error())
		return
	}
	defer db.Close()

	repos := repository.NewRepository(db)
	if err := repos.User.MigrateUp(); err != nil {
		log.Fatalf("%s", err.Error())
		return
	}
}
