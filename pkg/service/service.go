package service

import (
	pkg "dcard/pkg"
)

type service struct {
	repo pkg.IRepository
}

// NewService 依賴注入
func NewService(repo pkg.IRepository) pkg.IService {
	return &service{repo: repo}
}
