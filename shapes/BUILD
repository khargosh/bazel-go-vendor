package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_prefix")

go_prefix("github.com/khargosh/shapes")

go_binary(
  name = "shapes",
  srcs = ["shapes.go"],
  visibility = ["//visibility:public"],
  deps = ["//vendor/github.com/khargosh/bar:go_default_library"],
)
