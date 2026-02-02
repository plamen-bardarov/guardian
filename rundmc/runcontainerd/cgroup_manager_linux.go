package runcontainerd

import (
	"path/filepath"

	gardencgroups "code.cloudfoundry.org/guardian/rundmc/cgroups"
	"code.cloudfoundry.org/guardian/rundmc/goci"
	"github.com/opencontainers/cgroups"
	"github.com/opencontainers/cgroups/fs2"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

func (m cgroupManager) setUnifiedResources(bundle goci.Bndl) error {
	if bundle.Spec.Linux.CgroupsPath != "" && bundle.Spec.Annotations["container-type"] == "garden-init" && cgroups.IsCgroup2UnifiedMode() {
		// In cgroups v2 we move init process to "init" child cgroup
		// and set resources manually on parent cgroup
		newCgroupPath := filepath.Join(bundle.Spec.Linux.CgroupsPath, gardencgroups.InitCgroupName)

		// we are using UnifiedMountpoint because fs2.CreateCgroupPath checks that path starts with it
		cgroupPath := filepath.Join(fs2.UnifiedMountpoint, bundle.Spec.Linux.CgroupsPath)

		resources := convertSpecResourcesToCgroupResources(bundle.Spec.Linux.Resources)
		if resources != nil {
			cgroupManager, err := fs2.NewManager(&cgroups.Cgroup{}, cgroupPath)
			if err != nil {
				return err
			}
			err = fs2.CreateCgroupPath(cgroupPath, &cgroups.Cgroup{})
			if err != nil {
				return err
			}
			err = cgroupManager.Set(resources)
			if err != nil {
				return err
			}
		}

		bundle.Spec.Linux.CgroupsPath = newCgroupPath
	}

	return nil
}

func (m cgroupManager) mountUnifiedCgroup(bundle *goci.Bndl) {
	if bundle.Spec.Linux.CgroupsPath != "" && bundle.Spec.Annotations["container-type"] == "garden-init" && cgroups.IsCgroup2UnifiedMode() {
		// Mount the cgroup v2 hierarchy inside the container
		// at /sys/fs/cgroup so that processes inside the container
		// can see their cgroup settings. This is not done in "rundmc/runcontainerd/runcontainerd.go" because
		// we want to mount the "init" cgroup specifically instead of the whole hierarchy.
		bundle.Spec.Mounts = append(bundle.Spec.Mounts, specs.Mount{
			Destination: "/sys/fs/cgroup",
			Type:        "bind",
			Source:      filepath.Join("/sys/fs/cgroup", bundle.CGroupPath()),
			Options:     []string{"bind", "ro", "nosuid", "noexec", "nodev"},
		})
	}
}

func convertSpecResourcesToCgroupResources(specResources *specs.LinuxResources) *cgroups.Resources {
	if specResources == nil {
		return nil
	}

	resources := &cgroups.Resources{}
	resources.Unified = specResources.Unified

	if specResources.CPU != nil {
		if specResources.CPU.Shares != nil {
			resources.CpuShares = *specResources.CPU.Shares
		}
		if specResources.CPU.Quota != nil {
			resources.CpuQuota = *specResources.CPU.Quota
		}
		if specResources.CPU.Period != nil {
			resources.CpuPeriod = *specResources.CPU.Period
		}
	}
	if specResources.Memory != nil {
		if specResources.Memory.Limit != nil {
			resources.Memory = *specResources.Memory.Limit
		}
		if specResources.Memory.Swap != nil {
			resources.MemorySwap = *specResources.Memory.Swap
		}
	}

	return resources
}
