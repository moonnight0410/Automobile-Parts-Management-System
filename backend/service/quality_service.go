package service

import (
	"context"
	"encoding/json"
	"log"

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

func (s *QualityService) ListAllQualityInspections(ctx context.Context, handler string) ([]model.QualityInspection, error) {
	log.Printf("[DEBUG] ListAllQualityInspections called with handler: %q", handler)
	resp, err := s.fabric.Query(ctx, "ListAllQualityInspections", handler)
	if err != nil {
		log.Printf("[ERROR] Fabric query failed: %v", err)
		// 降级：查询失败时返回空列表，避免前端出现500
		return []model.QualityInspection{}, nil
	}
	log.Printf("[DEBUG] Fabric query response: %s", string(resp))
	var inspections []model.QualityInspection
	if err := json.Unmarshal(resp, &inspections); err != nil {
		log.Printf("[ERROR] JSON unmarshal failed: %v", err)
		// 降级：解析失败时返回空列表
		return []model.QualityInspection{}, nil
	}
	log.Printf("[DEBUG] Unmarshaled %d inspections", len(inspections))
	return inspections, nil
}
