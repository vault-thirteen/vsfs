package cli

import (
	"errors"
	"flag"
	"fmt"
)

const (
	ErrFServerListenHost       = "host name is not valid: '%v'"
	ErrFServerListenPort       = "port is not valid: '%v'"
	ErrFServerSharedFolderPath = "shared folder is not valid: '%v'"

	ErrNullPointer = "null pointer"
)

type Arguments struct {
	ServerListenHost string
	ServerListenPort uint16
	SharedFolderPath string
}

func (a *Arguments) IsValid() (bool, error) {
	if len(a.ServerListenHost) < 1 {
		return false, fmt.Errorf(ErrFServerListenHost, a.ServerListenHost)
	}

	if (a.ServerListenPort) < 1 {
		return false, fmt.Errorf(ErrFServerListenPort, a.ServerListenPort)
	}

	if len(a.SharedFolderPath) < 1 {
		return false, fmt.Errorf(ErrFServerSharedFolderPath, a.SharedFolderPath)
	}

	return true, nil
}

func NewArgumentsFromOs(
	argumentNameServerListenHost string,
	argumentNameServerListenPort string,
	argumentNameSharedFolderPath string,
) (args *Arguments, err error) {
	serverListenHostPtr := flag.String(argumentNameServerListenHost, "", "Host name of the server")
	serverListenPortPtr := flag.Uint(argumentNameServerListenPort, 0, "Port number of the server")
	sharedFolderPathPtr := flag.String(argumentNameSharedFolderPath, "", "Path to the shared folder")

	flag.Parse()

	if (serverListenHostPtr == nil) ||
		(serverListenPortPtr == nil) ||
		(sharedFolderPathPtr == nil) {
		return nil, errors.New(ErrNullPointer)
	}

	args = &Arguments{
		ServerListenHost: *serverListenHostPtr,
		ServerListenPort: uint16(*serverListenPortPtr),
		SharedFolderPath: *sharedFolderPathPtr,
	}

	_, err = args.IsValid()
	if err != nil {
		return nil, err
	}

	return args, nil
}
