#!/bin/sh
goctl api go -api test.api -dir . --style go_zero
#goctl api plugin -p goctl-rest-discover="rest-discover " -api exa.api -dir .