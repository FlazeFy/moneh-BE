package seeders

import (
	"log"
	"moneh/factories"
	"moneh/modules/flow"
	"moneh/modules/pocket"
	"moneh/modules/user"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

func SeedFlowRelations(repo flow.FlowRelationRepository, userRepo user.UserRepository, flowRepo flow.FlowRepository, pocketRepo pocket.PocketRepository, count int) {
	// Empty Table
	repo.DeleteAll()

	// Fill Table
	var success = 0
	for i := 0; i < count; i++ {
		user, err := userRepo.FindOneHasFlowAndPocketRandom()
		if err != nil {
			log.Printf("failed to seed flow relation %d: %v\n", i, err)
		}

		flow, err := flowRepo.FindOneRandomByUserID(user.ID)
		if err != nil {
			log.Printf("failed to seed flow relation %d: %v\n", i, err)
		}

		pocket, err := pocketRepo.FindOneRandomByUserID(user.ID)
		if err != nil {
			log.Printf("failed to seed flow relation %d: %v\n", i, err)
		}

		// Make some of flow split into multiple pocket
		var flowRelationAmmount int
		if gofakeit.Bool() {
			var maxSplit int
			if pocket.PocketAmmount < 100000 {
				maxSplit = 1
			} else if pocket.PocketAmmount < 1000000 {
				maxSplit = 2
			} else if pocket.PocketAmmount < 10000000 {
				maxSplit = 3
			} else {
				maxSplit = 4
			}

			// Take 60% - 90% of the pocket ammount
			min := int(0.6 * float64(pocket.PocketAmmount))
			max := int(0.9 * float64(pocket.PocketAmmount))
			targetTotal := gofakeit.Number(min, max)

			remaining := targetTotal
			splits := make([]int, 0, maxSplit)

			for i := 0; i < maxSplit; i++ {
				if i == maxSplit-1 {
					splits = append(splits, remaining)
				} else {
					maxSplit := remaining / (maxSplit - i)
					part := gofakeit.Number(1, maxSplit)
					splits = append(splits, part)
					remaining -= part
				}
			}

			for _, dt := range splits {
				err := addFlowRelation(repo, flow.ID, pocket.ID, user.ID, dt, i)
				if err == nil {
					success++
				}
			}
		} else {
			flowRelationAmmount = pocket.PocketAmmount
			err := addFlowRelation(repo, flow.ID, pocket.ID, user.ID, flowRelationAmmount, i)
			if err == nil {
				success++
			}
		}
	}
	log.Printf("Seeder : Success to seed %d Flow Relation", success)
}

func addFlowRelation(repo flow.FlowRelationRepository, flowID, pocketID, userID uuid.UUID, ammount, idx int) error {
	flowRelation := factories.FlowRelationFactory(ammount, flowID, pocketID)
	_, err := repo.CreateFlowRelation(&flowRelation, userID)
	if err != nil {
		log.Printf("failed to seed flow relation %d: %v\n", idx, err)
		return err
	}

	return nil
}
