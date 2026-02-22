package store

func (s *Store) populateDefaults(lang string) error {
	s.emitEvent(StoreEventDatabaseNotFound)
	s.emitEvent(StoreEventCreatingDefaults)

	translations := s.languageMap(lang)

	rules := []struct {
		ext    string
		folder string
	}{
		// Music
		{"mp3", translations["music"]},
		{"aac", translations["music"]},
		{"flac", translations["music"]},
		{"ogg", translations["music"]},
		{"wma", translations["music"]},
		{"m4a", translations["music"]},
		{"aiff", translations["music"]},
		{"wav", translations["music"]},
		{"amr", translations["music"]},

		// Videos
		{"flv", translations["videos"]},
		{"ogv", translations["videos"]},
		{"avi", translations["videos"]},
		{"mp4", translations["videos"]},
		{"mpg", translations["videos"]},
		{"mpeg", translations["videos"]},
		{"3gp", translations["videos"]},
		{"mkv", translations["videos"]},
		{"ts", translations["videos"]},
		{"webm", translations["videos"]},
		{"vob", translations["videos"]},
		{"wmv", translations["videos"]},

		// Pictures
		{"png", translations["pictures"]},
		{"jpeg", translations["pictures"]},
		{"gif", translations["pictures"]},
		{"jpg", translations["pictures"]},
		{"bmp", translations["pictures"]},
		{"svg", translations["pictures"]},
		{"webp", translations["pictures"]},
		{"psd", translations["pictures"]},
		{"tiff", translations["pictures"]},

		// Archives
		{"rar", translations["archives"]},
		{"zip", translations["archives"]},
		{"7z", translations["archives"]},
		{"gz", translations["archives"]},
		{"bz2", translations["archives"]},
		{"tar", translations["archives"]},
		{"dmg", translations["archives"]},
		{"tgz", translations["archives"]},
		{"xz", translations["archives"]},
		{"iso", translations["archives"]},
		{"cpio", translations["archives"]},

		// Documents
		{"txt", translations["documents"]},
		{"pdf", translations["documents"]},
		{"doc", translations["documents"]},
		{"docx", translations["documents"]},
		{"odf", translations["documents"]},
		{"xls", translations["documents"]},
		{"xlsv", translations["documents"]},
		{"xlsx", translations["documents"]},
		{"ppt", translations["documents"]},
		{"pptx", translations["documents"]},
		{"ppsx", translations["documents"]},
		{"odp", translations["documents"]},
		{"odt", translations["documents"]},
		{"ods", translations["documents"]},
		{"md", translations["documents"]},
		{"json", translations["documents"]},
		{"csv", translations["documents"]},

		// Books
		{"mobi", translations["books"]},
		{"epub", translations["books"]},
		{"chm", translations["books"]},

		// DEB Packages
		{"deb", translations["deb_packages"]},

		// Programs
		{"exe", translations["programs"]},
		{"msi", translations["programs"]},

		// RPM Packages
		{"rpm", translations["rpm_packages"]},
	}

	for _, r := range rules {
		if err := s.InsertRule(r.ext + ":" + r.folder); err != nil {
			return err
		}
	}

	s.emitEvent(StoreEventDefaultsInitialized)
	return nil
}

func (s *Store) languageMap(lang string) map[string]string {
	translations := make(map[string]string)

	switch lang {
	case "pt":
		translations["music"] = "Musicas"
		translations["videos"] = "Videos"
		translations["pictures"] = "Imagens"
		translations["archives"] = "Arquivos"
		translations["documents"] = "Documentos"
		translations["books"] = "Livros"
		translations["deb_packages"] = "PacotesDEB"
		translations["programs"] = "Programas"
		translations["rpm_packages"] = "PacotesRPM"

	case "tr":
		translations["music"] = "Müzikler"
		translations["videos"] = "Videolar"
		translations["pictures"] = "Resimler"
		translations["archives"] = "Arşivler"
		translations["documents"] = "Dokümanlar"
		translations["books"] = "Kitaplar"
		translations["deb_packages"] = "DEBPaketleri"
		translations["programs"] = "Programlar"
		translations["rpm_packages"] = "RPMPaketleri"

	default:
		translations["music"] = "Music"
		translations["videos"] = "Videos"
		translations["pictures"] = "Pictures"
		translations["archives"] = "Archives"
		translations["documents"] = "Documents"
		translations["books"] = "Books"
		translations["deb_packages"] = "DEBPackages"
		translations["programs"] = "Programs"
		translations["rpm_packages"] = "RPMPackages"
	}

	return translations
}
