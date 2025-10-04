package docgen

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

func (d *DocumentGenerator) parseMetaDirectory() (*AnsibleMeta, error) {
	var (
		metaInfo AnsibleMeta
		yamlFile []byte
		err      error
	)

	metaFile1 := path.Join(d.RoleDirectory, "meta", "main.yml")
	metaFile2 := path.Join(d.RoleDirectory, "meta", "main.yaml")

	log.Debug().Msgf("reading meta from %s", metaFile1)

	yamlFile, err = os.ReadFile(filepath.Clean(metaFile1))
	if errors.Is(err, os.ErrNotExist) {
		log.Debug().Msgf("reading meta from %s", metaFile2)

		yamlFile, err = os.ReadFile(filepath.Clean(metaFile2))
		if errors.Is(err, os.ErrNotExist) {
			log.Debug().Msg("meta file not found")

			return &metaInfo, nil
		}
	}

	err = yaml.Unmarshal(yamlFile, &metaInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal meta: %w", err)
	}

	return &metaInfo, nil
}
