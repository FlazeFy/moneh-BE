package factories

import (
	"moneh/models"

	"github.com/google/uuid"
)

func FlowRelationFactory(flowRelationAmmount int, flowID, pocketID uuid.UUID) models.FlowRelation {
	return models.FlowRelation{
		Ammount:  flowRelationAmmount,
		FlowId:   flowID,
		PocketId: pocketID,
	}
}
