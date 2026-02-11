package service

import (
	"context"
	"encoding/json"

	"automobile-parts-backend/model"
)

type QualityService struct {
	fabric *FabricService
}

func NewQualityService(fabric *FabricService) *QualityService {
	return &QualityService{fabric: fabric}
}

func (s *QualityService) CreateQualityInspection(ctx context.Context, dto model.QualityInspectionDTO) error {
	payload, err := json.Marshal(dto)
	if err != nil {
		return err
	}
	_, err = s.fabric.Submit(ctx, "CreateQualityInspection", string(payload))
	return err
}
