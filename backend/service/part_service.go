package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

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
		errMsg := err.Error()
		if strings.Contains(errMsg, "不存在") || strings.Contains(errMsg, "not found") {
			return nil, fmt.Errorf("零部件不存在")
		}
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

func (s *PartService) ListAllParts(ctx context.Context, manufacturer string) ([]model.Part, error) {
	resp, err := s.fabric.Query(ctx, "ListAllParts", manufacturer)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return []model.Part{}, nil
	}
	var parts []model.Part
	if err := json.Unmarshal(resp, &parts); err != nil {
		return nil, err
	}
	return parts, nil
}

func (s *PartService) DeletePart(ctx context.Context, partID string) error {
	_, err := s.fabric.Submit(ctx, "DeletePart", partID)
	return err
}

func (s *PartService) QueryPartLifecycle(ctx context.Context, partID string) (*model.PartLifecycle, error) {
	resp, err := s.fabric.Query(ctx, "QueryPartLifecycle", partID)
	if err != nil {
		errMsg := err.Error()
		if errMsg != "" && (strings.Contains(errMsg, "chaincode response 500") ||
			strings.Contains(errMsg, "lifecycle data not found") ||
			strings.Contains(errMsg, "lifecycle does not exist") ||
			strings.Contains(errMsg, "生命周期数据不存在")) {
			return &model.PartLifecycle{}, nil
		}
		return nil, err
	}
	if len(resp) == 0 {
		return &model.PartLifecycle{}, nil
	}

	respStr := string(resp)
	if strings.Contains(respStr, "生命周期数据不存在") ||
		strings.Contains(respStr, "lifecycle data not found") ||
		strings.Contains(respStr, "does not exist") {
		return &model.PartLifecycle{}, nil
	}

	var lifecycle model.PartLifecycle
	if err := json.Unmarshal(resp, &lifecycle); err != nil {
		return &model.PartLifecycle{}, nil
	}
	if lifecycle.AftersaleInfo == nil {
		lifecycle.AftersaleInfo = []model.AftersaleRecord{}
	}
	if lifecycle.SupplyChainInfo == nil {
		lifecycle.SupplyChainInfo = []model.SupplyChainData{}
	}
	return &lifecycle, nil
}
