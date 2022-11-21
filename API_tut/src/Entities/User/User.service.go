package User

// =============================================

// Get User by id if exist
func getUserById(id int) User {
	// idx := slices.IndexFunc(Users,
	// 	func(user User) bool { return user.Id == id })
	// return Users[idx]
}

// =============================================

func userGetListService() []User {

	// Users := Database.ORM.Find(&User{})
	// return Users
}

// =============================================

func userGetOneService(id int) User {
	return getUserById(id)
}

// =============================================

func userCreateService(user User) User {
	return user
}

// =============================================

func userUpdateService(id int, user User) User {
	return getUserById(id)
}

// =============================================

func userDeleteService(id int) User {
	return getUserById(id)
}

// =============================================
