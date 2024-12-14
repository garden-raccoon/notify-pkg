package models

import (
	proto "github.com/garden-raccoon/notify-pkg/protocols/notify"
)

type Notification struct {
	NoteUUID      string
	UserUUID      string
	VacancyUUID   string
	CandidateName string
	CandidateUrl  string
	IsReaded      bool
}

func (n Notification) Proto() *proto.Notification {

	pb := &proto.Notification{
		NoteUuid:      n.NoteUUID,
		CandidateName: n.CandidateName,
		CanditateUrl:  n.CandidateUrl,
		UserUuid:      n.UserUUID,
		VacancyUuid:   n.VacancyUUID,
		IsReaded:      n.IsReaded,
	}
	return pb
}

func NotesFronProto(pb *proto.Notifications) []*Notification {
	var notes []*Notification

	for i := range pb.Notifications {
		notes = append(notes, FromProto(pb.Notifications[i]))
	}
	return notes
}
func FromProto(pb *proto.Notification) *Notification {
	return &Notification{
		NoteUUID:      pb.NoteUuid,
		UserUUID:      pb.UserUuid,
		VacancyUUID:   pb.VacancyUuid,
		CandidateName: pb.CandidateName,
		CandidateUrl:  pb.CanditateUrl,
		IsReaded:      pb.IsReaded,
	}
}
