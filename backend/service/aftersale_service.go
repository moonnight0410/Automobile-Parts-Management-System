package service

import (
	"context"
	"encoding/json"

	"automobile-parts-backend/model"
)

type AftersaleService struct {
	fabric *FabricService
}

func NewAftersaleService(fabric *FabricService) *AftersaleService {
	return &AftersaleService{fabric: fabric}
}

func (s *AftersaleService) CreateFaultReport(ctx context.Context, dto model.FaultReportDTO) error {
	payload, err := json.Marshal(dto)
	if err != nil {
		return err
	}
	_, err = s.fabric.Submit(ctx, "CreateFaultReport", string(payload))
	return err
}

func (s *AftersaleService) CreateRecallRecord(ctx context.Context, dto model.RecallRecordDTO) error {
	payload, err := json.Marshal(dto)
	if err != nil {
		return err
	}
	_, err = s.fabric.Submit(ctx, "CreateRecallRecord", string(payload))
	return err
}

func (s *AftersaleService) CreateAftersaleRecord(ctx context.Context, dto model.AftersaleRecordDTO) error {
	payload, err := json.Marshal(dto)
	if err != nil {
		return err
	}
	_, err = s.fabric.Submit(ctx, "CreateAftersaleRecord", string(payload))
	return err
}
