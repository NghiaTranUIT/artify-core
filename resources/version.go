package resources

import (
	"github.com/NghiaTranUIT/artify-core/models"
	"github.com/mcuadros/go-version"
)

func (r *Resource) GetLatestAppVersion() (*models.Version, error) {
	version := models.Version{}
	err := r.PostgreSQL.Last(&version)

	if err != nil {
		return nil, err
	}

	return &version, nil
}

func (r *Resource) GetUpdateAppVersion(buildVersion string) (*models.Version, error) {
	latestVersion, err := r.GetLatestAppVersion()
	if latestVersion == nil {
		return nil, err
	}
	// Compare
	result := version.CompareSimple(latestVersion.BuildVersion, buildVersion)

	// Has No new version
	if result <= 0 {
		return nil, nil
	}

	// New version
	return latestVersion, nil
}

func (r *Resource) GenerateSampleVersion() (*models.Version, error) {
	version := models.Version{
		BuildVersion: "1.0.0",
		BuildNumber:  1000,
		Url:          "https://www.dropbox.com/s/cxwtqvrgblr4t0b/Artify_1.0.0.zip",
		Notes:        "First version Artify",
	}
	err := r.PostgreSQL.Create(&version)
	if err != nil {
		return nil, err
	}
	return &version, nil
}
