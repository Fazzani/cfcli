package service

import (
	. "github.com/fazzani/cfcli/domain"
)

type IRecordService interface {
	FindAll(zone_id string) ([]Record, []error)
	Create(zone_id string, r Record) (bool, []error)
	Delete(zone_id string, r Record) (bool, []error)
	Update(zone_id string, r Record) (bool, []error)
}
type HttpRecordService struct{}
