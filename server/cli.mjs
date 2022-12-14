#!/usr/bin/env zx

switch (argv["_"][0]) {
  case "gen":
    await $`go generate ./...`;
    break;
  case "migrate":
    await $`go run main.go migrate`;
    break;
  case "ping":
    await $`go run main.go ping`;
    break;
  case "start":
    await $`go run main.go start`;
    break;
  case "nm":
    await $`go run -mod=mod data/migrate/main.go ${argv["_"][1]}`;
    break;
  case "ip":
    await $`curl https://99wry.cf/qqwry.dat -o ./\pkg/\iploc/\qqwry.dat`;
    break;
  default:
    console.warn("Exit~");
}
