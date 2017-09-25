package conf

var Settings SettingsInterface = &GOstorageSettings{}

func init() {
	Settings.SetSettings(map[string]string{
		"HOST":          "0.0.0.0",
		"PROTOCOL":      "http",
		"PORT":          "8070",
		"USERNAME": "GOstorageRegister",
		"PASSWORD": "k1k2k3k4",
	})
}
