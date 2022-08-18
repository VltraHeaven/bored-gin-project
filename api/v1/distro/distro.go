package distro

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDistros(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, distros)
}

var distros = []distro{
	{
		ID:             "1",
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
		ID:             "2",
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
		ID:             "3",
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

func GetDistroByID(c *gin.Context) {
	id := c.Param("id")
	for _, r := range distros {
		if r.ID == id {
			c.IndentedJSON(http.StatusOK, r)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "distro not found"})
}

func PostDistro(c *gin.Context) {
	var d distro
	if err := c.BindJSON(&d); err != nil {
		return
	}
	distros = append(distros, d)
	c.IndentedJSON(http.StatusCreated, d)
}
