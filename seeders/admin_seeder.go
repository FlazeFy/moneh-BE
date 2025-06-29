package seeders

import (
	"log"
	"moneh/factories"
	"moneh/modules/admin"
)

func SeedAdmins(repo admin.AdminRepository, count int) {
	// Empty Table
	repo.DeleteAll()

	// Fill Table
	var success = 0
	for i := 0; i < count; i++ {
		admin := factories.AdminFactory()
		err := repo.CreateAdmin(&admin)
		if err != nil {
			log.Printf("failed to seed admin %d: %v\n", i, err)
		}
		success++
	}
	log.Printf("Seeder : Success to seed %d Admin", success)
}
