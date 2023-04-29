package friends

import (
	pUser "students/pkg/user"
)

type Friends struct {
	TargetId uint64     `gorm:"not null"`
	SourceId uint64     `gorm:"not null;"`
	Source   pUser.User `gorm:"foreignKey:SourceId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Target   pUser.User `gorm:"foreignKey:TargetId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func NewFriends(sourceId uint64, targetId uint64) *Friends {
	return &Friends{

		SourceId: sourceId,
		TargetId: targetId,
	}
}
