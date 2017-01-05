package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os/user"
	"path/filepath"
	"strings"
)

func testDb(file string) bool {
	cfgTest, err := ini.LoadSources(ini.LoadOptions{AllowBooleanKeys: true, Loose: false}, file)

	if err != nil {
		return false
	}

	cfg = cfgTest

	return true
}

func initDb() {

	if language == "" {
		language = "en"
	}

	//test if exist a configFile in the directory
	cfgFile = strings.Replace(configFile, "{lang}", language, 1)
	if testDb(cfgFile) {
		return
	}

	//test if exist a configFile in the user's home
	currentUser, _ := user.Current()
	cfgFile = filepath.Join(currentUser.HomeDir, cfgFile)
	if testDb(cfgFile) {
		return
	}

	//Create a default database in user's home
	cfg = ini.Empty()
	defaultDb()

}

func closeDb() {
	err := cfg.SaveTo(cfgFile)

	if err != nil {
		panic(err)
	}
}

func defaultDb() {

	lang := langVars()

	fmt.Println("No database found")
	fmt.Println("Creating default database")

	//Music
	insertRule("mp3:" + lang["music"])
	insertRule("aac:" + lang["music"])
	insertRule("flac:" + lang["music"])
	insertRule("ogg:" + lang["music"])
	insertRule("wma:" + lang["music"])
	insertRule("m4a:" + lang["music"])
	insertRule("aiff:" + lang["music"])
	insertRule("wav:" + lang["music"])
	insertRule("amr:" + lang["music"])

	//Videos
	insertRule("flv:" + lang["videos"])
	insertRule("ogv:" + lang["videos"])
	insertRule("avi:" + lang["videos"])
	insertRule("mp4:" + lang["videos"])
	insertRule("mpg:" + lang["videos"])
	insertRule("mpeg:" + lang["videos"])
	insertRule("3gp:" + lang["videos"])
	insertRule("mkv:" + lang["videos"])
	insertRule("ts:" + lang["videos"])
	insertRule("webm:" + lang["videos"])
	insertRule("vob:" + lang["videos"])
	insertRule("wmv:" + lang["videos"])

	//Pictures
	insertRule("png:" + lang["pictures"])
	insertRule("jpeg:" + lang["pictures"])
	insertRule("gif:" + lang["pictures"])
	insertRule("jpg:" + lang["pictures"])
	insertRule("bmp:" + lang["pictures"])
	insertRule("svg:" + lang["pictures"])
	insertRule("webp:" + lang["pictures"])
	insertRule("psd:" + lang["pictures"])
	insertRule("tiff:" + lang["pictures"])

	//Archives
	insertRule("rar:" + lang["archives"])
	insertRule("zip:" + lang["archives"])
	insertRule("7z:" + lang["archives"])
	insertRule("gz:" + lang["archives"])
	insertRule("bz2:" + lang["archives"])
	insertRule("tar:" + lang["archives"])
	insertRule("dmg:" + lang["archives"])
	insertRule("tgz:" + lang["archives"])
	insertRule("xz:" + lang["archives"])
	insertRule("iso:" + lang["archives"])
	insertRule("cpio:" + lang["archives"])

	//Documents
	insertRule("txt:" + lang["documents"])
	insertRule("pdf:" + lang["documents"])
	insertRule("doc:" + lang["documents"])
	insertRule("docx:" + lang["documents"])
	insertRule("odf:" + lang["documents"])
	insertRule("xls:" + lang["documents"])
	insertRule("xlsv:" + lang["documents"])
	insertRule("xlsx:" + lang["documents"])
	insertRule("ppt:" + lang["documents"])
	insertRule("pptx:" + lang["documents"])
	insertRule("ppsx:" + lang["documents"])
	insertRule("odp:" + lang["documents"])
	insertRule("odt:" + lang["documents"])
	insertRule("ods:" + lang["documents"])
	insertRule("md:" + lang["documents"])
	insertRule("json:" + lang["documents"])
	insertRule("csv:" + lang["documents"])

	//Books
	insertRule("mobi:" + lang["books"])
	insertRule("epub:" + lang["books"])
	insertRule("chm:" + lang["books"])

	//DEBPackages
	insertRule("deb:" + lang["deb_packages"])

	//Programs
	insertRule("exe:" + lang["programs"])
	insertRule("msi:" + lang["programs"])

	//RPMPackages
	insertRule("rpm:" + lang["rpm_packages"])

	fmt.Println("Default database initialized")
}

func langVars() map[string]string {

	lang := make(map[string]string)

	switch language {
	case "tr":
		lang["music"] = "Müzikler"
		lang["videos"] = "Videolar"
		lang["pictures"] = "Resimler"
		lang["archives"] = "Arşivler"
		lang["documents"] = "Dokümanlar"
		lang["books"] = "Kitaplar"
		lang["deb_packages"] = "DEBPaketleri"
		lang["programs"] = "Programlar"
		lang["rpm_packages"] = "RPMPaketleri"

	default:
		lang["music"] = "Music"
		lang["videos"] = "Videos"
		lang["pictures"] = "Pictures"
		lang["archives"] = "Archives"
		lang["documents"] = "Documents"
		lang["books"] = "Books"
		lang["deb_packages"] = "DEBPackages"
		lang["programs"] = "Programs"
		lang["rpm_packages"] = "RPMPackages"
	}

	return lang
}
