// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ContainerConfig Configuration for a container that is portable between hosts
// swagger:model ContainerConfig
type ContainerConfig struct {

	// Command is already escaped (Windows only)
	ArgsEscaped bool `json:"ArgsEscaped,omitempty"`

	// Whether to attach to `stderr`.
	AttachStderr bool `json:"AttachStderr,omitempty"`

	// Whether to attach to `stdin`.
	AttachStdin bool `json:"AttachStdin,omitempty"`

	// Whether to attach to `stdout`.
	AttachStdout bool `json:"AttachStdout,omitempty"`

	// Command to run specified an array of strings.
	Cmd []string `json:"Cmd"`

	// Whether to generate the network files(/etc/hostname, /etc/hosts and /etc/resolv.conf) for container.
	DisableNetworkFiles bool `json:"DisableNetworkFiles,omitempty"`

	// Set disk quota for container.
	// Key is the dir in container.
	// Value is disk quota size for the dir.
	// / means rootfs dir in container.
	// .* includes rootfs dir and all volume dir.
	//
	DiskQuota map[string]string `json:"DiskQuota,omitempty"`

	// The domain name to use for the container.
	Domainname string `json:"Domainname,omitempty"`

	// The entry point for the container as a string or an array of strings.
	// If the array consists of exactly one empty string (`[""]`) then the entry point is reset to system default.
	//
	Entrypoint []string `json:"Entrypoint"`

	// A list of environment variables to set inside the container in the form `["VAR=value", ...]`. A variable without `=` is removed from the environment, rather than to have an empty value.
	//
	Env []string `json:"Env"`

	// An object mapping ports to an empty object in the form:`{<port>/<tcp|udp>: {}}`
	ExposedPorts map[string]interface{} `json:"ExposedPorts,omitempty"`

	// The hostname to use for the container, as a valid RFC 1123 hostname.
	// Min Length: 1
	// Format: hostname
	Hostname strfmt.Hostname `json:"Hostname,omitempty"`

	// The name of the image to use when creating the container
	// Required: true
	Image string `json:"Image"`

	// Initial script executed in container. The script will be executed before entrypoint or command
	InitScript string `json:"InitScript,omitempty"`

	// User-defined key/value metadata.
	Labels map[string]string `json:"Labels,omitempty"`

	// MAC address of the container.
	MacAddress string `json:"MacAddress,omitempty"`

	// Masks over the provided paths inside the container.
	MaskedPaths []string `json:"MaskedPaths"`

	// net priority.
	NetPriority int64 `json:"NetPriority,omitempty"`

	// Disable networking for the container.
	NetworkDisabled bool `json:"NetworkDisabled,omitempty"`

	// `ONBUILD` metadata that were defined.
	OnBuild []string `json:"OnBuild"`

	// Open `stdin`
	OpenStdin bool `json:"OpenStdin,omitempty"`

	// Set disk quota by specified quota id.
	// If QuotaID <= 0, it means pouchd should allocate a unique quota id by sequence automatically.
	// By default, a quota ID is mapped to only one container. And one quota ID can include several mountpoint.
	//
	QuotaID string `json:"QuotaID,omitempty"`

	// Set the provided paths as RO inside the container.
	ReadonlyPaths []string `json:"ReadonlyPaths"`

	// Whether to start container in rich container mode. (default false)
	Rich bool `json:"Rich,omitempty"`

	// Choose one rich container mode.(default dumb-init)
	// Enum: [dumb-init sbin-init systemd]
	RichMode string `json:"RichMode,omitempty"`

	// Shell for when `RUN`, `CMD`, and `ENTRYPOINT` uses a shell.
	Shell []string `json:"Shell"`

	// annotations send to runtime spec.
	SpecAnnotation map[string]string `json:"SpecAnnotation,omitempty"`

	// Create container with given id.
	// MinLength: 64
	// MaxLength: 64
	// The characters of given id should be in 0123456789abcdef.
	// By default, given id is unnecessary.
	//
	SpecificID string `json:"SpecificID,omitempty"`

	// Close `stdin` after one attached client disconnects
	StdinOnce bool `json:"StdinOnce,omitempty"`

	// Signal to stop a container as a string or unsigned integer.
	StopSignal string `json:"StopSignal,omitempty"`

	// Timeout to stop a container in seconds.
	// Minimum: 0
	StopTimeout *int64 `json:"StopTimeout,omitempty"`

	// Attach standard streams to a TTY, including `stdin` if it is not closed.
	Tty bool `json:"Tty,omitempty"`

	// The user that commands are run as inside the container.
	User string `json:"User,omitempty"`

	// An object mapping mount point paths inside the container to empty objects.
	Volumes map[string]interface{} `json:"Volumes,omitempty"`

	// The working directory for commands to run in.
	WorkingDir string `json:"WorkingDir,omitempty"`
}

