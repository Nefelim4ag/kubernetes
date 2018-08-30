load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "container.go",
        "factory.go",
    ],
    importpath = "git.synesis.ru/KipodDependencies/cadvisor/container",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/google/cadvisor/info/v1:go_default_library",
        "//vendor/github.com/google/cadvisor/manager/watcher:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//vendor/git.synesis.ru/KipodDependencies/cadvisor/container/common:all-srcs",
        "//vendor/git.synesis.ru/KipodDependencies/cadvisor/container/containerd:all-srcs",
        "//vendor/git.synesis.ru/KipodDependencies/cadvisor/container/crio:all-srcs",
        "//vendor/git.synesis.ru/KipodDependencies/cadvisor/container/docker:all-srcs",
        "//vendor/git.synesis.ru/KipodDependencies/cadvisor/container/libcontainer:all-srcs",
        "//vendor/git.synesis.ru/KipodDependencies/cadvisor/container/raw:all-srcs",
        "//vendor/git.synesis.ru/KipodDependencies/cadvisor/container/rkt:all-srcs",
        "//vendor/git.synesis.ru/KipodDependencies/cadvisor/container/systemd:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
