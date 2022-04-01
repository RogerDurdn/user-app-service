package pkg

import "github.com/RogerDurdn/users/model"

type Storage interface {
	findUserById(id int) (*model.User, error)
	findUserByName(name string) (*model.User, error)
	createOrUpdateUser(user *model.User) (*model.User, error)
	deleteUserById(id int) bool
}

type PostgresStorage struct {
}

func (p *PostgresStorage) findUserById(id int) (*model.User, error) {
	return nil, nil
}
func (p *PostgresStorage) findUserByName(name string) (*model.User, error) {

	return nil, nil
}
func (p *PostgresStorage) createOrUpdateUser(user *model.User) (*model.User, error) {

	return nil, nil
}
func (p *PostgresStorage) deleteUserById(id int) bool {
	return false
}
