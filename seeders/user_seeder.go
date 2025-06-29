package seeders

import (
	"log"
	"moneh/factories"
	"moneh/modules/user"
)

func SeedUsers(repo user.UserRepository, count int) {
	// Empty Table
	repo.DeleteAll()

	// Fill Table
	var success = 0
	for i := 0; i < count; i++ {
		user := factories.UserFactory()
		_, err := repo.CreateUser(&user)
		if err != nil {
			log.Printf("failed to seed user %d: %v\n", i, err)
		}
		success++
	}
	log.Printf("Seeder : Success to seed %d User", success)
}
