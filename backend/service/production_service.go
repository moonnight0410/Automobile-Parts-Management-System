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
	log.Printf("[Production] 开始创建生产数据: %+v", dto)
	
	payload, err := json.Marshal(dto)
	if err != nil {
		log.Printf("[Production] JSON序列化失败: %v", err)
		return err
	}
	
	log.Printf("[Production] 调用Fabric链码: CreateProductionData")
	result, err := s.fabric.Submit(ctx, "CreateProductionData", string(payload))
	if err != nil {
		log.Printf("[Production] Fabric提交失败: %v", err)
		return err
	}
	
	log.Printf("[Production] Fabric响应: %s", string(result))
	return nil
}

func (s *ProductionService) ListAllProductionData(ctx context.Context, operator string) ([]model.ProductionData, error) {
	log.Printf("[Production] 查询生产数据, operator: %q", operator)
	resp, err := s.fabric.Query(ctx, "ListAllProductionData", operator)
	if err != nil {
		log.Printf("[Production] Fabric查询失败: %v", err)
		return nil, err
	}
	log.Printf("[Production] Fabric响应: %s", string(resp))
	var productionList []model.ProductionData
	if err := json.Unmarshal(resp, &productionList); err != nil {
		log.Printf("[Production] JSON反序列化失败: %v", err)
		return nil, err
	}
	log.Printf("[Production] 查询到 %d 条生产数据", len(productionList))
	return productionList, nil
}
