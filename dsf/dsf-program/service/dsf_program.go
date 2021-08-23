package service

import (
	"log"
	"middleware-mmksi/dsf/dsf-program/repo"
)

type DsfProgramService interface {
	GetAdditionalInsurance() (*repo.AdditionalResponse, error)
}

type dsfProgramService struct {
	dsfProgramRepo repo.DsfProgramRepo
}

func NewDsfProgramService(
	dsfProgramRepo repo.DsfProgramRepo,
) DsfProgramService {
	return &dsfProgramService{
		dsfProgramRepo: dsfProgramRepo,
	}
}

func (s *dsfProgramService) GetAdditionalInsurance() (*repo.AdditionalResponse, error) {

	result, err := s.dsfProgramRepo.GetAdditionalInsurance()
	if err != nil {
		return nil, err
	}

	log.Print("res servis", result)
	return result, nil
}
