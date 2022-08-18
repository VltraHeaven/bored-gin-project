package distro

type distro struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	ReleaseCycle   string `json:"releaseCycle"`
	LatestVersion  string `json:"latestVersion"`
	LatestKernel   string `json:"latestKernel"`
	PackageManager string `json:"packageManager"`
	IsImmutable    bool   `json:"isImmutable"`
	LibC           string `json:"libc"`
	Compiler       string `json:"compiler"`
}
