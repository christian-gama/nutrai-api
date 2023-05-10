package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/christian-gama/nutrai-api/internal/shared/infra/env"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/sql"
	"github.com/christian-gama/nutrai-api/internal/user/app/command"
	"github.com/christian-gama/nutrai-api/internal/user/app/service"
	"github.com/christian-gama/nutrai-api/internal/user/infra/hash"
	"github.com/christian-gama/nutrai-api/internal/user/infra/persistence"
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

	patientRepo := persistence.NewPatient(db)
	hashPasswordHandler := service.NewHashPasswordHandler(hash.New())
	savePatientHandler := command.NewSavePatientHandler(patientRepo, hashPasswordHandler)
	ctx := context.Background()

	err = savePatientHandler.Handle(ctx, &command.SavePatientInput{
		Age: 18,
		User: &command.SaveUserInput{
			Name:     faker.Name(),
			Email:    faker.Email(),
			Password: "12345678",
		},
		HeightM:  1.8,
		WeightKG: 70,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Patient created successfully!")

	server := http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 5,
	}
	defer server.Close()

	server.ListenAndServe()
}
