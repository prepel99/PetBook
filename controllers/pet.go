package controllers

import (
	"PetBook/models"
	"PetBook/pkg/utils"

	//"PetBook/pkg/utils"
	//"fmt"
	//"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
)

func (c *Controller) CreatePetPostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		pet := &models.Pet{}
		pet.Name = r.FormValue("nickname")
		pet.PetType = r.FormValue("pet-type")
		pet.Breed = r.FormValue("breed")
		//pet.Age = r.FormValue("age")
		//pet.Weight = r.FormValue("weight")
		pet.Gender = r.FormValue("gender")
		//petpetDescription := r.FormValue("description")
		err = c.PetStore.RegisterPet(pet)
		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/main", 301)
	}
}
func (c *Controller) CreatePetGetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.GenerateHTML(w, nil, "cabinetPet")
	}
}
