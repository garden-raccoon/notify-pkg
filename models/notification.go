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
	CandidateName string
	CandidateUrl  string
	IsReaded      bool
}

func (n Notification) Proto() *proto.Notification {

	pb := &proto.Notification{
		NoteUuid:      n.NoteUUID,
		EmployerUuid:  n.EmployerUUID,
		CandidateName: n.CandidateName,
		CanditateUrl:  n.CandidateUrl,
		EmployeeUuid:  n.EmployeeUUID,
		VacancyUuid:   n.VacancyUUID,
		IsReaded:      n.IsReaded,
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
		IsReaded:      pb.IsReaded,
		NoteUUID:      noteUuid,
	}
}

type MessageNotification struct {
	NoteUUID      gocql.UUID
	EmployeeUUID  gocql.UUID
	EmployerUUID  gocql.UUID // is recieved while update !
	VacancyUUID   gocql.UUID
	CandidateName string
	CandidateUrl  string
	IsReaded      bool

	Mu sync.Mutex
}

func NewMessageNotification(candidateName, candidateUrl, employeeUUID, employerUUID, vacancyUUID string) (*MessageNotification, error) {
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
	return &MessageNotification{
		NoteUUID:      newUUID,
		EmployeeUUID:  employee,
		EmployerUUID:  employer,
		VacancyUUID:   vacUuid,
		CandidateName: candidateName,
		CandidateUrl:  candidateUrl,
		IsReaded:      false,
	}, nil
}
