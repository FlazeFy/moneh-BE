package seeders

import (
	"log"
	"moneh/factories"
	"moneh/modules/feedback"
	"moneh/modules/user"
)

func SeedFeedbacks(repo feedback.FeedbackRepository, userRepo user.UserRepository, count int) {
	// Empty Table
	repo.DeleteAll()

	// Fill Table
	var success = 0
	for i := 0; i < count; i++ {
		feedback := factories.FeedbackFactory()
		user, err := userRepo.FindOneRandom()
		if err != nil {
			log.Printf("failed to seed feedback %d: %v\n", i, err)
		}

		err = repo.CreateFeedback(&feedback, user.ID)
		if err != nil {
			log.Printf("failed to seed feedback %d: %v\n", i, err)
		}
		success++
	}
	log.Printf("Seeder : Success to seed %d Feedback", success)
}
