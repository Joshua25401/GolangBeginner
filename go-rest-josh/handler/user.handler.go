package handler

import (
	"log"
	"strconv"
	"time"

	"github.com/joshua25401/go-test/validators"

	"github.com/joshua25401/go-test/common/response"

	"gorm.io/gorm"

	"github.com/joshua25401/go-test/database"
	"github.com/joshua25401/go-test/model/entity"

	"github.com/gofiber/fiber/v2"
)

func ShowAllUsers(ctx *fiber.Ctx) error {
	log.Println("Accepting request from route : ", ctx.OriginalURL())
	// TODO : implement get query here
	var users []entity.User

	results := database.DB.Debug().Find(&users)

	// if there is an error when inserting data
	// do log.Fatal(error)
	if results.Error != nil {
		log.Fatal(results.Error.Error())
	}

	log.Println("Success Getting " + strconv.FormatInt(results.RowsAffected, 10) + " Rows in DB")
	success := response.BuildResponse(true, "Succesfully Get Users Data", users)
	return ctx.Status(200).JSON(success)
}

func InsertUsers(ctx *fiber.Ctx) error {
	log.Println("Accepting request from route : ", ctx.OriginalURL())
	userRequest := new(entity.User)
	// TODO : implement insert query here
	if err := ctx.BodyParser(userRequest); err != nil {
		log.Fatal(err.Error())
	}

	newUser := entity.User{
		ID:        userRequest.ID,
		Name:      userRequest.Name,
		Email:     userRequest.Email,
		Address:   userRequest.Address,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}

	if err := validators.ValidateStruct(newUser); err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "failed",
			"error":   err,
		})
	}

	errInsertUser := database.DB.Create(&newUser)

	if errInsertUser.Error != nil {
		failed := response.BuildErrorResponse("Error Inserting User", errInsertUser.Error.Error(), nil)
		return ctx.Status(404).JSON(failed)
	}

	log.Println("Creating ", strconv.FormatInt(errInsertUser.RowsAffected, 10), " User")

	success := response.BuildResponse(true, "Successfuly create new user", newUser)
	return ctx.Status(200).JSON(success)
}

func GetUserById(ctx *fiber.Ctx) error {
	uID := ctx.Params("id")

	var user entity.User

	result := database.DB.Debug().First(&user, uID)

	if result.Error != nil {
		failed := response.BuildErrorResponse("Error Getting User w/ ID "+uID, result.Error.Error(), nil)
		return ctx.Status(404).JSON(failed)
	}

	log.Println("Succes get " + strconv.FormatInt(result.RowsAffected, 10) + " Data")

	success := response.BuildResponse(true, "Success Getting Data", user)
	return ctx.Status(200).JSON(success)
}

func UpdateUser(ctx *fiber.Ctx) error {

	updateUserRequest := new(entity.User)
	if err := ctx.BodyParser(updateUserRequest); err != nil {
		failed := response.BuildErrorResponse("Internal Server Error", err.Error(), nil)
		return ctx.Status(400).JSON(failed)
	}

	uID := ctx.Params("id")
	// Check user if it's available
	// if it is then continue update
	// if not then return Error response
	var dbUser entity.User

	result := database.DB.Debug().First(&dbUser, uID)

	if result.Error != nil {
		failed := response.BuildErrorResponse("User w/ ID"+uID+"  not found", result.Error.Error(), nil)
		return ctx.Status(404).JSON(failed)
	}

	// Binding user request data to dbData then save
	dbUser.Name = updateUserRequest.Name
	dbUser.Address = updateUserRequest.Address
	dbUser.Email = updateUserRequest.Email
	dbUser.UpdatedAt = time.Time{}

	updateResult := database.DB.Debug().Save(&dbUser)

	if updateResult.Error != nil {
		failed := response.BuildErrorResponse("Error Update User w/ "+uID, result.Error.Error(), nil)
		return ctx.Status(404).JSON(failed)
	}

	log.Println("Success Update " + strconv.FormatInt(updateResult.RowsAffected, 10) + " Data")
	success := response.BuildResponse(true, "Success Update", nil)
	return ctx.Status(200).JSON(success)
}

func DeleteUser(ctx *fiber.Ctx) error {
	uID := ctx.Params("id")
	var dbUser entity.User

	result := database.DB.Debug().First(&dbUser, uID)

	if result.Error != nil {
		failed := response.BuildErrorResponse("User w/ ID"+uID+"  not found", result.Error.Error(), nil)
		return ctx.Status(404).JSON(failed)
	}

	deleteResult := database.DB.Debug().Delete(&dbUser)

	if deleteResult.Error != nil {
		failed := response.BuildErrorResponse("Internal Server Error", result.Error.Error(), nil)
		return ctx.Status(500).JSON(failed)
	}

	log.Println("Success Deleting " + strconv.FormatInt(deleteResult.RowsAffected, 10) + " Data")
	success := response.BuildResponse(true, "Success Delete", nil)
	return ctx.Status(200).JSON(success)
}
