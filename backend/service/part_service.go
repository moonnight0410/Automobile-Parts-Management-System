package service

import (
	"context"
	"encoding/json"

	"automobile-parts-backend/model"
)

type PartService struct {
	fabric *FabricService
}

func NewPartService(fabric *FabricService) *PartService {
	return &PartService{fabric: fabric}
}

func (s *PartService) CreatePart(ctx context.Context, dto model.PartDTO) error {
	payload, err := json.Marshal(dto)
	if err != nil {
		return err
	}
	_, err = s.fabric.Submit(ctx, "CreatePart", string(payload))
	return err
}

func (s *PartService) QueryPart(ctx context.Context, partID string) (*model.Part, error) {
	resp, err := s.fabric.Query(ctx, "QueryPart", partID)
	if err != nil {
		return nil, err
	}
	var part model.Part
	if err := json.Unmarshal(resp, &part); err != nil {
		return nil, err
	}
	return &part, nil
}

func (s *PartService) QueryPartByBatchNo(ctx context.Context, batchNo string) ([]model.Part, error) {
	resp, err := s.fabric.Query(ctx, "QueryPartByBatchNo", batchNo)
	if err != nil {
		return nil, err
	}
	var parts []model.Part
	if err := json.Unmarshal(resp, &parts); err != nil {
		return nil, err
	}
	return parts, nil
}

func (s *PartService) QueryPartByVIN(ctx context.Context, vin string) ([]model.Part, error) {
	resp, err := s.fabric.Query(ctx, "QueryPartByVIN", vin)
	if err != nil {
		return nil, err
	}
	var parts []model.Part
	if err := json.Unmarshal(resp, &parts); err != nil {
		return nil, err
	}
	return parts, nil
}
