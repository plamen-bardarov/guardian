package runcontainerd

import (
	"code.cloudfoundry.org/guardian/rundmc/goci"
)

func (m cgroupManager) setUnifiedResources(bundle goci.Bndl) error {
	return nil
}

func (m cgroupManager) mountUnifiedCgroup(bundle *goci.Bndl) { return nil }
