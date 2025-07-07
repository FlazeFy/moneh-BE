package seeders

import (
	"log"
	"moneh/factories"
	"moneh/modules/errors"
)

func SeedErrors(repo errors.ErrorRepository, count int) {
	// Empty Table
	repo.DeleteAll()

	// Fill Table
	var success = 0
	for i := 0; i < count; i++ {
		admin := factories.ErrorFactory()
		err := repo.CreateError(&admin)
		if err != nil {
			log.Printf("failed to seed error %d: %v\n", i, err)
		}
		success++
	}
	log.Printf("Seeder : Success to seed %d Error", success)
}
