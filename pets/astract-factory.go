package pets

import (
	"errors"
	"fmt"
	"go-breeder/models"
)

type AnimalInterface interface {
	Show() string
}

type DogFromFactory struct {
	Pet *models.Dog
}
type CatFromFactory struct {
	Pet *models.Cat
}

func (dff *DogFromFactory) Show() string {
	return fmt.Sprintf("This animal is a %s", dff.Pet.Breed.Breed)
}
func (cff *CatFromFactory) Show() string {
	return fmt.Sprintf("This animal is a %s", cff.Pet.Breed.Breed)
}

type PetFactoryInterface interface {
	newPet() AnimalInterface
}

type DogAbstractFactory struct {
}
type CatAbstractFactory struct {
}

func (df *DogAbstractFactory) newPet() AnimalInterface {

	return &DogFromFactory{
		Pet: &models.Dog{},
	}
}
func (df *CatAbstractFactory) newPet() AnimalInterface {

	return &CatFromFactory{
		Pet: &models.Cat{},
	}
}

func NewPetFromAbstractFactory(species string) (AnimalInterface, error) {

	switch species {

	case "dog":
		var dogFactory DogAbstractFactory
		dog := dogFactory.newPet()
		return dog, nil

	case "cat":
		var catFactory CatAbstractFactory
		cat := catFactory.newPet()
		return cat, nil

	default:
		return nil, errors.New("Invalid pet type")
	}
}
