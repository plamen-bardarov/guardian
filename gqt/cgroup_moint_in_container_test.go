package gqt_test

import (
	"strings"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/guardian/gqt/runner"
	gardencgroups "code.cloudfoundry.org/guardian/rundmc/cgroups"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("cgroup mounts inside container", func() {
	var (
		client     *runner.RunningGarden
		container  garden.Container
		limits     garden.Limits
		privileged bool
	)

	BeforeEach(func() {
		privileged = false
	})

	JustBeforeEach(func() {
		client = runner.Start(config)
		var err error
		container, err = client.Create(garden.ContainerSpec{
			Limits:     limits,
			Privileged: privileged,
		})
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(client.DestroyAndStop()).To(Succeed())
	})

	It("mounts cgroups in the container", func() {
		// Read mountinfo from inside the container's mount namespace
		out := runInContainerCombinedOutput(container, "cat", []string{"/proc/self/mountinfo"})

		if gardencgroups.IsCgroup2UnifiedMode() {
			// cgroup v2 unified hierarchy
			// Look for the ' - cgroup2 cgroup2 ' separator fields and the /sys/fs/cgroup mountpoint
			Expect(out).To(ContainSubstring(" - cgroup2 cgroup2 "))
			Expect(out).To(Or(
				ContainSubstring(" /sys/fs/cgroup "),
				ContainSubstring(" /sys/fs/cgroup/ "),
			))
		} else {
			// cgroup v1 — at least one cgroup controller mount is expected
			// e.g., entries like: ' - cgroup cgroup ... devices,' or cpu, memory, etc.
			Expect(out).To(ContainSubstring(" - cgroup cgroup "))
			// Optionally assert known controllers or options
			Expect(strings.ToLower(out)).To(SatisfyAny(
				ContainSubstring(" devices"),
				ContainSubstring(" cpu"),
				ContainSubstring(" memory"),
			))
		}
	})
})
