package main

import (
	"github.com/cucumber/godog"
)

func InitializeScenarioComments(ctx *godog.ScenarioContext) {
	testRunner := &TestRunner{}
	ctx.Step(`^I send "([^"]*)" request to "([^"]*)"$`, testRunner.SendRequest)
	ctx.Step(`^the response code should be (\d+)$`, testRunner.IsEqual)
	ctx.Step(`^the response should match json:$`, testRunner.IsEqualsJson)
}
