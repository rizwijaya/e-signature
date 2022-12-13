package ipfs

import (
	"e-signature/app/config"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

func ConnectIPFS() *shell.Shell {
	var sh *shell.Shell
	conf, _ := config.Init()
	sh = shell.NewShell(conf.IPFS.Host + ":" + conf.IPFS.Port)
	return sh
}

func UploadIPFS(path string) (string, error) {
	sh := ConnectIPFS()
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	var r = f
	cid, err := sh.Add(r)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_ = os.Remove(path)
	return cid, nil
}

func GetFileIPFS(hash string, output string, directory string) (string, error) {
	sh := ConnectIPFS()
	outputName := directory + output
	err := sh.Get(hash, outputName)
	if err != nil {
		return "", err
	}
	return outputName, nil
}
