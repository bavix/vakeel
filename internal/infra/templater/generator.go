package templater

import (
	"bytes"
	_ "embed"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/google/uuid"

	"github.com/bavix/vakeel/pkg/featnix"
)

// errUnsupportedOS is the error returned when the operating system is unsupported.
var errUnsupportedOS = errors.New("unsupported operating system")

// errStubNotFound is the error returned when the stub template is not found.
//
// It is used when the stub template file is not found during the generation of the
// init script or systemd service file.
var errStubNotFound = errors.New("stub template not found")

// openwrtServicePath is the path to the init script for the vakeel agent
// on OpenWRT systems.
//
// The init script is responsible for starting and stopping the vakeel agent.
const openwrtServicePath = "/tmp/etc/init.d/vakeel"

// systemdServicePath is the path to the systemd service file for the vakeel agent
// on systems that use systemd as the service manager.
//
// The systemd service file defines the configuration for the vakeel agent as a
// systemd service. It specifies the command to start the agent, the behavior
// when the agent crashes or is terminated, and other options.
const systemdServicePath = "/tmp/etc/systemd/system/vakeel.service"

//go:embed openwrt.stub
var openwrtTemplate string

//go:embed systemd.stub
var systemdTemplate string

// Data contains the data used to fill the stub agent template.
type Data struct {
	// AppPath is the path to the application binary.
	AppPath string
	// ID is the ID of the agent.
	ID uuid.UUID
	// Host is the hostname of the vakeel server.
	Host string
	// Port is the port of the vakeel server.
	Port int
}

// ServiceGenerator is a template for creating stub agents.
//
// It contains the data used to fill the stub agent template.
type ServiceGenerator struct {
	// context contains the context used to fill the stub agent template.
	// It contains the path to the application binary, the ID of the agent,
	// the hostname of the vakeel server, and the port of the vakeel server.
	context Data
}

// New creates a new ServiceGenerator instance with the given ID, host, and port.
//
// Parameters:
// - id: The ID of the agent.
// - host: The hostname of the vakeel server.
// - port: The port of the vakeel server.
//
// Returns:
// - *ServiceGenerator: A pointer to the ServiceGenerator instance.
// - error: An error if any.
func New(
	id uuid.UUID, // The ID of the agent.
	host string, // The hostname of the vakeel server.
	port int, // The port of the vakeel server.
) (*ServiceGenerator, error) {
	// Get the path to the application binary.
	//
	// The os.Executable function returns the path name for the executable file that is currently
	// running. If there is an error, it returns the error.
	appPath, err := os.Executable()
	if err != nil {
		return nil, err // Return the error if any.
	}

	// Create a new ServiceGenerator instance with the given ID, host, port, and appPath.
	return &ServiceGenerator{
		// Initialize the context with the given ID, host, port, and appPath.
		context: Data{
			AppPath: appPath, // The path to the application binary.
			ID:      id,      // The ID of the agent.
			Host:    host,    // The hostname of the vakeel server.
			Port:    port,    // The port of the vakeel server.
		},
	}, nil
}

// stub returns the stub agent template based on the operating system.
//
// This function checks the operating system and returns the appropriate template
// based on the result. If the operating system is OpenWrt, it returns the
// openwrtTemplate. If the operating system supports systemd, it returns the
// systemdTemplate. If the operating system is neither OpenWrt nor supports
// systemd, it returns nil.
//
// The stub agent template is a Go template that is used to create stub agents.
// The template contains placeholders for the application binary path, agent ID,
// the hostname of the vakeel server, and the port of the vakeel server.
//
// Returns:
//
//	*string: The stub agent template. Nil if the operating system is neither
//	OpenWrt nor supports systemd.
func (t *ServiceGenerator) stub() *string {
	// Check if the operating system is OpenWrt.
	// If it is, return the openwrtTemplate.
	if featnix.IsOpenWrt() {
		return &openwrtTemplate
	}

	// Check if the operating system supports systemd.
	// If it does, return the systemdTemplate.
	if featnix.HasSystemd() {
		return &systemdTemplate
	}

	// The operating system is neither OpenWrt nor supports systemd.
	// Return nil.
	return nil
}

// Register registers the stub agent service file based on the operating system.
//
// This function calls the generate function to generate the stub agent service
// file. Then, it checks if the operating system supports systemd or is OpenWrt.
// If it does, it enables and starts the service using the systemctl or
// /etc/init.d/vakeel commands respectively. If the operating system is neither
// OpenWrt nor supports systemd, it returns an error indicating that the
// operating system is unsupported.
//
// Returns:
//
//	error: An error if the operating system is unsupported or if there is an
//	error enabling or starting the service. Nil if the registration is
//	successful.
func (t *ServiceGenerator) Register() error {
	if _, err := t.generate(); err != nil {
		return err
	}

	if featnix.HasSystemd() {
		return t.registerSystemd()
	}

	if featnix.IsOpenWrt() {
		return t.registerOpenWrt()
	}

	return errStubNotFound
}

