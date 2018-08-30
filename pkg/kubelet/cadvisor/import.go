// +build mock_godep

package cadvisor

import (
	_ "git.synesis.ru/KipodDependencies/cadvisor/zfs"
	_ "git.synesis.ru/KipodDependencies/cadvisor/container"
	_ "git.synesis.ru/KipodDependencies/cadvisor/container/common"
	_ "git.synesis.ru/KipodDependencies/cadvisor/container/containerd"
	_ "git.synesis.ru/KipodDependencies/cadvisor/container/systemd"
	_ "git.synesis.ru/KipodDependencies/cadvisor/container/crio"
	_ "git.synesis.ru/KipodDependencies/cadvisor/container/docker"
	_ "git.synesis.ru/KipodDependencies/cadvisor/container/libcontainer"
	_ "git.synesis.ru/KipodDependencies/cadvisor/container/rkt"
	_ "git.synesis.ru/KipodDependencies/cadvisor/container/raw"
	_ "git.synesis.ru/KipodDependencies/cadvisor/machine"
	_ "git.synesis.ru/KipodDependencies/cadvisor/devicemapper"
	_ "git.synesis.ru/KipodDependencies/cadvisor/events"
	_ "git.synesis.ru/KipodDependencies/cadvisor/version"
	_ "git.synesis.ru/KipodDependencies/cadvisor/summary"
	_ "git.synesis.ru/KipodDependencies/cadvisor/client/v2"
	_ "git.synesis.ru/KipodDependencies/cadvisor/info/v1"
	_ "git.synesis.ru/KipodDependencies/cadvisor/info/v2"
	_ "git.synesis.ru/KipodDependencies/cadvisor/healthz"
	_ "git.synesis.ru/KipodDependencies/cadvisor/manager"
	_ "git.synesis.ru/KipodDependencies/cadvisor/manager/watcher"
	_ "git.synesis.ru/KipodDependencies/cadvisor/manager/watcher/rkt"
	_ "git.synesis.ru/KipodDependencies/cadvisor/manager/watcher/raw"
	_ "git.synesis.ru/KipodDependencies/cadvisor/storage"
	_ "git.synesis.ru/KipodDependencies/cadvisor/pages"
	_ "git.synesis.ru/KipodDependencies/cadvisor/pages/static"
	_ "git.synesis.ru/KipodDependencies/cadvisor/fs"
	_ "git.synesis.ru/KipodDependencies/cadvisor/validate"
	_ "git.synesis.ru/KipodDependencies/cadvisor/utils"
	_ "git.synesis.ru/KipodDependencies/cadvisor/utils/sysinfo"
	_ "git.synesis.ru/KipodDependencies/cadvisor/utils/docker"
	_ "git.synesis.ru/KipodDependencies/cadvisor/utils/cpuload"
	_ "git.synesis.ru/KipodDependencies/cadvisor/utils/cpuload/netlink"
	_ "git.synesis.ru/KipodDependencies/cadvisor/utils/sysfs"
	_ "git.synesis.ru/KipodDependencies/cadvisor/utils/oomparser"
	_ "git.synesis.ru/KipodDependencies/cadvisor/utils/cloudinfo"
	_ "git.synesis.ru/KipodDependencies/cadvisor/api"
	_ "git.synesis.ru/KipodDependencies/cadvisor/collector"
	_ "git.synesis.ru/KipodDependencies/cadvisor/cache/memory"
	_ "git.synesis.ru/KipodDependencies/cadvisor/http"
	_ "git.synesis.ru/KipodDependencies/cadvisor/http/mux"
	_ "git.synesis.ru/KipodDependencies/cadvisor/metrics"
	_ "git.synesis.ru/KipodDependencies/cadvisor/accelerators"
)
