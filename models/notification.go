package models

import (
	"fmt"
	proto "github.com/garden-raccoon/notify-pkg/protocols/notify"
	"github.com/gocql/gocql"
	"sync"
)

type Notification struct {
	EmployeeUUID  string
	VacancyUUID   string
	EmployerUUID  string
	NoteUUID      string
	ResumeUUID    string
	CandidateName string
	CandidateUrl  string
	IsRead        bool
}

func (n Notification) Proto() *proto.Notification {

	pb := &proto.Notification{
		NoteUuid:      n.NoteUUID,
		EmployerUuid:  n.EmployerUUID,
		CandidateName: n.CandidateName,
		CanditateUrl:  n.CandidateUrl,
		EmployeeUuid:  n.EmployeeUUID,
		VacancyUuid:   n.VacancyUUID,
		ResumeUuid:    n.ResumeUUID,
		IsRead:        n.IsRead,
	}
	return pb
}

func AppliedNotesFromProto(pb *proto.Notifications, employerUuid string) []*Notification {
	var notes []*Notification

	for i := range pb.Notifications {
		notes = append(notes, AppliedFromProto(pb.Notifications[i], employerUuid))
	}
	return notes
}
func AppliedFromProto(pb *proto.Notification, noteUuid string) *Notification {
	return &Notification{
		EmployeeUUID:  pb.EmployeeUuid,
		VacancyUUID:   pb.VacancyUuid,
		EmployerUUID:  pb.EmployerUuid,
		CandidateName: pb.CandidateName,
		CandidateUrl:  pb.CanditateUrl,
		ResumeUUID:    pb.ResumeUuid,
		IsRead:        pb.IsRead,
		NoteUUID:      noteUuid,
	}
}

type MessageNotification struct {
	NoteUUID      gocql.UUID
	EmployeeUUID  gocql.UUID
	EmployerUUID  gocql.UUID
	VacancyUUID   gocql.UUID
	ResumeUUID    gocql.UUID
	CandidateName string
	CandidateUrl  string
	IsRead        bool

	Mu sync.Mutex
}

func NewMessageNotification(candidateName, candidateUrl, employeeUUID, employerUUID, vacancyUUID, resumeUUID string) (*MessageNotification, error) {
	newUUID, err := gocql.RandomUUID()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	employee, err := gocql.ParseUUID(employeeUUID)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	employer, err := gocql.ParseUUID(employerUUID)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	vacUuid, err := gocql.ParseUUID(vacancyUUID)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	resume, err := gocql.ParseUUID(resumeUUID)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return &MessageNotification{
		NoteUUID:      newUUID,
		EmployeeUUID:  employee,
		EmployerUUID:  employer,
		VacancyUUID:   vacUuid,
		CandidateName: candidateName,
		CandidateUrl:  candidateUrl,
		ResumeUUID:    resume,
		IsRead:        false,
	}, nil
}

func UpdateReadNoty(noteUUID string) (*MessageNotification, error) {
	noteUuid, err := gocql.ParseUUID(noteUUID)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return &MessageNotification{
		NoteUUID: noteUuid,
		IsRead:   true,
	}, nil
}
