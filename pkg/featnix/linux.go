package featnix

// OSReleaseID represents the ID field in the /etc/os-release file.
// It is used to identify the distribution of the operating system.
type OSReleaseID string

const (
	// OSReleaseIDUnknown represents an unknown OS.
	OSReleaseIDUnknown OSReleaseID = ""

	// OSReleaseIDDebian represents Debian.
	OSReleaseIDDebian OSReleaseID = "debian"

	// OSReleaseIDUbuntu represents Ubuntu.
	OSReleaseIDUbuntu OSReleaseID = "ubuntu"

	// OSReleaseIDFedora represents Fedora.
	OSReleaseIDFedora OSReleaseID = "fedora"

	// OSReleaseIDFreeBSD represents FreeBSD.
	OSReleaseIDFreeBSD OSReleaseID = "freebsd"

	// OSReleaseIDArch represents Arch Linux.
	OSReleaseIDArch OSReleaseID = "arch"

	// OSReleaseIDOpenWrt represents OpenWrt.
	OSReleaseIDOpenWrt OSReleaseID = "openwrt"

	// OSReleaseIDAlpine represents Alpine Linux.
	OSReleaseIDAlpine OSReleaseID = "alpine"

	// OSReleaseIDOpenSUSE represents openSUSE.
	OSReleaseIDOpenSUSE OSReleaseID = "opensuse"
)

// IsDebian checks if the current system is running Debian by reading the contents
// of the /etc/os-release file.
//
// Returns:
//
//	bool: True if the system is running Debian, false otherwise.
func IsDebian() bool {
	return readOSReleaseID() == OSReleaseIDDebian
}

// IsUbuntu checks if the current operating system is Ubuntu by reading the contents of the
// /etc/os-release file.
//
// Returns:
//
//	bool: True if the operating system is Ubuntu, false otherwise.
func IsUbuntu() bool {
	return readOSReleaseID() == OSReleaseIDUbuntu
}

// IsFedora checks if the current system is running Fedora by reading the contents
// of the /etc/os-release file.
//
// Returns:
//
//	bool: True if the system is running Fedora, false otherwise.
func IsFedora() bool {
	// Read the contents of the /etc/os-release file and check if the field "ID"
	// contains the string "fedora". If it does, it returns true indicating that
	// Fedora is running. Otherwise, it returns false.
	return readOSReleaseID() == OSReleaseIDFedora
}

// IsFreeBSD checks if the current system is running FreeBSD by reading the contents
// of the /etc/os-release file.
//
// It returns true if the field "ID" in the /etc/os-release file contains the string "freebsd",
// indicating that FreeBSD is running. Otherwise, it returns false.
//
// Returns:
//
//	bool: True if the system is running FreeBSD, false otherwise.
func IsFreeBSD() bool {
	// Read the contents of the /etc/os-release file and check if the field "ID"
	// contains the string "freebsd". If it does, it returns true indicating that
	// FreeBSD is running. Otherwise, it returns false.
	return readOSReleaseID() == OSReleaseIDFreeBSD
}

// IsArchLinux checks if the current system is running Arch Linux by reading the contents
// of the /etc/os-release file.
//
// It returns true if the field "ID" in the /etc/os-release file contains the string "arch",
// indicating that Arch Linux is running. Otherwise, it returns false.
//
// Returns:
//
//	bool: True if the system is running Arch Linux, false otherwise.
func IsArchLinux() bool {
	// Read the contents of the /etc/os-release file and check if the field "ID"
	// contains the string "arch". If it does, it returns true indicating that
	// Arch Linux is running. Otherwise, it returns false.
	return readOSReleaseID() == OSReleaseIDArch
}

// IsOpenWrt checks if the current system is running OpenWrt by reading the contents
// of the /etc/os-release file.
//
// It returns true if the field "ID" in the /etc/os-release file contains the string "openwrt",
// indicating that OpenWrt is running. Otherwise, it returns false.
//
// Returns:
//
//	bool: True if the system is running OpenWrt, false otherwise.
func IsOpenWrt() bool {
	// Read the contents of the /etc/os-release file and check if the field "ID"
	// contains the string "openwrt". If it does, it returns true indicating that
	// OpenWrt is running. Otherwise, it returns false.
	return readOSReleaseID() == OSReleaseIDOpenWrt
}

// IsAlpine checks if the current system is running Alpine Linux by reading the contents
// of the /etc/os-release file.
//
// It returns true if the field "ID" in the /etc/os-release file contains the string "alpine",
// indicating that Alpine Linux is running. Otherwise, it returns false.
//
// Returns:
//
//	bool: True if the system is running Alpine Linux, false otherwise.
func IsAlpine() bool {
	// Reads the contents of the /etc/os-release file and checks if the field "ID"
	// contains the string "alpine". If it does, it returns true indicating that
	// Alpine Linux is running. Otherwise, it returns false.
	return readOSReleaseID() == OSReleaseIDAlpine
}

// IsOpenSUSE checks if the current system is running openSUSE by reading the contents
// of the /etc/os-release file.
//
// It returns true if the field "ID" in the /etc/os-release file contains the string "opensuse",
// indicating that openSUSE is running. Otherwise, it returns false.
//
// Returns:
//
//	bool: True if the system is running openSUSE, false otherwise.
func IsOpenSUSE() bool {
	// Read the contents of the /etc/os-release file and check if the field "ID"
	// contains the string "opensuse". If it does, it returns true indicating that
	// openSUSE is running. Otherwise, it returns false.
	return readOSReleaseID() == OSReleaseIDOpenSUSE
}
