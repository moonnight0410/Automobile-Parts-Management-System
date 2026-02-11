package service

import (
	"context"
	"encoding/json"

	"automobile-parts-backend/model"
)

type BOMService struct {
	fabric *FabricService
}

func NewBOMService(fabric *FabricService) *BOMService {
	return &BOMService{fabric: fabric}
}

func (s *BOMService) CreateBOM(ctx context.Context, dto model.BOMDTO) error {
	payload, err := json.Marshal(dto)
	if err != nil {
		return err
	}
	_, err = s.fabric.Submit(ctx, "CreateBOM", string(payload))
	return err
}

func (s *BOMService) QueryBOM(ctx context.Context, bomID string) (*model.BOM, error) {
	resp, err := s.fabric.Query(ctx, "QueryBOM", bomID)
	if err != nil {
		return nil, err
	}
	var bom model.BOM
	if err := json.Unmarshal(resp, &bom); err != nil {
		return nil, err
	}
	return &bom, nil
}
