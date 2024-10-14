package pets

import "go-breeder/models"

func NewPet(species string) *models.Pet {

	pet := models.Pet{
		Species:     species,
		Breed:       "",
		MinWeight:   0,
		MaxWeight:   0,
		Description: "No description",
		LifeSpan:    0,
	}
	return &pet
}
