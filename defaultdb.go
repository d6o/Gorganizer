package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

func initDb() {

	cfg, _ = ini.LoadSources(ini.LoadOptions{AllowBooleanKeys: true, Loose: true}, configFile)

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		cfg = ini.Empty()
		defaultDb()
	}
}

func closeDb() {
	err := cfg.SaveTo(configFile)

	if err != nil {
		panic(err)
	}
}

func defaultDb() {
	fmt.Println("No database found")
	fmt.Println("Creating default database")

	//Music
	insertRule("mp3:Musics")
	insertRule("aac:Musics")
	insertRule("flac:Musics")
	insertRule("ogg:Musics")
	insertRule("wma:Musics")
	insertRule("m4a:Musics")
	insertRule("aiff:Musics")
	insertRule("wav:Musics")
	insertRule("amr:Musics")

	//Videos
	insertRule("flv:Videos")
	insertRule("ogv:Videos")
	insertRule("avi:Videos")
	insertRule("mp4:Videos")
	insertRule("mpg:Videos")
	insertRule("mpeg:Videos")
	insertRule("3gp:Videos")
	insertRule("mkv:Videos")
	insertRule("ts:Videos")
	insertRule("webm:Videos")
	insertRule("vob:Videos")
	insertRule("wmv:Videos")

	//Pictures
	insertRule("png:Pictures")
	insertRule("jpeg:Pictures")
	insertRule("gif:Pictures")
	insertRule("jpg:Pictures")
	insertRule("bmp:Pictures")
	insertRule("svg:Pictures")
	insertRule("webp:Pictures")
	insertRule("psd:Pictures")
	insertRule("tiff:Pictures")

	//Archives
	insertRule("rar:Archives")
	insertRule("zip:Archives")
	insertRule("7z:Archives")
	insertRule("gz:Archives")
	insertRule("bz2:Archives")
	insertRule("tar:Archives")
	insertRule("dmg:Archives")
	insertRule("tgz:Archives")
	insertRule("xz:Archives")
	insertRule("iso:Archives")
	insertRule("cpio:Archives")

	//Documents
	insertRule("txt:Documents")
	insertRule("pdf:Documents")
	insertRule("doc:Documents")
	insertRule("docx:Documents")
	insertRule("odf:Documents")
	insertRule("xls:Documents")
	insertRule("xlsv:Documents")
	insertRule("xlsx:Documents")
	insertRule("ppt:Documents")
	insertRule("pptx:Documents")
	insertRule("ppsx:Documents")
	insertRule("odp:Documents")
	insertRule("odt:Documents")
	insertRule("ods:Documents")
	insertRule("md:Documents")
	insertRule("json:Documents")
	insertRule("csv:Documents")

	//Books
	insertRule("mobi:Books")
	insertRule("epub:Books")
	insertRule("chm:Books")

	//DEBPackages
	insertRule("deb:DEBPackages")

	//Programs
	insertRule("exe:Programs")
	insertRule("msi:Programs")

	//RPMPackages
	insertRule("rpm:RPMPackages")

	fmt.Println("Default database initialized")
}
