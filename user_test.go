package main

import (
	"github.com/cucumber/godog"
)

func InitializeScenarioUser(ctx *godog.ScenarioContext) {
	testRunner := &TestRunner{}
	ctx.Step(`^Enviarei as seguintes informacoes$`, testRunner.GetPayload)
	ctx.Step(`^Vou enviar "([^"]*)" uma requisicao para "([^"]*)"$`, testRunner.SendRequest)
	ctx.Step(`^Verificar o status code (\d+)$`, testRunner.IsEqual)
	ctx.Step(`^Por fim verificar o corpo da requisicao criada$`, testRunner.IsEqualsJson)

	ctx.Step(`^Vou solicitar uma requisicao "([^"]*)" para "([^"]*)"$`, testRunner.GetRequest)
	ctx.Step(`^Verificar o status code (\d+)$`, testRunner.IsEqual)
	ctx.Step(`^Por fim verificar o corpo da requisicao$`, testRunner.IsEqualsJson)

	ctx.Step(`^Enviarei as seguintes informacoes atualizadas$`, testRunner.GetPayload)
	ctx.Step(`^Vou enviar "([^"]*)" uma requisicao para "([^"]*)"$`, testRunner.SendRequest)
	ctx.Step(`^Verificar o status code (\d+)$`, testRunner.IsEqual)
	ctx.Step(`^Por fim verificar o corpo da requisicao atualizada$`, testRunner.IsEqualsJson)
}
