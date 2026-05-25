# Changelog

## 0.14.0 (2026-05-25)

Full Changelog: [v0.13.0...v0.14.0](https://github.com/CloudCIX/gocloudcix/compare/v0.13.0...v0.14.0)

### Features

* **api:** manual updates ([be3dcd6](https://github.com/CloudCIX/gocloudcix/commit/be3dcd6feb5f6ae60815bfbd3f24fac45775f3ad))

## 0.13.0 (2026-05-25)

Full Changelog: [v0.12.0...v0.13.0](https://github.com/CloudCIX/gocloudcix/compare/v0.12.0...v0.13.0)

### Features

* **api:** manual updates ([b049c22](https://github.com/CloudCIX/gocloudcix/commit/b049c2230c190e92eb8c56deb533fbb44e92631f))

## 0.12.0 (2026-05-25)

Full Changelog: [v0.11.0...v0.12.0](https://github.com/CloudCIX/gocloudcix/compare/v0.11.0...v0.12.0)

### Features

* **api:** manual updates ([9a6efa4](https://github.com/CloudCIX/gocloudcix/commit/9a6efa4f2941540b28a0d4aca9c8ba45c3c7fb71))

## 0.11.0 (2026-05-14)

Full Changelog: [v0.10.1...v0.11.0](https://github.com/CloudCIX/gocloudcix/compare/v0.10.1...v0.11.0)

### Features

* **client:** optimize json encoder for internal types ([bb677c7](https://github.com/CloudCIX/gocloudcix/commit/bb677c78a7fba02f58d4153315f42fb0b2cdced3))

## 0.10.1 (2026-05-13)

Full Changelog: [v0.10.0...v0.10.1](https://github.com/CloudCIX/gocloudcix/compare/v0.10.0...v0.10.1)

## 0.10.0 (2026-05-12)

Full Changelog: [v0.9.0...v0.10.0](https://github.com/CloudCIX/gocloudcix/compare/v0.9.0...v0.10.0)

### Features

* **api:** manual updates ([38410f1](https://github.com/CloudCIX/gocloudcix/commit/38410f136a0ccadc69971f2ea3bed5ffe4e09c8d))

## 0.9.0 (2026-05-12)

Full Changelog: [v0.8.0...v0.9.0](https://github.com/CloudCIX/gocloudcix/compare/v0.8.0...v0.9.0)

### Features

* **api:** manual update to compute 5.0.4 ([e6596ca](https://github.com/CloudCIX/gocloudcix/commit/e6596ca04c78077e6daf76af4f36560e367dc8ab))

## 0.8.0 (2026-05-08)

Full Changelog: [v0.7.0...v0.8.0](https://github.com/CloudCIX/gocloudcix/compare/v0.7.0...v0.8.0)

### Features

* **go:** add default http client with timeout ([5b9a9da](https://github.com/CloudCIX/gocloudcix/commit/5b9a9da9d9be3192a2d9e0bfece3e9e07bfafce6))
* support setting headers via env ([3db2de1](https://github.com/CloudCIX/gocloudcix/commit/3db2de17d6873582c11f9e5b3179aa708336392d))


### Bug Fixes

* **go:** avoid panic when http.DefaultTransport is wrapped ([eba5590](https://github.com/CloudCIX/gocloudcix/commit/eba559036a37e83790bcd3573753b9b569d1012b))


### Chores

* avoid embedding reflect.Type for dead code elimination ([12b896a](https://github.com/CloudCIX/gocloudcix/commit/12b896a24dd2360bef2fc6e6eed3267239028045))
* **internal:** more robust bootstrap script ([73d1799](https://github.com/CloudCIX/gocloudcix/commit/73d17990aa592b87d1f377786681f7155b15cb86))
* redact api-key headers in debug logs ([a143c3a](https://github.com/CloudCIX/gocloudcix/commit/a143c3a50fc95f29b12e67dd01cad7be89ec8c0d))

## 0.7.0 (2026-04-01)

Full Changelog: [v0.6.4...v0.7.0](https://github.com/CloudCIX/gocloudcix/compare/v0.6.4...v0.7.0)

### Features

* **internal:** support comma format in multipart form encoding ([ad102c3](https://github.com/CloudCIX/gocloudcix/commit/ad102c3290153625e5121dc94f24f40318935833))


### Bug Fixes

* fix issue with unmarshaling in some cases ([1a8191d](https://github.com/CloudCIX/gocloudcix/commit/1a8191d0aea5ca8586614a0caa78228c2405dc3e))
* prevent duplicate ? in query params ([06bfce6](https://github.com/CloudCIX/gocloudcix/commit/06bfce6828e49a2b18577816128a1b709d2c23f6))


### Chores

* **ci:** skip lint on metadata-only changes ([ba07b7c](https://github.com/CloudCIX/gocloudcix/commit/ba07b7c58eb4116c88942d344866d407adde9138))
* **ci:** support opting out of skipping builds on metadata-only commits ([3d7ef64](https://github.com/CloudCIX/gocloudcix/commit/3d7ef64b996200282be17f683a4aaa64117d8e52))
* **client:** fix multipart serialisation of Default() fields ([a056ef9](https://github.com/CloudCIX/gocloudcix/commit/a056ef973ce34f3fdb6cdf67a7e5dabc0ec170fb))
* **internal:** support default value struct tag ([825224a](https://github.com/CloudCIX/gocloudcix/commit/825224afeb2e5da6a7597ac57e28b61a1bf5f471))
* **internal:** tweak CI branches ([087ff74](https://github.com/CloudCIX/gocloudcix/commit/087ff74d277d7ed94029f352dfbd444ed3ea7f42))
* **internal:** update gitignore ([3a81868](https://github.com/CloudCIX/gocloudcix/commit/3a81868d4b4be7c6b91a708959798966532a9efc))
* remove unnecessary error check for url parsing ([4c06b50](https://github.com/CloudCIX/gocloudcix/commit/4c06b50afdae0d0f3ee21f7992ccbc2cf3508d60))
* update docs for api:"required" ([238980e](https://github.com/CloudCIX/gocloudcix/commit/238980e9b342b0ac98361a33e09238d13dc6cdea))

## 0.6.4 (2026-03-11)

Full Changelog: [v0.6.3...v0.6.4](https://github.com/CloudCIX/gocloudcix/compare/v0.6.3...v0.6.4)

### Chores

* **internal:** minor cleanup ([f47c031](https://github.com/CloudCIX/gocloudcix/commit/f47c031d4a6368860a4291f2f0cfc6ae7a60f676))
* **internal:** use explicit returns ([2350c4e](https://github.com/CloudCIX/gocloudcix/commit/2350c4e842a52d724c1f3eec795736332f1bbdab))
* **internal:** use explicit returns in more places ([f7c59df](https://github.com/CloudCIX/gocloudcix/commit/f7c59dfe60d5fa159be3f8913b4928041a5173c7))

## 0.6.3 (2026-03-07)

Full Changelog: [v0.6.2...v0.6.3](https://github.com/CloudCIX/gocloudcix/compare/v0.6.2...v0.6.3)

### Chores

* **ci:** skip uploading artifacts on stainless-internal branches ([2f12332](https://github.com/CloudCIX/gocloudcix/commit/2f1233254fe5910b92aa16a4abd81323282948bb))
* **internal:** codegen related update ([db29f5b](https://github.com/CloudCIX/gocloudcix/commit/db29f5b5b1df4d27f1a961c73efc92a94404c4fa))
* **internal:** codegen related update ([ba9ad61](https://github.com/CloudCIX/gocloudcix/commit/ba9ad617485ade9a30c14c005deca8b6ed6f0607))

## 0.6.2 (2026-02-25)

Full Changelog: [v0.6.1...v0.6.2](https://github.com/CloudCIX/gocloudcix/compare/v0.6.1...v0.6.2)

### Bug Fixes

* allow canceling a request while it is waiting to retry ([f655c82](https://github.com/CloudCIX/gocloudcix/commit/f655c828dbcce86d9a8295426311f22a05eb385d))
* **encoder:** correctly serialize NullStruct ([e0f6b90](https://github.com/CloudCIX/gocloudcix/commit/e0f6b906b707959b911c60b73ceea7ea91f4976a))


### Chores

* **internal:** move custom custom `json` tags to `api` ([89fb981](https://github.com/CloudCIX/gocloudcix/commit/89fb9810ddbe05bbf6638d749212c0ad1b30691a))
* **internal:** remove mock server code ([e98926b](https://github.com/CloudCIX/gocloudcix/commit/e98926b1e4d78d1624a8a4791c1ad32224714895))
* update mock server docs ([7ef2f51](https://github.com/CloudCIX/gocloudcix/commit/7ef2f51cba7790f1dccee822bbbb69c678f5d15d))

## 0.6.1 (2026-01-28)

Full Changelog: [v0.6.0...v0.6.1](https://github.com/CloudCIX/gocloudcix/compare/v0.6.0...v0.6.1)

## 0.6.0 (2026-01-24)

Full Changelog: [v0.5.3...v0.6.0](https://github.com/CloudCIX/gocloudcix/compare/v0.5.3...v0.6.0)

### Features

* **client:** add a convenient param.SetJSON helper ([c224e7a](https://github.com/CloudCIX/gocloudcix/commit/c224e7abc723c9f4bdee203fd604ea5a49d7f0e0))

## 0.5.3 (2026-01-17)

Full Changelog: [v0.5.2...v0.5.3](https://github.com/CloudCIX/gocloudcix/compare/v0.5.2...v0.5.3)

### Bug Fixes

* **docs:** add missing pointer prefix to api.md return types ([cd8f4c9](https://github.com/CloudCIX/gocloudcix/commit/cd8f4c9118117f414df6e5141487a43cf8be3f77))


### Chores

* **internal:** update `actions/checkout` version ([d6da822](https://github.com/CloudCIX/gocloudcix/commit/d6da822700176005336f4c1b7e053b8154692cc5))

## 0.5.2 (2026-01-15)

Full Changelog: [v0.5.1...v0.5.2](https://github.com/CloudCIX/gocloudcix/compare/v0.5.1...v0.5.2)

## 0.5.1 (2026-01-15)

Full Changelog: [v0.5.0...v0.5.1](https://github.com/CloudCIX/gocloudcix/compare/v0.5.0...v0.5.1)

## 0.5.0 (2025-12-19)

Full Changelog: [v0.4.0...v0.5.0](https://github.com/CloudCIX/gocloudcix/compare/v0.4.0...v0.5.0)

### Features

* **encoder:** support bracket encoding form-data object members ([59db02b](https://github.com/CloudCIX/gocloudcix/commit/59db02b404635149e2ce7c8b66c56f7821857ce7))


### Bug Fixes

* skip usage tests that don't work with Prism ([37926c8](https://github.com/CloudCIX/gocloudcix/commit/37926c863f95232b3a03e2f109e728a378f3b927))


### Chores

* add float64 to valid types for RegisterFieldValidator ([ded3d6b](https://github.com/CloudCIX/gocloudcix/commit/ded3d6b44b3e00ff6492731405a8fa600d9de47b))

## 0.4.0 (2025-12-10)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/CloudCIX/gocloudcix/compare/v0.3.0...v0.4.0)

### Features

* **api:** manual updates ([88d321a](https://github.com/CloudCIX/gocloudcix/commit/88d321a6a3afdd0598f0969353428f9cb19b9906))

## 0.3.0 (2025-12-08)

Full Changelog: [v0.2.1...v0.3.0](https://github.com/CloudCIX/gocloudcix/compare/v0.2.1...v0.3.0)

### Features

* **api:** manual updates ([0a59dd2](https://github.com/CloudCIX/gocloudcix/commit/0a59dd211a52748280e572f071e13408be7a3d89))

## 0.2.1 (2025-12-06)

Full Changelog: [v0.2.0...v0.2.1](https://github.com/CloudCIX/gocloudcix/compare/v0.2.0...v0.2.1)

### Bug Fixes

* **mcp:** correct code tool API endpoint ([d67d795](https://github.com/CloudCIX/gocloudcix/commit/d67d7958725f995bd533a6905a987a3f29f55cfa))
* rename param to avoid collision ([ef1a84b](https://github.com/CloudCIX/gocloudcix/commit/ef1a84ba78d7d03d62c2a45e7a3b3242c5186a8e))


### Chores

* elide duplicate aliases ([170142d](https://github.com/CloudCIX/gocloudcix/commit/170142de251d5bdfd708921f01572acffe17dff9))
* **internal:** codegen related update ([c4a6413](https://github.com/CloudCIX/gocloudcix/commit/c4a64136a9c7dfc1b8b983b273d459f7059841a2))

## 0.2.0 (2025-12-04)

Full Changelog: [v0.1.1...v0.2.0](https://github.com/CloudCIX/gocloudcix/compare/v0.1.1...v0.2.0)

### Features

* **api:** manual updates ([52e803b](https://github.com/CloudCIX/gocloudcix/commit/52e803b6bc57e00098c83fce45b7efb83cc314a9))

## 0.1.1 (2025-12-04)

Full Changelog: [v0.1.0...v0.1.1](https://github.com/CloudCIX/gocloudcix/compare/v0.1.0...v0.1.1)

## 0.1.0 (2025-12-04)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/CloudCIX/gocloudcix/compare/v0.0.1...v0.1.0)

### Features

* **api:** manual updates ([cdc9731](https://github.com/CloudCIX/gocloudcix/commit/cdc97311a176d056d6e976f6db7cb96c28e92335))


### Chores

* update SDK settings ([34f1a3c](https://github.com/CloudCIX/gocloudcix/commit/34f1a3c7fb53e3fb0c7260978f77d90dafbd254b))
