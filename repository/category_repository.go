package repository

//layer repository sebagai jembatan ke database

import "belajar-golang-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
