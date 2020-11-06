package service

import (
	"dcard/internal/testsuite"
	"dcard/pkg"
	"dcard/pkg/repository"
	"os"
	"testing"

	"go.uber.org/fx"
)

type serviceSuite struct {
	repo pkg.IRepository
}

var suite serviceSuite

func TestMain(m *testing.M) {
	var exit int
	defer os.Exit(exit)
	defer testsuite.Close()
	testsuite.Initialize(
		fx.Provide(
			repository.NewRepository,
		),
		fx.Populate(&suite.repo),
	)
	exit = m.Run()
}
