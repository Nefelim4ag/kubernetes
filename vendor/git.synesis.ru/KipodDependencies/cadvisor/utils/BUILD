load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "path.go",
        "timed_store.go",
        "utils.go",
    ],
    importpath = "git.synesis.ru/KipodDependencies/cadvisor/utils",
    visibility = ["//visibility:public"],
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
        "//vendor/git.synesis.ru/KipodDependencies/cadvisor/utils/cloudinfo:all-srcs",
        "//vendor/git.synesis.ru/KipodDependencies/cadvisor/utils/cpuload:all-srcs",
        "//vendor/git.synesis.ru/KipodDependencies/cadvisor/utils/docker:all-srcs",
        "//vendor/git.synesis.ru/KipodDependencies/cadvisor/utils/oomparser:all-srcs",
        "//vendor/git.synesis.ru/KipodDependencies/cadvisor/utils/sysfs:all-srcs",
        "//vendor/git.synesis.ru/KipodDependencies/cadvisor/utils/sysinfo:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
