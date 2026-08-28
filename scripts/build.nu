#!/usr/bin/env nu

def main [] {
  mkdir target
  go build -o target ./cmd/tenpo/*.go
}
