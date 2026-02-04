//go:build !windows

package gqt_test

import (
	"bufio"
	"path/filepath"
	"strings"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/guardian/gqt/cgrouper"
	"code.cloudfoundry.org/guardian/gqt/runner"
	gardencgroups "code.cloudfoundry.org/guardian/rundmc/cgroups"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("cgroup mounts inside container", func() {
	var (
		client     *runner.RunningGarden
		container  garden.Container
		cgroupType string
		limits     garden.Limits
		privileged bool
	)

	BeforeEach(func() {
		privileged = false

		limits = garden.Limits{
			Memory: garden.MemoryLimits{
				LimitInBytes: 256 * 1024 * 1024, // 256MB
			},
		}
	})

	JustBeforeEach(func() {
		client = runner.Start(config)

		var err error
		container, err = client.Create(garden.ContainerSpec{
			Limits:     limits,
			Privileged: privileged,
		})
		Expect(err).NotTo(HaveOccurred())

		// Verify cgroup path exists
		parentPath, err := cgrouper.GetCGroupPath(client.CgroupsRootPath(), cgroupType, config.Tag, privileged, cpuThrottlingEnabled())
		Expect(err).NotTo(HaveOccurred())
		cgroupPath := filepath.Join(parentPath, container.Handle())
		Expect(cgroupPath).To(BeADirectory())
	})

	AfterEach(func() {
		Expect(client.DestroyAndStop()).To(Succeed())
	})

	When("hybrid cgroup is enabled", func() {
		BeforeEach(func() {
			if gardencgroups.IsCgroup2UnifiedMode() {
				Skip("Skipping hybrid cgroups tests when cgroups v2 is enabled")
			}

			cgroupType = "memory"
		})

		It("mounts v1 controllers in the container", func() {
			// Read mountinfo from inside the container
			mountinfo := runInContainerCombinedOutput(container, "cat", []string{"/proc/self/mountinfo"})

			// Parse mount entries to find cgroup v1 controllers
			var foundMemoryCgroup bool
			scanner := bufio.NewScanner(strings.NewReader(mountinfo))
			for scanner.Scan() {
				line := scanner.Text()
				// Check for memory controller mount
				if strings.Contains(line, "/sys/fs/cgroup/memory") && strings.Contains(line, "- cgroup") {
					foundMemoryCgroup = true
				}
			}

			Expect(foundMemoryCgroup).To(BeTrue(), "Expected cgroup v1 memory controller to be mounted at /sys/fs/cgroup/memory")
		})
	})

	When("unified cgroup(v2) is enabled", func() {
		BeforeEach(func() {
			if !gardencgroups.IsCgroup2UnifiedMode() {
				Skip("Skipping unified cgroups(v2) tests when hybrid cgroups is enabled")
			}
		})

		It("mounts unified root in the container", func() {
			// Read mountinfo from inside the container
			mountinfo := runInContainerCombinedOutput(container, "cat", []string{"/proc/self/mountinfo"})

			// Parse mount entries to find cgroup v2 unified mount
			// Format: mountID parentID major:minor root mountpoint options - fstype source superOptions
			// Example: 1317 1310 0:29 / /sys/fs/cgroup ro,nosuid,nodev,noexec,relatime - cgroup2 cgroup rw,nsdelegate,memory_recursiveprot
			var foundCgroup2 bool
			scanner := bufio.NewScanner(strings.NewReader(mountinfo))
			for scanner.Scan() {
				line := scanner.Text()
				fields := strings.Fields(line)

				if len(fields) < 8 {
					continue
				}

				mountPoint := fields[4]
				fsType := fields[7] // Separator "-" is at field 6, fstype at field 7

				// Check for cgroup2 unified hierarchy at /sys/fs/cgroup
				if mountPoint == "/sys/fs/cgroup" && fsType == "cgroup2" {
					foundCgroup2 = true
					break
				}
			}

			Expect(foundCgroup2).To(BeTrue(), "Expected cgroup v2 unified hierarchy to be mounted at /sys/fs/cgroup with type cgroup2")
		})
	})
})