// Validate validates this container config
func (m *ContainerConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExposedPorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRichMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopTimeout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// additional properties value enum
var containerConfigExposedPortsValueEnum []interface{}

func init() {
	var res []interface{}
	if err := json.Unmarshal([]byte(`[{}]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		containerConfigExposedPortsValueEnum = append(containerConfigExposedPortsValueEnum, v)
	}
}

func (m *ContainerConfig) validateExposedPortsValueEnum(path, location string, value interface{}) error {
	if err := validate.Enum(path, location, value, containerConfigExposedPortsValueEnum); err != nil {
		return err
	}
	return nil
}

func (m *ContainerConfig) validateExposedPorts(formats strfmt.Registry) error {

	if swag.IsZero(m.ExposedPorts) { // not required
		return nil
	}

	for k := range m.ExposedPorts {

		if err := m.validateExposedPortsValueEnum("ExposedPorts"+"."+k, "body", m.ExposedPorts[k]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ContainerConfig) validateHostname(formats strfmt.Registry) error {

	if swag.IsZero(m.Hostname) { // not required
		return nil
	}

	if err := validate.MinLength("Hostname", "body", string(m.Hostname), 1); err != nil {
		return err
	}

	if err := validate.FormatOf("Hostname", "body", "hostname", m.Hostname.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ContainerConfig) validateImage(formats strfmt.Registry) error {

	if err := validate.RequiredString("Image", "body", string(m.Image)); err != nil {
		return err
	}

	return nil
}

var containerConfigTypeRichModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dumb-init","sbin-init","systemd"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		containerConfigTypeRichModePropEnum = append(containerConfigTypeRichModePropEnum, v)
	}
}

const (

	// ContainerConfigRichModeDumbInit captures enum value "dumb-init"
	ContainerConfigRichModeDumbInit string = "dumb-init"

	// ContainerConfigRichModeSbinInit captures enum value "sbin-init"
	ContainerConfigRichModeSbinInit string = "sbin-init"

	// ContainerConfigRichModeSystemd captures enum value "systemd"
	ContainerConfigRichModeSystemd string = "systemd"
)

// prop value enum
func (m *ContainerConfig) validateRichModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, containerConfigTypeRichModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ContainerConfig) validateRichMode(formats strfmt.Registry) error {

	if swag.IsZero(m.RichMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateRichModeEnum("RichMode", "body", m.RichMode); err != nil {
		return err
	}

	return nil
}

func (m *ContainerConfig) validateStopTimeout(formats strfmt.Registry) error {

	if swag.IsZero(m.StopTimeout) { // not required
		return nil
	}

	if err := validate.MinimumInt("StopTimeout", "body", int64(*m.StopTimeout), 0, false); err != nil {
		return err
	}

	return nil
}

// additional properties value enum
var containerConfigVolumesValueEnum []interface{}

func init() {
	var res []interface{}
	if err := json.Unmarshal([]byte(`[{}]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		containerConfigVolumesValueEnum = append(containerConfigVolumesValueEnum, v)
	}
}

func (m *ContainerConfig) validateVolumesValueEnum(path, location string, value interface{}) error {
	if err := validate.Enum(path, location, value, containerConfigVolumesValueEnum); err != nil {
		return err
	}
	return nil
}

func (m *ContainerConfig) validateVolumes(formats strfmt.Registry) error {

	if swag.IsZero(m.Volumes) { // not required
		return nil
	}

	for k := range m.Volumes {

		if err := m.validateVolumesValueEnum("Volumes"+"."+k, "body", m.Volumes[k]); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerConfig) UnmarshalBinary(b []byte) error {
	var res ContainerConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
