variable "GO_VERSION" {
  default = null
}

variable "DESTDIR" {
  default = "./bin"
}

group "default" {
  targets = ["build"]
}

target "build" {
  args = {
    GO_VERSION = "${GO_VERSION}"
  }
}

group "test" {
  targets = ["test-root", "test-noroot"]
}

target "test-root" {
  inherits = ["build"]
  target = "test-coverage"
  output = ["${DESTDIR}/coverage"]
}

target "test-noroot" {
  inherits = ["build"]
  target = "test-noroot-coverage"
  output = ["${DESTDIR}/coverage"]
}

target "lint" {
  dockerfile = "./hack/dockerfiles/lint.Dockerfile"
  output = ["type=cacheonly"]
  args = {
    GO_VERSION = "${GO_VERSION}"
  }
}

target "validate-generated-files" {
  dockerfile = "./hack/dockerfiles/generated-files.Dockerfile"
  output = ["type=cacheonly"]
  target = "validate"
  args = {
    GO_VERSION = "${GO_VERSION}"
  }
}

target "generated-files" {
  inherits = ["validate-generated-files"]
  output = ["."]
  target = "update"
}

target "validate-gomod" {
  dockerfile = "./hack/dockerfiles/gomod.Dockerfile"
  output = ["type=cacheonly"]
  target = "validate"
  args = {
    # go mod may produce different results between go versions,
    # if this becomes a problem, this should be switched to use
    # a fixed go version.
    GO_VERSION = "${GO_VERSION}"
  }
}

target "gomod" {
  inherits = ["validate-gomod"]
  output = ["."]
  target = "update"
}

target "validate-shfmt" {
  dockerfile = "./hack/dockerfiles/shfmt.Dockerfile"
  output = ["type=cacheonly"]
  target = "validate"
}

target "shfmt" {
  inherits = ["validate-shfmt"]
  output = ["."]
  target = "update"
}

target "cross" {
  inherits = ["build"]
  platforms = ["linux/amd64", "linux/386", "linux/arm64", "linux/arm", "linux/ppc64le", "linux/s390x", "darwin/amd64", "darwin/arm64", "windows/amd64", "windows/arm64", "freebsd/amd64", "freebsd/arm64"]
}