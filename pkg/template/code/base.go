package code

import (
	_ "embed"
	"os"
	"path"
)

//go:embed .dockerignore
var dockerignore string

//go:embed Dockerfile
var dockerfile string

//go:embed .gitignore
var gitignore string

func RenderFiles(targetPath string) error {
	contents := []string{dockerignore, dockerfile, gitignore}
	names := []string{".dockerignore", "Dockerfile", ".gitignore"}

	for i, c := range contents {
		name := names[i]
		filePath := path.Join(targetPath, name)

		f, err := os.Create(filePath)
		if err != nil {
			return err
		}

		_, err = f.WriteString(c)
	}

	return nil
}
