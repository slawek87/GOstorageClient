package conf


type SettingsInterface interface {
	SetSettings(settings map[string]string)
	GetSettings(name string) string
}


type GOstorageSettings struct {
	settings	map[string]string
}

func (goStorageSettings *GOstorageSettings) SetSettings(settings map[string]string) {
	goStorageSettings.settings = settings
}

func (goStorageSettings *GOstorageSettings) GetSettings(name string) string {
	return goStorageSettings.settings[name]
}
