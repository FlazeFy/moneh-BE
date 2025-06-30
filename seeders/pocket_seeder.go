package seeders

import (
	"log"
	"moneh/factories"
	"moneh/modules/pocket"
	"moneh/modules/user"
)

func SeedPockets(repo pocket.PocketRepository, userRepo user.UserRepository, count int) {
	// Empty Table
	repo.DeleteAll()

	// Fill Table
	var success = 0
	for i := 0; i < count; i++ {
		pocket := factories.PocketFactory()
		user, err := userRepo.FindOneRandom()
		if err != nil {
			log.Printf("failed to seed pocket %d: %v\n", i, err)
		}

		err = repo.CreatePocket(&pocket, user.ID)
		if err != nil {
			log.Printf("failed to seed pocket %d: %v\n", i, err)
		}
		success++
	}
	log.Printf("Seeder : Success to seed %d Flow", success)
}
