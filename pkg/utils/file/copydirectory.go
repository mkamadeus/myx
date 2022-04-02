package file

import (
	"io"
	"io/ioutil"
	"os"
	"path"
)

func CopyDirectory(srcDir, destDir string) error {
	entries, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return err
	}

	for _, e := range entries {
		srcPath := path.Join(srcDir, e.Name())
		destPath := path.Join(destDir, e.Name())

		fileInfo, err := os.Stat(srcPath)
		if err != nil {
			return err
		}

		modeType := fileInfo.Mode() & os.ModeType
		if modeType == os.ModeDir {
			return CopyDirectory(srcPath, destPath)
		}

		out, err := os.Create(destPath)
		if err != nil {
			return err
		}

		defer out.Close()

		in, err := os.Open(srcPath)
		defer in.Close()
		if err != nil {
			return err
		}

		_, err = io.Copy(out, in)
		if err != nil {
			return err
		}

	}
	return nil
}
