package models

import (
	"fmt"
	proto "github.com/garden-raccoon/notify-pkg/protocols/notify"
	"github.com/gocql/gocql"
	"sync"
)

type Notification struct {
	UserUUID      string
	VacancyUUID   string
	NoteUUID      string
	CandidateName string
	CandidateUrl  string
	IsReaded      bool
}

func (n Notification) Proto() *proto.Notification {

	pb := &proto.Notification{
		CandidateName: n.CandidateName,
		CanditateUrl:  n.CandidateUrl,
		UserUuid:      n.UserUUID,
		VacancyUuid:   n.VacancyUUID,
		IsReaded:      n.IsReaded,
	}
	return pb
}

func AppliedNotesFronProto(pb *proto.Notifications, noteUuid string) []*Notification {
	var notes []*Notification

	for i := range pb.Notifications {
		notes = append(notes, AppliedFromProto(pb.Notifications[i], noteUuid))
	}
	return notes
}
func AppliedFromProto(pb *proto.Notification, noteUuid string) *Notification {
	return &Notification{
		UserUUID:      pb.UserUuid,
		VacancyUUID:   pb.VacancyUuid,
		CandidateName: pb.CandidateName,
		CandidateUrl:  pb.CanditateUrl,
		IsReaded:      pb.IsReaded,
		NoteUUID:      noteUuid,
	}
}

type MessageNotification struct {
	NoteUUID      gocql.UUID
	UserUUID      gocql.UUID
	VacancyUUID   gocql.UUID
	CandidateName string
	CandidateUrl  string
	IsReaded      bool

	Mu sync.Mutex
}

func NewMessageNotification(candidateName, candidateUrl, userUUID, vacancyUUID string) (*MessageNotification, error) {
	newUUID, err := gocql.RandomUUID()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	userUuid, err := gocql.ParseUUID(userUUID)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	vacUuid, err := gocql.ParseUUID(vacancyUUID)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return &MessageNotification{
		NoteUUID:      newUUID,
		UserUUID:      userUuid,
		VacancyUUID:   vacUuid,
		CandidateName: candidateName,
		CandidateUrl:  candidateUrl,
		IsReaded:      false,
	}, nil
}
