package conf

var Settings SettingsInterface = &GOstorageSettings{}

func init() {
	Settings.SetSettings(map[string]string{
		"HOST":          "0.0.0.0",
		"PROTOCOL":      "http",
		"PORT":          "8070",
		"USERNAME": "fb008237-c46a-4823-b027-c2ea2f49340d",
		"PASSWORD": "GOstorage",
	})
}
