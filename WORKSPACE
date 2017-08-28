workspace(name = "cryptdo")

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.5.4",
)

load("@io_bazel_rules_go//go:def.bzl", "gazelle", "go_repositories", "go_repository")
go_repositories()

load("@io_bazel_rules_go//proto:go_proto_library.bzl", "go_proto_repositories")
go_proto_repositories()

go_repository(
    name = "com_github_golang_protobuf",
    commit = "ab9f9a6dab164b7d1246e0e688b0ab7b94d8553e",
    importpath = "github.com/golang/protobuf",
)

go_repository(
    name = "com_github_jessevdk_go_flags",
    commit = "6cf8f02b4ae8ba723ddc64dcfd403e530c06d927",
    importpath = "github.com/jessevdk/go-flags",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "81e90905daefcd6fd217b62423c0908922eadb30",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "2d6f6f883a06fc0d5f4b14a81e4c28705ea64c15",
    importpath = "golang.org/x/sys",
)
