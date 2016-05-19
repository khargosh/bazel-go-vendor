package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_prefix")

go_prefix("github.com/khargosh/bar")

go_library(
  name = "bar",
  srcs = [
    "doc.go",
    "interface.go",
    "circle.go",
    "rectangle.go",
    "square.go",
  ],
  visibility = ["//visibility:public"],
)
