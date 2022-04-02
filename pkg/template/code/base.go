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

//go:embed requirements.txt
var requirements string

func RenderFiles(targetPath string) error {
	contents := []string{dockerignore, dockerfile, gitignore, requirements}
	names := []string{".dockerignore", "Dockerfile", ".gitignore", "requirements.txt"}

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
