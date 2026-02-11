package service

import (
	"context"
	"encoding/json"

	"automobile-parts-backend/model"
)

type SupplyChainService struct {
	fabric *FabricService
}

func NewSupplyChainService(fabric *FabricService) *SupplyChainService {
	return &SupplyChainService{fabric: fabric}
}

func (s *SupplyChainService) CreateSupplyOrder(ctx context.Context, dto model.SupplyOrderDTO) error {
	payload, err := json.Marshal(dto)
	if err != nil {
		return err
	}
	_, err = s.fabric.Submit(ctx, "CreateSupplyOrder", string(payload))
	return err
}

func (s *SupplyChainService) CreateLogisticsData(ctx context.Context, dto model.LogisticsDataDTO) error {
	payload, err := json.Marshal(dto)
	if err != nil {
		return err
	}
	_, err = s.fabric.Submit(ctx, "CreateLogisticsData", string(payload))
	return err
}
