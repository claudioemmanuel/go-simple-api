package repositories

import (
	"devbook-api/pkg/database"
	"devbook-api/pkg/models"
)

// UserRepository is a struct that represents a user repository
type UserRepository struct{}

// FindAll is responsible for getting all users
func (r UserRepository) FindAll() ([]models.User, error) {

	// open a database connection
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	// close the database connection
	defer database.Close(db)

	// get all users
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// FindByID is responsible for getting a user by id
func (r UserRepository) FindByID(id int) (models.User, error) {

	// open a database connection
	db, err := database.Connect()
	if err != nil {
		return models.User{}, err
	}

	// close the database connection
	defer database.Close(db)

	// get the user by id
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

// Create is responsible for creating a new user
func (r UserRepository) Create(user models.User) (models.User, error) {

	// open a database connection
	db, err := database.Connect()
	if err != nil {
		return models.User{}, err
	}

	// close the database connection
	defer database.Close(db)

	// create a new user
	if err := db.Create(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

// Update is responsible for updating a user
func (r UserRepository) Update(id int, user models.User) (models.User, error) {

	// open a database connection
	db, err := database.Connect()
	if err != nil {
		return models.User{}, err
	}

	// close the database connection
	defer database.Close(db)

	// update the user
	db = db.Model(&models.User{}).Where("id = ?", id).Take(&models.User{}).UpdateColumns(
		map[string]interface{}{
			"name":     user.Name,
			"nick":     user.Nick,
			"email":    user.Email,
			"password": user.Password,
		},
	)

	db.First(&user, id)

	return user, nil
}

// Delete is responsible for deleting a user
func (r UserRepository) Delete(id int) error {

	// open a database connection
	db, err := database.Connect()
	if err != nil {
		return err
	}

	// close the database connection
	defer database.Close(db)

	// delete the user
	if err := db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}

	return nil
}