// registerSystemd enables and starts the Vakeel service using systemctl.
//
// This function uses the systemctl command to enable and start the Vakeel service.
// If there is an error enabling or starting the service, it will be returned.
//
// Returns:
//
//	error: An error if there is an error enabling or starting the service.
//	Nil if the service is enabled and started successfully.
func (t *ServiceGenerator) registerSystemd() error {
	// Enable the Vakeel service using systemctl.
	// If there is an error enabling the service, return the error.
	if err := exec.Command("systemctl", "enable", "vakeel.service").Run(); err != nil {
		return err
	}

	// Start the Vakeel service using systemctl.
	// If there is an error starting the service, return the error.
	return exec.Command("systemctl", "start", "vakeel.service").Run()
}

// registerOpenWrt registers the Vakeel service using /etc/init.d/vakeel.
//
// This function uses the /etc/init.d/vakeel script to enable and start the Vakeel
// service. If there is an error enabling or starting the service, it will be
// returned.
//
// Returns:
//
//	error: An error if there is an error enabling or starting the service.
//	Nil if the service is enabled and started successfully.
func (t *ServiceGenerator) registerOpenWrt() error {
	// Enable the Vakeel service using /etc/init.d/vakeel.
	// If there is an error enabling the service, return the error.
	if err := exec.Command(openwrtServicePath, "enable").Run(); err != nil {
		return err
	}

	// Start the Vakeel service using /etc/init.d/vakeel.
	// If there is an error starting the service, return the error.
	return exec.Command(openwrtServicePath, "start").Run()
}

// generate generates the stub agent service file based on the operating system.
//
// This function gets the stub agent template based on the operating system,
// parses the template, and executes it with the StubTemplate instance as the
// data. Finally, it returns the generated template as a string.
//
// Returns:
//
//	string: The generated stub agent service file.
//	error:  An error if there is an error parsing or executing the template.
func (t *ServiceGenerator) generate() (string, error) {
	// Render the stub agent template.
	// The template is parsed and executed with the StubTemplate instance as the data.
	// The generated template is returned as a string.
	content, err := t.render()
	if err != nil {
		return "", err
	}

	// Check if the operating system is OpenWrt.
	// If it is, write the content to the openwrt service file.
	if featnix.IsOpenWrt() {
		return t.writeToFile(openwrtServicePath, content)
	}

	// Check if the operating system supports systemd.
	// If it does, write the content to the systemd service file.
	if featnix.HasSystemd() {
		return t.writeToFile(systemdServicePath, content)
	}

	// If the operating system is neither OpenWrt nor supports systemd,
	// return an error indicating that the operating system is unsupported.
	return "", errUnsupportedOS
}

// writeToFile writes the content to the specified file path, if the file does not already exist.
//
// Parameters:
// - filePath: The path to the file where the content will be written.
// - content: The content to be written to the file.
//
// Returns:
// - string: The file path.
// - error: An error if there is an error writing the file.
func (t *ServiceGenerator) writeToFile(filePath string, content string) (string, error) {
	// Create the directory for the file if it doesn't exist.
	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return "", err
	}

	// Check if the file already exists.
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		// If the file exists, return the file path.
		return filePath, nil
	}

	// Open the file for writing.
	f, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// Write the content to the file.
	_, err = f.WriteString(content)
	if err != nil {
		return "", err
	}

	// Return the file path.
	return filePath, nil
}

// render generates the stub agent service file based on the operating system.
//
// This function gets the stub agent template based on the operating system,
// parses the template, and executes it with the StubTemplate instance as the
// data. Finally, it returns the generated template as a string.
//
// Returns:
//
//	string: The generated stub agent service file.
//	error:  An error if there is an error parsing or executing the template.
func (t *ServiceGenerator) render() (string, error) {
	// Get the stub agent template based on the operating system.
	stubTemplate := t.stub()

	// Return an error if the stub agent template is nil.
	if stubTemplate == nil {
		return "", errStubNotFound
	}

	// Parse the stub agent template.
	//
	// The template is parsed using the template package's Parse function. The
	// template is named "stub" and the template string is stored in the
	// stubTemplate variable.
	tmpl, err := template.New("stub").Parse(*stubTemplate)
	if err != nil {
		// Return an error if there is an error parsing the template.
		return "", err
	}

	// Execute the template with the StubTemplate instance as the data.
	//
	// The template is executed using the Execute function of the template. The
	// result is stored in a buffer.
	var buf bytes.Buffer

	err = tmpl.Execute(&buf, t.context)
	if err != nil {
		// Return an error if there is an error executing the template.
		return "", err
	}

	// Return the generated template as a string.
	//
	// The generated template is stored in the buffer and returned as a string.
	return buf.String(), nil
}
