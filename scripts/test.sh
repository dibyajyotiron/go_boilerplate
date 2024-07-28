#!/bin/sh

# Run tests
go test ./...

# Check if tests passed or failed
if [ $? -ne 0 ]; then
  echo "Tests failed"
  exit 1
else
  echo "Tests passed"
  exit 0
fi
