#!/bin/bash

# ============================================
# Configure sandbox for quick workspace setup.
# ============================================

mkdir -p build
go build -o build/server main.go
