#!make

# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
#                                                                             #
#      ____                               _                                   #
#     / __ \____  ___  ____  ____ _____  (_) ®                                #
#    / / / / __ \/ _ \/ __ \/ __ `/ __ \/ /                                   #
#   / /_/ / /_/ /  __/ / / / /_/ / /_/ / /                                    #
#   \____/ .___/\___/_/ /_/\__,_/ .___/_/                                     #
#       /_/                    /_/                                            #
#                                                                             #
#   The Largest Certified API Marketplace                                     #
#   Accelerate Digital Transformation • Simplify Processes • Lead Industry    #
#                                                                             #
#   ═══════════════════════════════════════════════════════════════════════   #
#                                                                             #
#   Project:        openapi-go-sdk                                            #
#   Version:        0.1.0                                                     #
#   Author:         L. Paderi (@lpaderiAltravia)                              #
#   Copyright:      (c) 2025 Openapi®. All rights reserved.                   #
#   License:        MIT                                                       #
#   Maintainer:     Francesco Bianco                                          #
#   Contact:        https://openapi.com/                                      #
#   Repository:     https://github.com/openapi/openapi-go-sdk/               #
#   Documentation:  https://console.openapi.com/                              #
#                                                                             #
#   ═══════════════════════════════════════════════════════════════════════   #
#                                                                             #
#   "Truth lies at the source of the stream."                                 #
#                                  — English Proverb                          #
#                                                                             #
# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #

## =========
## Variables
## =========

VERSION := 1.2.1
TAG     := v$(VERSION)

## ====================
## Development Commands
## ====================

test:
	@go test ./...

vet:
	@go vet ./...

dev-push:
	@git config credential.helper 'cache --timeout=3600'
	@git add .
	@git commit -m "$$(read -p 'Commit message: ' msg; echo $$msg)" || true
	@git push

## ================
## Release Commands
## ================

push:
	@git add .
	@git commit -am "Updated at $$(date)" || true
	@git push

release:
	@echo "==> Releasing $(TAG)..."
	@if [ "$$(git rev-parse --abbrev-ref HEAD)" != "main" ]; then \
		echo "ERROR: releases must be cut from the main branch"; exit 1; \
	fi
	@if [ -n "$$(git status --porcelain)" ]; then \
		echo "ERROR: working directory is not clean"; exit 1; \
	fi
	@echo "==> Running tests..."
	@go test ./...
	@echo "==> Running vet..."
	@go vet ./...
	@echo "==> Tagging $(TAG)..."
	@git tag -a "$(TAG)" -m "Release $(TAG)"
	@git push origin "$(TAG)"
	@echo "==> Creating GitHub release..."
	@gh release create "$(TAG)" \
		--title "$(TAG)" \
		--generate-notes \
		--verify-tag
	@echo "==> Done. Release $(TAG) is live."