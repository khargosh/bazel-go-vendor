# Load Golang repository and build definitions.
git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    commit = "8b2c913a841c7ea4c3f0514fbc8f132f5dd03dfd",
)

load('@io_bazel_rules_go//go:def.bzl', 'go_repositories')
go_repositories()


# Load Closure repository and build definitions.

git_repository(
    name = "io_bazel_rules_closure",
    remote = "https://github.com/bazelbuild/rules_closure.git",
    tag = "0.1.0",
)

load("@io_bazel_rules_closure//closure:defs.bzl", "closure_repositories")
closure_repositories()
