package sync

func (o Operation) syncReleaseNext() error {
	return runSteps([]step{
		o.resetReleaseNext,
		o.addForkFiles(nextRelease{}),
		o.applyPatches,
		o.pushBranch(o.Config.Branches.ReleaseNext),
	})
}

func (o Operation) pushBranch(branch string) step {
	return func() error {
		p := push{
			State:  o.State,
			branch: branch,
		}
		return runSteps(p.steps())
	}
}

type nextRelease struct{}

func (n nextRelease) String() string {
	return "next"
}

func (n nextRelease) Name(string) (string, error) {
	return "release-" + n.String(), nil
}

func (n nextRelease) less(release) bool {
	return false
}

func (n nextRelease) Tag() string {
	return "knative-next"
}
