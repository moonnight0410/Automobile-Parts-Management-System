package service

import (
	"context"
	"encoding/json"
	"log"

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

func (s *ProductionService) ListAllProductionData(ctx context.Context, operator string) ([]model.ProductionData, error) {
	log.Printf("[DEBUG] ListAllProductionData called with operator: %q", operator)
	resp, err := s.fabric.Query(ctx, "ListAllProductionData", operator)
	if err != nil {
		log.Printf("[ERROR] Fabric query failed: %v", err)
		return nil, err
	}
	log.Printf("[DEBUG] Fabric response: %s", string(resp))
	var productionList []model.ProductionData
	if err := json.Unmarshal(resp, &productionList); err != nil {
		log.Printf("[ERROR] JSON unmarshal failed: %v", err)
		return nil, err
	}
	log.Printf("[DEBUG] Unmarshaled %d production records", len(productionList))
	return productionList, nil
}
