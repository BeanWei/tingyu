#!/usr/bin/env zx

switch (argv["_"][0]) {
  case "gen":
    await $`go generate ./...`;
    break;
  default:
    console.warn("Exit~");
}
