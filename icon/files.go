// Initial config from:
// - https://github.com/athityakumar/colorls/blob/master/lib/yaml/files.yaml
// - https://github.com/athityakumar/colorls/blob/master/lib/yaml/file_aliases.yaml

package icon

import (
	"path/filepath"
	"strings"
)

// ForFile - Return rune icon corresponding to input file name
func ForFile(name string) rune {
	name = strings.ToLower(name)
	ext := filepath.Ext(name)

	if files[name] != 0 {
		return files[name]
	}

	if len(ext) != 0 {
		ext = ext[1:]
	} else {
		ext = name
	}

	alias := fileAliases[ext]
	if alias != "" {
		ext = alias
	}

	icon := files[ext]
	if icon == 0 {
		return fileDefault
	}

	return icon
}

var fileDefault = '\uf15b'

var fileAliases = map[string]string{
	"apk":              "android",
	"gradle":           "android",
	"ds_store":         "apple",
	"localized":        "apple",
	"flac":             "audio",
	"m4a":              "audio",
	"mp3":              "audio",
	"ogg":              "audio",
	"wav":              "audio",
	"editorconfig":     "conf",
	"scss":             "css",
	"docx":             "doc",
	"gdoc":             "doc",
	"dockerfile":       "docker",
	"mobi":             "ebook",
	"eot":              "font",
	"otf":              "font",
	"ttf":              "font",
	"woff":             "font",
	"woff2":            "font",
	"gitconfig":        "git",
	"gitignore":        "git",
	"gitignore_global": "git",
	"lhs":              "hs",
	"bmp":              "image",
	"gif":              "image",
	"ico":              "image",
	"jpeg":             "image",
	"jpg":              "image",
	"png":              "image",
	"pxm":              "image",
	"svg":              "image",
	"tiff":             "image",
	"webp":             "image",
	"jar":              "java",
	"properties":       "json",
	"cjs":              "js",
	"mjs":              "js",
	"tsx":              "jsx",
	"license":          "md",
	"markdown":         "md",
	"mkd":              "md",
	"rdoc":             "md",
	"readme":           "md",
	"gslides":          "ppt",
	"pptx":             "ppt",
	"ipynb":            "py",
	"pyc":              "py",
	"rdata":            "r",
	"rds":              "r",
	"gemfile":          "rb",
	"gemspec":          "rb",
	"guardfile":        "rb",
	"lock":             "rb",
	"procfile":         "rb",
	"rakefile":         "rb",
	"rspec":            "rb",
	"rspec_parallel":   "rb",
	"rspec_status":     "rb",
	"ru":               "rb",
	"erb":              "rubydoc",
	"slim":             "rubydoc",
	"bash":             "shell",
	"bash_history":     "shell",
	"bash_profile":     "shell",
	"bashrc":           "shell",
	"fish":             "shell",
	"sh":               "shell",
	"zsh":              "shell",
	"zsh-theme":        "shell",
	"zshrc":            "shell",
	"stylus":           "styl",
	"cls":              "tex",
	"avi":              "video",
	"flv":              "video",
	"mkv":              "video",
	"mov":              "video",
	"mp4":              "video",
	"ogv":              "video",
	"webm":             "video",
	"bat":              "windows",
	"exe":              "windows",
	"ini":              "windows",
	"csv":              "xls",
	"gsheet":           "xls",
	"xlsx":             "xls",
	"xul":              "xml",
	"yaml":             "yml",
	"toml":             "toml",
	"gz":               "zip",
	"rar":              "zip",
	"tar":              "zip",
	"mod":              "go",
	"sum":              "go",
}

var files = map[string]rune{
	"ai":           '\ue7b4',
	"android":      '\ue70e',
	"apple":        '\uf179',
	"audio":        '\ufc58',
	"avro":         '\ue60b',
	"c":            '\ue61e',
	"clj":          '\ue768',
	"coffee":       '\uf0f4',
	"conf":         '\ue615',
	"cpp":          '\ue61d',
	"css":          '\ue749',
	"d":            '\ue7af',
	"dart":         '\ue798',
	"db":           '\uf1c0',
	"diff":         '\uf440',
	"doc":          '\uf1c2',
	"docker":       '\uf308',
	"ebook":        '\ue28b',
	"env":          '\uf462',
	"epub":         '\ue28a',
	"erl":          '\ue7b1',
	"font":         '\uf031',
	"gform":        '\uf298',
	"git":          '\uf7a1',
	"go":           '\ue65e',
	"gruntfile.js": '\ue74c',
	"hs":           '\ue777',
	"html":         '\uf13b',
	"image":        '\uf1c5',
	"iml":          '\ue7b5',
	"java":         '\ue204',
	"js":           '\ue74e',
	"json":         '\ue60b',
	"jsx":          '\ue7ba',
	"less":         '\ue758',
	"log":          '\uf18d',
	"lua":          '\ue620',
	"md":           '\uf48a',
	"mustache":     '\ue60f',
	"npmignore":    '\ue71e',
	"pdf":          '\uf1c1',
	"php":          '\ue73d',
	"pl":           '\ue769',
	"ppt":          '\uf1c4',
	"psd":          '\ue7b8',
	"py":           '\ue606',
	"r":            '\uf25d',
	"rb":           '\ue21e',
	"rdb":          '\ue76d',
	"rss":          '\uf09e',
	"rubydoc":      '\ue73b',
	"sass":         '\ue603',
	"scala":        '\ue737',
	"shell":        '\uf489',
	"sqlite3":      '\ue7c4',
	"styl":         '\ue600',
	"tex":          '\ue600',
	"ts":           '\ue628',
	"twig":         '\ue61c',
	"txt":          '\uf15c',
	"video":        '\uf03d',
	"vim":          '\ue62b',
	"windows":      '\uf17a',
	"xls":          '\uf1c3',
	"xml":          '\ue619',
	"yarn.lock":    '\ue718',
	"yml":          '\uf481',
	"toml":         '\ue6b2',
	"zip":          '\uf410',
	"sql":          '\ue64d',
	"rs":           '\ue68b',
	"cargo.lock":   '\ue68b',
}
