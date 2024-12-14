package notify

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"

	"github.com/misnaged/scriptorium/logger"

	"github.com/garden-raccoon/notify-pkg/models"
)

type INotificator interface {
	GetAllAppliedCandidatesByNoty(notyUuid string) ([]*models.Notification, error)
	//HealthCheck() error

	NewRegisterNotification(req *models.MessageNotification)
	NewUpdateNotification(req *models.MessageNotification)

	Stop()
	// Close GRPC Api connection
}

type notificator struct {
	errHandler chan *string
	conn       *kafka.Conn
	mq         *models.MessageNotification
	stop       chan struct{}

	address        string
	api            *Api
	registerWriter *kafka.Writer
	updaterWriter  *kafka.Writer

	timeout time.Duration
}

func (noty *notificator) NewRegisterNotification(req *models.MessageNotification) {
	defer req.Mu.Unlock()
	req.Mu.Lock()

	recordJSON, err := json.Marshal(req)
	if err != nil {
		logger.Log().Error(fmt.Errorf("failed to marshal message: %w", err).Error())
		return
	}

	if err := noty.registerWriter.WriteMessages(context.Background(),
		kafka.Message{
			Value: recordJSON,
		}); err != nil {
		logger.Log().Error(fmt.Errorf("failed to write message: %w", err).Error())
		return
	}
}
func (noty *notificator) NewKafkaWriter(topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:       kafka.TCP(noty.address),
		Topic:      topic,
		Balancer:   &kafka.LeastBytes{},
		BatchBytes: 104857600,
	}
}

func (noty *notificator) NewKafkaConn(transport, address string, part int) (*kafka.Conn, error) {
	var err error
	noty.conn, err = kafka.DialLeader(context.Background(), transport, address, "register", part)
	if err != nil {
		return nil, fmt.Errorf("couldn't run kafka dialeader: %w", err)
	}
	logger.Log().Info("kafka client started")
	return noty.conn, nil
}
func (noty *notificator) Stop() {
	close(noty.errHandler)
}

func (noty *notificator) Check() bool {
	if noty.conn == nil {
		return false
	}
	_, _, err := noty.conn.ReadOffsets()
	if err != nil {
		logger.Log().Error(err.Error())
		return false
	}
	return true

}

//func (noty *notificator) discoverLoop(cfg *config.Scheme) {
//	var err error
//
//	for {
//		select {
//		case <-noty.stop:
//			return
//		default:
//		Checking:
//			if !noty.Check() {
//				logger.Log().Error("connection to kafka server has been lost. reconnecting...")
//				sleep(cfg.Kafka.RetryTimeout, noty.stop)
//				noty.conn, err = kafka.DialLeader(context.Background(), cfg.Kafka.Transport, cfg.Kafka.Address, cfg.Kafka.Topic, cfg.Kafka.Partition)
//				if err != nil {
//					goto Checking
//				}
//			}
//
//		}
//	}
//}

func (noty *notificator) NewUpdateNotification(req *models.MessageNotification) {
	defer req.Mu.Unlock()
	req.Mu.Lock()

	recordJSON, err := json.Marshal(req)
	if err != nil {
		logger.Log().Error(fmt.Errorf("failed to marshal message: %w", err).Error())
		return
	}

	if err := noty.updaterWriter.WriteMessages(context.Background(),
		kafka.Message{
			Value: recordJSON,
		}); err != nil {
		logger.Log().Error(fmt.Errorf("failed to write message: %w", err).Error())
		return
	}
}
