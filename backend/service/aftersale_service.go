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

func (s *AftersaleService) ListAllFaultReports(ctx context.Context, reporter string) ([]model.FaultReport, error) {
	resp, err := s.fabric.Query(ctx, "ListAllFaultReports", reporter)
	if err != nil {
		return nil, err
	}
	var faults []model.FaultReport
	if err := json.Unmarshal(resp, &faults); err != nil {
		return nil, err
	}
	return faults, nil
}

func (s *AftersaleService) ListAllRecallRecords(ctx context.Context) ([]model.RecallRecord, error) {
	resp, err := s.fabric.Query(ctx, "ListAllRecallRecords")
	if err != nil {
		return nil, err
	}
	var recalls []model.RecallRecord
	if err := json.Unmarshal(resp, &recalls); err != nil {
		return nil, err
	}
	return recalls, nil
}

func (s *AftersaleService) ListAllAftersaleRecords(ctx context.Context, handlerOrgID string) ([]model.AftersaleRecord, error) {
	resp, err := s.fabric.Query(ctx, "ListAllAftersaleRecords", handlerOrgID)
	if err != nil {
		return nil, err
	}
	var records []model.AftersaleRecord
	if err := json.Unmarshal(resp, &records); err != nil {
		return nil, err
	}
	return records, nil
}
