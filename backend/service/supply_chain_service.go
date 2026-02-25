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

func (s *SupplyChainService) ListAllSupplyOrders(ctx context.Context, buyer string, seller string) ([]model.SupplyOrder, error) {
	resp, err := s.fabric.Query(ctx, "ListAllSupplyOrders", buyer, seller)
	if err != nil {
		return nil, err
	}
	var orders []model.SupplyOrder
	if err := json.Unmarshal(resp, &orders); err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *SupplyChainService) ListAllLogisticsData(ctx context.Context) ([]model.LogisticsData, error) {
	resp, err := s.fabric.Query(ctx, "ListAllLogisticsData")
	if err != nil {
		return nil, err
	}
	var logisticsList []model.LogisticsData
	if err := json.Unmarshal(resp, &logisticsList); err != nil {
		return nil, err
	}
	return logisticsList, nil
}
