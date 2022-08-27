#!/usr/bin/env zx

switch (argv["_"][0]) {
  case "gen":
    await $`go generate ./...`;
    break;
  case "migrate":
    await $`go run main.go migrate`;
    break;
  case "nm":
    await $`go run -mod=mod data/migrate/main.go ${argv["_"][1]}`;
    break;
  default:
    console.warn("Exit~");
}
