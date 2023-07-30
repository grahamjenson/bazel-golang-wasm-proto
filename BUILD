load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")
load("@com_github_bazelbuild_buildtools//buildifier:def.bzl", "buildifier")

# gazelle:prefix github.com/grahamjenson/bazel-golang-wasm-proto
gazelle(name = "gazelle")

buildifier(name = "buildifier")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/grahamjenson/bazel-golang-wasm-proto",
    visibility = ["//visibility:private"],
    deps = [
        "//protos/api:go_default_library",
        "//server:go_default_library",
        "@com_github_maxence_charriere_go_app_v6//pkg/app:go_default_library",
    ],
)

go_binary(
    name = "server",
    data = [
        "//wasm:app.wasm",
        "@com_github_bootstrap//file:bootstrap.css",
        "@com_github_ec2instances//file:instances.json",
    ],
    args = [
        "--bootstrap-css-path=$(location @com_github_bootstrap//file:bootstrap.css)",
        "--wasm-path=$(location //wasm:app.wasm)",
    ],
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
