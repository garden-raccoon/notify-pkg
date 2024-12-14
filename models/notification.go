package models

import (
	proto "github.com/garden-raccoon/notify-pkg/protocols/notify"
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
