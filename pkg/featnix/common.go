package featnix

import (
	"bufio"
	"os"
	"strings"
)

// OSReleaseField represents a field in the /etc/os-release file.
//
// The /etc/os-release file contains information about the operating system.
// This file is used by various Linux distributions to provide information
// about the distribution and its version.
//
// The file contains key-value pairs of the form "KEY=VALUE".
// The keys and values are separated by an equal sign.
//
// The OSReleaseField type represents a field in the /etc/os-release file.
type OSReleaseField string

// String returns the string representation of an OSReleaseField.
//
// It returns the string value of the OSReleaseFieldID constant.
//
// The String method satisfies the Stringer interface.
func (f OSReleaseField) String() string {
	return string(OSReleaseFieldID)
}

// OSReleaseFieldID is the constant representing the "ID" field in the
// "/etc/os-release" file.
const OSReleaseFieldID OSReleaseField = "ID"

// osReleasePath is the path to the file that contains the operating system release
// information on Linux systems.
//
// The ID field in the file is used to identify the Linux distribution.
// The file is typically located at "/etc/os-release".
const osReleasePath = "/etc/os-release"

// readOSReleaseField reads the contents of the specified file and returns the value of the
// specified field.
//
// Parameters:
//   - filePath: The path to the file to read.
//   - field: The field to search for in the file.
//
// Returns:
//   - The value of the specified field in the file. If the field is not found, an empty string
//     is returned.
func readOSReleaseField(filePath, field string) string {
	// Open the file.
	file, err := os.Open(filePath)
	if err != nil {
		return "" // Return an empty string if the file cannot be opened.
	}
	defer file.Close()

	// Create a scanner to read the file line by line.
	scanner := bufio.NewScanner(file)

	// Iterate over each line in the file.
	for scanner.Scan() {
		// Read the line.
		line := scanner.Text()
		// Split the line into key-value pairs.
		parts := strings.Split(line, "=")

		// Check if the line contains the specified field.
		if len(parts) == 2 && parts[0] == field {
			// Trim the value of the field to remove any surrounding quotes.
			value := strings.Trim(parts[1], "\"'")

			return value // Return the value of the field.
		}
	}

	// Return an empty string if the field is not found.
	return ""
}

// readOSReleaseID reads the contents of the /etc/os-release file and returns the value of the
// "ID" field as an OSReleaseID.
//
// This function reads the contents of the file line by line and retrieves the value of the "ID"
// field. The value is returned as an OSReleaseID type. If the "ID" field is not found in the file,
// an empty string is returned.
//
// Returns:
//
//	OSReleaseID: The value of the "ID" field in the file as an OSReleaseID. If the "ID" field
//	  is not found in the file, an empty string is returned.
func readOSReleaseID() OSReleaseID {
	// Read the contents of the file specified by osReleasePath and find the value of the "ID" field.
	// The value of the "ID" field is returned as an OSReleaseID type.
	// If the "ID" field is not found in the file, an empty string is returned.
	//
	// Call the readOSReleaseField function to read the contents of the file and find the value of the "ID" field.
	// The readOSReleaseField function takes the file path and the field name as parameters.
	// It returns a string containing the value of the field.
	// Convert the string to an OSReleaseID type and return it.
	return OSReleaseID(
		readOSReleaseField(
			osReleasePath,             // The path to the file to read.
			OSReleaseFieldID.String(), // The field to search for in the file.
		),
	)
}
