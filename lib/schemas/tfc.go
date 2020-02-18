package schemas

type TFC struct {
	Organization            string `hcl:"organization"`
	Workspace               string `hcl:"workspace"`
	PurgeUnmanagedVariables *bool  `hcl:"purge-unmanaged-variables"`
}