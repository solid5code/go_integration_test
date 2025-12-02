package main

import (
	"testing"

	// 引入 allure 核心定義

	// 引入 runner 執行器
	"github.com/ozontech/allure-go/pkg/framework/runner"
	// 引入 provider，這是 Test/Step 介面定義的來源
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func TestRunner(t *testing.T) {
	runner.Run(t, "My first test", func(t provider.T) {
		t.NewStep("My First Step!")
	})
	runner.Run(t, "My second test", func(t provider.T) {
		t.WithNewStep("My Second Step!", func(sCtx provider.StepCtx) {
			sCtx.NewStep("My First SubStep!")
		})
	})
}
