# Changelog

## [0.26.3](https://github.com/blaxel-ai/sdk-go/compare/v0.26.2...v0.26.3) (2026-08-13)


### Bug Fixes

* **openapi:** keep internal fields out of the public API reference ([4252c7c](https://github.com/blaxel-ai/sdk-go/commit/4252c7c3ec912aaf5fd1973ac676048f67ca17df))

## [0.26.2](https://github.com/blaxel-ai/sdk-go/compare/v0.26.1...v0.26.2) (2026-08-12)


### Bug Fixes

* **controlplane:** resolve volume and sandbox node type based on mk3.1 capabilities ([47bfd4e](https://github.com/blaxel-ai/sdk-go/commit/47bfd4ee7467a192b402a52d3dcc584caca2ae31))

## [0.26.1](https://github.com/blaxel-ai/sdk-go/compare/v0.26.0...v0.26.1) (2026-08-05)


### Chores

* **deps:** bump github.com/modelcontextprotocol/go-sdk to v1.4.1 ([6d7c1cf](https://github.com/blaxel-ai/sdk-go/commit/6d7c1cf109dbb150792a960a91045b62c47d57da))
* **deps:** fix critical/high Dependabot alerts in sdk-go ([2aa9ced](https://github.com/blaxel-ai/sdk-go/commit/2aa9ced50d26cda22d7d1f6687c3df29b4cdd38a))

## [0.26.0](https://github.com/blaxel-ai/sdk-go/compare/v0.25.1...v0.26.0) (2026-07-31)


### Features

* Add agent drive management functionality ([0068eb5](https://github.com/blaxel-ai/sdk-go/commit/0068eb5067f5b7ba90326aa215a4cf0442b8bbe7))
* add configurable terminated sandbox retention with quota support ([1965913](https://github.com/blaxel-ai/sdk-go/commit/19659137fca3381ac953466da5e4b56562b29922))
* add ecr for model-gateway ([1ecf9bb](https://github.com/blaxel-ai/sdk-go/commit/1ecf9bb0d7b8f5ea7359c6e9220c096de6754aa6))
* Add mk2 to RuntimeGeneration enum in ModelRuntime ([f6d98c6](https://github.com/blaxel-ai/sdk-go/commit/f6d98c605e1e8c874af918d5e70a900c36d4ffca))
* add POST /v0/images endpoint for image-only builds (ENG-2192) ([bf8fb9b](https://github.com/blaxel-ai/sdk-go/commit/bf8fb9b3447a818fc64db4aca80f9b741070bb1c))
* Add reference field to FunctionSpec for MCP code mode with base64 normalization ([1bf7385](https://github.com/blaxel-ai/sdk-go/commit/1bf73854206fbf184bdc98903005c1880e3c1c72))
* Add region and volume to agent ([ff80c1e](https://github.com/blaxel-ai/sdk-go/commit/ff80c1ec7e30952c9ae30c78d0acf73634fa4e8d))
* add typed gateway error fields to Error struct ([#32](https://github.com/blaxel-ai/sdk-go/issues/32)) ([5532564](https://github.com/blaxel-ai/sdk-go/commit/5532564f95c9006457cc300c35a31925f505b919))
* Adds public access control to deployments ([aa5fc3a](https://github.com/blaxel-ai/sdk-go/commit/aa5fc3a1260287bb261017fdd5f187f685a663e6))
* allow UpdateInstanceTTL and UpdateInstanceLifecycle to clear values (ENG-3189) ([#43](https://github.com/blaxel-ai/sdk-go/issues/43)) ([f999f04](https://github.com/blaxel-ai/sdk-go/commit/f999f04a4902a5761535b6fb1bdda2433c7bb885))
* **api:** api update ([290fba3](https://github.com/blaxel-ai/sdk-go/commit/290fba3ca90a6f5abfacfb3abdfb33a872a4466f))
* **api:** api update ([0c56e1e](https://github.com/blaxel-ai/sdk-go/commit/0c56e1e0af9bff967376b3eabdad2c08cf932e6c))
* **api:** api update ([5d6b16d](https://github.com/blaxel-ai/sdk-go/commit/5d6b16d01b1a34d0b3e1d89fd752b977951054a3))
* **api:** api update ([7652208](https://github.com/blaxel-ai/sdk-go/commit/7652208a8f0d539308932ea6cfc906636fe80ff4))
* **api:** api update ([b54c20e](https://github.com/blaxel-ai/sdk-go/commit/b54c20e0db9bc39dfc310a0e0757fbf835812ffa))
* **api:** api update ([9977c13](https://github.com/blaxel-ai/sdk-go/commit/9977c13074e962803649afa349e9f65058fa4a4d))
* **api:** api update ([254c90d](https://github.com/blaxel-ai/sdk-go/commit/254c90dc013adf7c6f34e4dcdc0b3cdef72750d4))
* **api:** manual updates ([6a9addc](https://github.com/blaxel-ai/sdk-go/commit/6a9addc85f46dd1d15df70c4f207a8e754bdbb41))
* **api:** manual updates ([a0d22a7](https://github.com/blaxel-ai/sdk-go/commit/a0d22a74dc310fc8de9495135d9b5384ce16c77c))
* **api:** manual updates ([99e6b9c](https://github.com/blaxel-ai/sdk-go/commit/99e6b9c8df567111b3e178edccc88c18a73adb1e))
* **api:** manual updates ([86f8b5f](https://github.com/blaxel-ai/sdk-go/commit/86f8b5fafae3420341f4d94c2813d14fcbcfe46f))
* **api:** manual updates ([eb7ca5d](https://github.com/blaxel-ai/sdk-go/commit/eb7ca5dace7e6af0f5a9858d2a3988fe41b68aa9))
* **api:** manual updates ([b9cc972](https://github.com/blaxel-ai/sdk-go/commit/b9cc97207710a094783504cf25a0f99758e18596))
* **api:** manual updates ([1034c4d](https://github.com/blaxel-ai/sdk-go/commit/1034c4d84dfceb75a47a57880ceff84129140efd))
* **api:** manual updates ([f906521](https://github.com/blaxel-ai/sdk-go/commit/f90652178894dd97dc8d3aa470a44d00ea1da375))
* **api:** manual updates ([60a8e76](https://github.com/blaxel-ai/sdk-go/commit/60a8e764c492d0a0637feb5006c9bc3d242b7bcf))
* **api:** manual updates ([1d756f4](https://github.com/blaxel-ai/sdk-go/commit/1d756f4687ae5727c2a84df6f867004040039060))
* **api:** manual updates ([7559443](https://github.com/blaxel-ai/sdk-go/commit/75594437a9afdc580776b2bb17857f1854f1e96e))
* **api:** manual updates ([737f42c](https://github.com/blaxel-ai/sdk-go/commit/737f42cca1c9cbc63161c1dd3d8fc26546a8f68b))
* **api:** manual updates ([5199115](https://github.com/blaxel-ai/sdk-go/commit/51991156e1f223d50280b5a180175e0e95140184))
* **api:** manual updates ([4bb88fa](https://github.com/blaxel-ai/sdk-go/commit/4bb88fa2830d7bed9c0b408ce90752b31b9abbec))
* **api:** manual updates ([320213c](https://github.com/blaxel-ai/sdk-go/commit/320213cd3336e8af02131392fa2089d000c5df28))
* **api:** manual updates ([9f2a6c9](https://github.com/blaxel-ai/sdk-go/commit/9f2a6c9bb1e641bc7e7e1d9dbf12784c6dcb9737))
* **api:** manual updates ([2d5fa98](https://github.com/blaxel-ai/sdk-go/commit/2d5fa98b706846658831450a9db77bc47a45efec))
* **api:** manual updates ([3f72f51](https://github.com/blaxel-ai/sdk-go/commit/3f72f51f8893f1365468fb206f09fba126ed1cd0))
* **api:** manual updates ([9e05d75](https://github.com/blaxel-ai/sdk-go/commit/9e05d75e577ada0a1b7e93dd1a695c375a1bbcd3))
* **api:** manual updates ([1b31a12](https://github.com/blaxel-ai/sdk-go/commit/1b31a1265138a9654470275b876188a96085b973))
* **api:** manual updates ([1561602](https://github.com/blaxel-ai/sdk-go/commit/1561602eed3ad6992bf22c70f3c063fbccbb5a08))
* **api:** manual updates ([453f0e2](https://github.com/blaxel-ai/sdk-go/commit/453f0e28691568b11d331122dca9e68a57cf2aa6))
* **api:** manual updates ([a35ad91](https://github.com/blaxel-ai/sdk-go/commit/a35ad910d1d4ec2e8f15b79954541dea47ba84c7))
* **api:** manual updates ([3a7959d](https://github.com/blaxel-ai/sdk-go/commit/3a7959deb5514a2fa35762670f442d45cd803a43))
* Cdrappier/increase job speed ([32e7d60](https://github.com/blaxel-ai/sdk-go/commit/32e7d60512a628975ab5d319da190a5b962310e4))
* Cdrappier/state in dynamo k8s ([0101e2a](https://github.com/blaxel-ai/sdk-go/commit/0101e2ab4a20746e80bcc3ed74efe2c802891fd2))
* **client:** add a convenient param.SetJSON helper ([b8bf4c8](https://github.com/blaxel-ai/sdk-go/commit/b8bf4c87904265e387197373eac73ea05fe21635))
* **config:** add environment to workspace configuration ([2f7ad7d](https://github.com/blaxel-ai/sdk-go/commit/2f7ad7d068939ee9b07a959fd9cb578d78ff7a8a))
* **config:** add SetCurrentWorkspace function to update the current workspace in the configuration ([2f7ad7d](https://github.com/blaxel-ai/sdk-go/commit/2f7ad7d068939ee9b07a959fd9cb578d78ff7a8a))
* **config:** implement AuthHeaders method for dynamic authentication header generation ([2f7ad7d](https://github.com/blaxel-ai/sdk-go/commit/2f7ad7d068939ee9b07a959fd9cb578d78ff7a8a))
* Enhances authentication with SSO, Directory Sync, and domain policies ([9b46033](https://github.com/blaxel-ai/sdk-go/commit/9b46033043e4002754efe1ec098b5da123846b8d))
* expose gateway error codes as typed fields on apierror.Error ([#33](https://github.com/blaxel-ai/sdk-go/issues/33)) ([0caa620](https://github.com/blaxel-ai/sdk-go/commit/0caa6207e3200fc6c35c05aab5fc862fb24261f5))
* GitHub Actions Self-Hosted Runner Integration (ENG-2198) ([ead4c2e](https://github.com/blaxel-ai/sdk-go/commit/ead4c2e1ee621a916a09a22701d70e9e9c289588))
* **internal:** support comma format in multipart form encoding ([e01b695](https://github.com/blaxel-ai/sdk-go/commit/e01b695c01f70e330f1689de968612e5a59b0691))
* Majoffre/dedicated ips ([5128af3](https://github.com/blaxel-ai/sdk-go/commit/5128af3793043dd907b9b7829c426d01ecf3d84d))
* Majoffre/drive browser ([3e5c96f](https://github.com/blaxel-ai/sdk-go/commit/3e5c96f00ad2595ac8c4f78a212e39fd237fcf74))
* Majoffre/support proxy config ([243b2d5](https://github.com/blaxel-ai/sdk-go/commit/243b2d5f7100b724b41d7a83983b9a2c51877ced))
* Remove seaweedfs from doc generated ([6e00399](https://github.com/blaxel-ai/sdk-go/commit/6e00399a3ed932d75fe8e7abd60067b2fb9671b6))
* Sandbox deletionin field ([a88cb22](https://github.com/blaxel-ai/sdk-go/commit/a88cb22907686eb35f769bec8820a16a645fcaec))
* Thomas/add filter billing metrics on resource info ([f9724db](https://github.com/blaxel-ai/sdk-go/commit/f9724db573e2117e44352e91164b029e8c86bea5))
* Update pre-commit configuration to include models directory and adjust file types; modify import order in sandbox.go for consistency. ([f475e5a](https://github.com/blaxel-ai/sdk-go/commit/f475e5aac3c967f889ce9fda21c6501dda90c32e))
* Update stainless openapi definition ([cfa7186](https://github.com/blaxel-ai/sdk-go/commit/cfa7186e44c335d6ffd9fba62180837f3af413ac))


### Bug Fixes

* add retry loop for URL checks and wait-after-delete in PreviewRaceConditions test ([7c9c671](https://github.com/blaxel-ai/sdk-go/commit/7c9c67172fa9fb7518beb04cdcbb5e7d2a1ccdf5))
* allow canceling a request while it is waiting to retry ([d3aff84](https://github.com/blaxel-ai/sdk-go/commit/d3aff84df378e10fc07caa28dece88a71b883f93))
* **api:** remove egressIpName from SandboxNetwork descriptions [ENG-1970] (cherry-pick to main) ([ac84637](https://github.com/blaxel-ai/sdk-go/commit/ac84637b2deae41c38e7ced98773823c0a0ef515))
* **client/oauth:** send grant_type in the right location ([ab15867](https://github.com/blaxel-ai/sdk-go/commit/ab15867ad420f3ca5e03c607c4dd13690f8106dd))
* **docs:** add missing pointer prefix to api.md return types ([2bfb9da](https://github.com/blaxel-ai/sdk-go/commit/2bfb9da8db85e48d1be51e3c0a37191cfe8695ae))
* **encoder:** correctly serialize NullStruct ([e1df07c](https://github.com/blaxel-ai/sdk-go/commit/e1df07cd49ffae1bd1ec306c497f79ab8897fac9))
* fix issue with unmarshaling in some cases ([ba8612d](https://github.com/blaxel-ai/sdk-go/commit/ba8612df12ddf516cdeb2bda88c22893f1b309a7))
* fix request delays for retrying to be more respectful of high requested delays ([33842ba](https://github.com/blaxel-ai/sdk-go/commit/33842baa419e0fc0505a2c49f9cd92a2936b252b))
* prevent duplicate ? in query params ([f1f5eef](https://github.com/blaxel-ai/sdk-go/commit/f1f5eefdee01bba41b25605802b0cbe737ce5c03))
* **stlc:** unblock all three SDK targets ([0b72118](https://github.com/blaxel-ai/sdk-go/commit/0b72118569bacc7da5b9be42f0834b80d23dda4f))
* **tracking:** default tracking to false and check config.yaml ([#24](https://github.com/blaxel-ai/sdk-go/issues/24)) ([4dc4c42](https://github.com/blaxel-ai/sdk-go/commit/4dc4c42bf32012e8bb6a85f9dd354b825b0cca38))


### Reverts

* restore custom code dropped by the one-repo build ([45a1ba0](https://github.com/blaxel-ai/sdk-go/commit/45a1ba06111a99004e8f01ca2e113fca5daa55a6))


### Chores

* **ci:** add build step ([78735fc](https://github.com/blaxel-ai/sdk-go/commit/78735fc0413ece828e7290d564992d80e62431ee))
* **ci:** skip lint on metadata-only changes ([b56cc28](https://github.com/blaxel-ai/sdk-go/commit/b56cc28c645cd9399c76d91ba0b40b9ec8e21dc1))
* **ci:** skip uploading artifacts on stainless-internal branches ([0e74d4a](https://github.com/blaxel-ai/sdk-go/commit/0e74d4a758a711b331ae685a9436d35316afa6db))
* **ci:** support opting out of skipping builds on metadata-only commits ([992e67f](https://github.com/blaxel-ai/sdk-go/commit/992e67f97a6e17f14fbeb281fa61be81531d4b7c))
* **client:** fix multipart serialisation of Default() fields ([8d79964](https://github.com/blaxel-ai/sdk-go/commit/8d79964837bf709396c468d167a67657dfe30d53))
* **deps:** bump github.com/modelcontextprotocol/go-sdk to v1.4.1 ([6bffcd3](https://github.com/blaxel-ai/sdk-go/commit/6bffcd38e45f4a8bb0289a59266ebb6ea27ed7a8))
* **deps:** bump golang.org/x/sys to v0.44.0 (GO-2026-5024) ([b9912a4](https://github.com/blaxel-ai/sdk-go/commit/b9912a4413d2fec6038957226048c359ae10cfa0))
* **internal:** codegen related update ([e7461c9](https://github.com/blaxel-ai/sdk-go/commit/e7461c90989c5b1a1c5c70b28975ee8cd6e3e6aa))
* **internal:** minor cleanup ([78c99fc](https://github.com/blaxel-ai/sdk-go/commit/78c99fc066add92b6d538bed6b03857a9dc81f71))
* **internal:** move custom custom `json` tags to `api` ([8f3951c](https://github.com/blaxel-ai/sdk-go/commit/8f3951cdc0a7f52ba2e6b213c65197ef905bee44))
* **internal:** support default value struct tag ([fdc1409](https://github.com/blaxel-ai/sdk-go/commit/fdc14090651b64f5992c8f55eb5b19196bd658c8))
* **internal:** tweak CI branches ([2624941](https://github.com/blaxel-ai/sdk-go/commit/26249417f8fcc52174268866d42aaa03279d6b1f))
* **internal:** update `actions/checkout` version ([965b6f3](https://github.com/blaxel-ai/sdk-go/commit/965b6f36b1aaff091ebcd1503da798ba29042647))
* **internal:** update gitignore ([3c32bd3](https://github.com/blaxel-ai/sdk-go/commit/3c32bd3de7c45de038e8502c2ea0a5cec39676e3))
* **internal:** use explicit returns ([65de7a6](https://github.com/blaxel-ai/sdk-go/commit/65de7a657e0b66bb721250a847db4c36fb898da0))
* **internal:** use explicit returns in more places ([9c6239f](https://github.com/blaxel-ai/sdk-go/commit/9c6239f2b4d1d149592380dcabb7d8a2755df854))
* remove dead two-repo promote/sync workflows ([3a7cf4a](https://github.com/blaxel-ai/sdk-go/commit/3a7cf4ac5f9a7e6d37268a41211811794d4d7d9a))
* remove dead two-repo promote/sync workflows (again) ([5437e30](https://github.com/blaxel-ai/sdk-go/commit/5437e309345325c4be81bd3b68c598cf160f37da))
* remove unnecessary error check for url parsing ([965fae5](https://github.com/blaxel-ai/sdk-go/commit/965fae5ca669fca09fc9712959e436affdc3cf34))
* **stlc:** switch every target to a one-repo setup ([8a4924e](https://github.com/blaxel-ai/sdk-go/commit/8a4924ef4e9d7d42e4cb25400a962e2e73f6250c))
* **stlc:** switch the Go SDK to a one-repo setup ([8cbd085](https://github.com/blaxel-ai/sdk-go/commit/8cbd085ad11cdede495e77ceb6e231d15142be15))
* **test:** do not count install time for mock server timeout ([f4e0f45](https://github.com/blaxel-ai/sdk-go/commit/f4e0f450d010cc8480e7e1b26aa7c439d489db9d))
* **tests:** bump steady to v0.19.4 ([2b33e36](https://github.com/blaxel-ai/sdk-go/commit/2b33e3654aa1813dbcc97f316b086f6f7e2ff5cb))
* **tests:** bump steady to v0.19.5 ([1e09115](https://github.com/blaxel-ai/sdk-go/commit/1e091159c837d22085e3f42b39aa495f55aed1e6))
* **tests:** bump steady to v0.19.6 ([10ebe39](https://github.com/blaxel-ai/sdk-go/commit/10ebe3930c7b03ea238c8ccc25cc3c1a4e9a5220))
* **tests:** bump steady to v0.19.7 ([85cf24a](https://github.com/blaxel-ai/sdk-go/commit/85cf24abe824ef63942bc3238759587f17ee5251))
* **tests:** bump steady to v0.20.1 ([a28d558](https://github.com/blaxel-ai/sdk-go/commit/a28d558f078d44d5c3187bfccdb4370356b1c7b1))
* **tests:** bump steady to v0.20.2 ([cc04fd6](https://github.com/blaxel-ai/sdk-go/commit/cc04fd6c20f337fb6decc9fa4bfdc4b1006dc6ae))
* update docs for api:"required" ([f753bbe](https://github.com/blaxel-ai/sdk-go/commit/f753bbef8845aeb4e8c5b81249294b705c7fe2ff))
* update mock server docs ([e4ff897](https://github.com/blaxel-ai/sdk-go/commit/e4ff89787282be3a2fef652c4067f0bf411809e3))
* update placeholder string ([ba6fd7e](https://github.com/blaxel-ai/sdk-go/commit/ba6fd7e2cdc850f848f9da471d6ecb4e66354116))
* update SDK settings ([28c4a0f](https://github.com/blaxel-ai/sdk-go/commit/28c4a0f7518a9b0b138a4d665017cb9df280491a))


### Documentation

* add data collection section in README ([bfb5a08](https://github.com/blaxel-ai/sdk-go/commit/bfb5a08736fa14b1423243c3e7e8576ddc7c6abd))
* add Prerequisites section with BL_WORKSPACE and BL_API_KEY (ENG-2261) ([0305fee](https://github.com/blaxel-ai/sdk-go/commit/0305feeb7e7986b651385461f4e5e5f6f4950425))
* add Prerequisites section with required env vars (ENG-2261) ([f760565](https://github.com/blaxel-ai/sdk-go/commit/f760565422e053e7b48fe75be6a3e481fca77723))
* move Prerequisites clear of the x-release-please regions ([b0031ed](https://github.com/blaxel-ai/sdk-go/commit/b0031ed494ba0abfd86279cec65586561e094b25))
* move Prerequisites clear of the x-release-please regions ([ebad9d3](https://github.com/blaxel-ai/sdk-go/commit/ebad9d37a0dc9a1a4f5e703add2de4213b60248c))
* update README ([9112f73](https://github.com/blaxel-ai/sdk-go/commit/9112f730a239baf934f57469ab3c9303ac3a89fc))
* update README for data collection ([6e3eb14](https://github.com/blaxel-ai/sdk-go/commit/6e3eb14a8513db45b9426c877576d441b8266006))


### Refactors

* **config:** rename writeConfig to WriteConfig for consistency and improved readability ([2f7ad7d](https://github.com/blaxel-ai/sdk-go/commit/2f7ad7d068939ee9b07a959fd9cb578d78ff7a8a))
* **tests:** switch from prism to steady ([47207a8](https://github.com/blaxel-ai/sdk-go/commit/47207a8b743b34daf31caf3e48b9ea51b8be89b7))
