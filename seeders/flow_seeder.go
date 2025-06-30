package seeders

import (
	"log"
	"moneh/factories"
	"moneh/modules/flow"
	"moneh/modules/user"
)

func SeedFlows(repo flow.FlowRepository, userRepo user.UserRepository, count int) {
	// Empty Table
	repo.DeleteAll()

	// Fill Table
	var success = 0
	for i := 0; i < count; i++ {
		flow := factories.FlowFactory()
		user, err := userRepo.FindOneRandom()
		if err != nil {
			log.Printf("failed to seed flow %d: %v\n", i, err)
		}

		err = repo.CreateFlow(&flow, user.ID)
		if err != nil {
			log.Printf("failed to seed flow %d: %v\n", i, err)
		}
		success++
	}
	log.Printf("Seeder : Success to seed %d Flow", success)
}
