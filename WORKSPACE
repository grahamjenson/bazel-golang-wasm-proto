workspace(name = "bazel_golang_wasm_proto")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

###
# Rules
###

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "51dc53293afe317d2696d4d6433a4c33feedb7748a9e352072e2ec3c0dafd2c6",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.40.1/rules_go-v0.40.1.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.40.1/rules_go-v0.40.1.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "727f3e4edd96ea20c29e8c2ca9e8d2af724d8c7778e7923a854b2c80952bc405",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.30.0/bazel-gazelle-v0.30.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.30.0/bazel-gazelle-v0.30.0.tar.gz",
    ],
)

git_repository(
    name = "com_google_protobuf",
    commit = "21edc3b26fc80771baea78d8bd8df1285b12cae2",
    remote = "https://github.com/protocolbuffers/protobuf",
)

http_archive(
    name = "rules_proto_grpc",
    sha256 = "5f0f2fc0199810c65a2de148a52ba0aff14d631d4e8202f41aff6a9d590a471b",
    strip_prefix = "rules_proto_grpc-1.0.2",
    urls = ["https://github.com/rules-proto-grpc/rules_proto_grpc/archive/1.0.2.tar.gz"],
)

###
# Overrides to get WASM working with protoc
###

http_archive(
    name = "com_github_gogo_protobuf",
    patch_args = ["-p1"],
    patches = [
        "//third_party:gogo.patch",
    ],
    sha256 = "2056a39c922c7315530fc5b7a6ce10cc83b58c844388c9b2e903a0d8867a8b66",
    strip_prefix = "protobuf-1.3.1",
    # v1.3.1, latest as of 2020-01-03
    urls = [
        "https://mirror.bazel.build/github.com/gogo/protobuf/archive/v1.3.1.zip",
        "https://github.com/gogo/protobuf/archive/v1.3.1.zip",
    ],
)

###
# Bootrap
###

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(go_version = "1.20.5")

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

# gazelle:repository_macro third_party/go_repositories.bzl%go_repositories
load("//third_party:go_repositories.bzl", "go_repositories")

go_repositories()

###
# Protobuf
###

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

###
# GRPC
###

load("@rules_proto_grpc//:repositories.bzl", "rules_proto_grpc_repos", "rules_proto_grpc_toolchains")

rules_proto_grpc_toolchains()

rules_proto_grpc_repos()

load("@com_github_grpc_grpc//bazel:grpc_deps.bzl", "grpc_deps")

grpc_deps()

load("@com_github_grpc_grpc//bazel:grpc_extra_deps.bzl", "grpc_extra_deps")

grpc_extra_deps()

###
# Data Files
###

http_file(
    name = "com_github_bootstrap",
    downloaded_file_path = "bootstrap.css",
    sha256 = "038ecec312ff9c0374c9d8831534865fb7ed6df4c94ca822274cea0ae4cf0e1e",
    urls = ["https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.css"],
)

http_file(
    name = "com_github_ec2instances",
    downloaded_file_path = "instances.json",
    sha256 = "8cf2c06b485cfef6567a1554589b4e3ce4ad8e61116a5edf32ed6233010b0fba",
    urls = ["https://raw.githubusercontent.com/powdahound/ec2instances.info/b6664cf095405e806d69ea2c8b1d3f02b5951cf1/www/instances.json"],
)
