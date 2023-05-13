package main

import (
	"context"
	"fmt"
	"time"

	"github.com/christian-gama/nutrai-api/internal/shared/infra/env"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/sql"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/sql/querying"
	"github.com/christian-gama/nutrai-api/internal/user/app/command"
	"github.com/christian-gama/nutrai-api/internal/user/app/query"
	"github.com/christian-gama/nutrai-api/internal/user/app/service"
	value "github.com/christian-gama/nutrai-api/internal/user/domain/value/user"
	"github.com/christian-gama/nutrai-api/internal/user/infra/hash"
	persistence "github.com/christian-gama/nutrai-api/internal/user/infra/persistence/sql"
	"github.com/go-faker/faker/v4"
)

func main() {
	env.Load(".env.dev")
	db := sql.MakePostgres()
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	patientRepo := persistence.NewSQLPatient(db)
	hashPasswordHandler := service.NewHashPasswordHandler(hash.New())
	savePatientHandler := command.NewSavePatientHandler(patientRepo, hashPasswordHandler)
	ctx := context.Background()

	email := value.Email(faker.Email())
	start := time.Now()
	err = savePatientHandler.Handle(ctx, &command.SavePatientInput{
		Age: 18,
		User: &command.SaveUserInput{
			Name:     "christian",
			Email:    email,
			Password: "12345678",
		},
		HeightM:  1.8,
		WeightKG: 70,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Took %v to save patient\n", time.Since(start))

	findPatientHandler := query.NewFindPatientHandler(patientRepo)
	patient, err := findPatientHandler.Handle(ctx, &query.FindPatientInput{
		ID:      5,
		Preload: querying.AddPreload("user").Slice(),
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Saved patient: %+v\n", patient)
}
