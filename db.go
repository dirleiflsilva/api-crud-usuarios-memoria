package main

import "github.com/google/uuid"

// FindAll retorna a lista completa de usuários
func (app *Application) FindAll() []User {
	users := make([]User, 0, len(app.Data))
	for _, user := range app.Data {
		users = append(users, user)
	}
	return users
}

// FindById retorna o usuário correspondente ao id ou nil se não existir
func (app *Application) FindById(id uuid.UUID) *User {
	user, exists := app.Data[id]
	if !exists {
		return nil
	}
	return &user
}

// Insert adiciona um novo usuário e retorna o usuário recém-criado com seu id
func (app *Application) Insert(user User) User {
	user.ID = uuid.New()
	app.Data[user.ID] = user
	return user
}

// Update atualiza um usuário existente e retorna a versão atualizada
func (app *Application) Update(id uuid.UUID, userUpdates User) *User {
	if _, exists := app.Data[id]; !exists {
		return nil
	}
	userUpdates.ID = id
	app.Data[id] = userUpdates
	return &userUpdates
}

// Delete remove um usuário e retorna o usuário que foi deletado
func (app *Application) Delete(id uuid.UUID) *User {
	user, exists := app.Data[id]
	if !exists {
		return nil
	}
	delete(app.Data, id)
	return &user
}
