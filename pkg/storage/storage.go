package storage

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	pFriends "students/pkg/friends"
	pUser "students/pkg/user"
)

type Storage struct {
	DB *gorm.DB
}

func NewStorage() Storage {
	storage := Storage{}

	dsn := "host=skillbox-users-db port=5432 user=admin dbname=api sslmode=disable password=admin"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Cannot connect to postgres database")
		log.Fatal("This is the error:", err)
	} else {
		log.Println("We are connected to the postgres database")
	}

	err = db.Debug().Migrator().DropTable(&pUser.User{}, &pFriends.Friends{})
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&pUser.User{}, &pFriends.Friends{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	storage.DB = db

	return storage
}

func (s Storage) Get(id uint64) (*pUser.User, error) {
	var dto *pUser.User

	err := s.DB.Debug().Model(pUser.User{}).Where("id = ?", id).Take(&dto).Error
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (s Storage) GetFriends(id uint64) ([]*pUser.User, error) {
	user, err := s.Get(id)
	if err != nil {
		return nil, err
	}

	var dtos []*pUser.User

	err = s.DB.Debug().Model(dtos).
		Joins("INNER JOIN friends f ON f.target_id = id").
		Where("f.source_id = ?", user.Id).
		Limit(100).Find(&dtos).Error
	if err != nil {
		return []*pUser.User{}, err
	}

	return dtos, nil
}

func (s Storage) Add(user *pUser.User) (uint64, error) {
	user.Prepare()

	err := s.DB.Debug().Create(&user).Error
	if err != nil {
		return 0, err
	}

	return user.Id, nil
}

func (s Storage) Put(id uint64, user *pUser.User) (bool, error) {
	existUser, err := s.Get(id)
	if err != nil {
		return false, err
	}

	if user.Name != "" {
		existUser.Name = user.Name
	}

	if user.Age != 0 {
		existUser.Age = user.Age
	}

	existUser.Prepare()

	err = s.DB.Debug().Save(&existUser).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s Storage) Attach(suId uint64, tuId uint64) (bool, error) {
	_, err := s.Get(suId)
	if err != nil {
		return false, err
	}

	_, err = s.Get(tuId)
	if err != nil {
		return false, err
	}

	var dto *pFriends.Friends

	err = s.DB.Debug().Model(pFriends.Friends{}).Where("source_id = ? AND target_id = ?", suId, tuId).Find(&dto).Error
	if err != nil {
		return false, err
	}

	if dto == nil || (dto.SourceId == 0 && dto.TargetId == 0) {
		friends := pFriends.NewFriends(suId, tuId)

		err = s.DB.Debug().Create(&friends).Error
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func (s Storage) Delete(id uint64) (bool, error) {
	user, err := s.Get(id)
	if err != nil {
		return false, err
	}

	if err = s.DB.Debug().Delete(&pUser.User{}, &pUser.User{Id: user.Id}).Error; err != nil {
		return false, err
	}

	return true, nil
}
