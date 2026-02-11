package service

import (
	"context"
	"encoding/json"

	"automobile-parts-backend/model"
)

type ProductionService struct {
	fabric *FabricService
}

func NewProductionService(fabric *FabricService) *ProductionService {
	return &ProductionService{fabric: fabric}
}

func (s *ProductionService) CreateProductionData(ctx context.Context, dto model.ProductionDataDTO) error {
	payload, err := json.Marshal(dto)
	if err != nil {
		return err
	}
	_, err = s.fabric.Submit(ctx, "CreateProductionData", string(payload))
	return err
}
