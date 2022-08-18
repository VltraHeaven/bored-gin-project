package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type distro struct {
	Name           string `json:"name"`
	ReleaseCycle   string `json:"releaseCycle"`
	LatestVersion  string `json:"latest"`
	LatestKernel   string `json:"latestKernel"`
	PackageManager string `json:"packageManager"`
	IsImmutable    bool   `json:"isImmutable"`
	LibC           string `json:"libc"`
	Compiler       string `json:"compiler"`
}

var distros = []distro{
	{
		Name:           "ArchLinux",
		ReleaseCycle:   "Rolling",
		LatestVersion:  "N/A",
		LatestKernel:   "5.19.1",
		PackageManager: "pacman",
		IsImmutable:    false,
		LibC:           "glibc",
		Compiler:       "gcc",
	},
	{
		Name:           "Debian",
		ReleaseCycle:   "Standard",
		LatestVersion:  "11 (Bullseye)",
		LatestKernel:   "5.10.0",
		PackageManager: "apt/dpkg",
		IsImmutable:    false,
		LibC:           "glibc",
		Compiler:       "gcc",
	},
	{
		Name:           "Fedora Silverblue",
		ReleaseCycle:   "Standard",
		LatestVersion:  "36",
		LatestKernel:   "5.18",
		PackageManager: "rpm-ostree",
		IsImmutable:    true,
		LibC:           "glibc",
		Compiler:       "gcc",
	},
}

func getDistros(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, distros)
}

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1", "localhost"})
	router.GET("/v1/distros", getDistros)
	router.Run("localhost:8080")
}
