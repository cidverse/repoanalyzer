package analyzerapi

import (
	"github.com/rs/zerolog/log"
	ignore "github.com/sabhiram/go-gitignore"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func GetAnalyzerContext(projectDir string) AnalyzerContext {
	// respect gitignore
	ignoreMatcher := ProcessIgnoreFiles([]string{filepath.Join(projectDir, ".gitignore"), filepath.Join(projectDir, ".cidignore")})

	// files
	var files []string
	filesByExtension := make(map[string][]string)
	var filesWithoutExtension []string

	err := filepath.WalkDir(projectDir, func(path string, d fs.DirEntry, err error) error {
		// check for directory skip
		if err != nil {
			log.Warn().Err(err).Str("path", projectDir).Msg("output")
			return nil
		}

		if d.IsDir() {
			if d.Name() == ".git" || ignoreMatcher.MatchesPath(path) {
				return filepath.SkipDir
			}
		} else {
			if ignoreMatcher.MatchesPath(path) {
				return nil
			}

			files = append(files, path)
			splitByExt := strings.SplitN(d.Name(), ".", 2)
			if len(splitByExt) == 2 {
				filesByExtension[splitByExt[1]] = append(filesByExtension[splitByExt[1]], path)
			} else {
				filesWithoutExtension = append(filesWithoutExtension, path)
			}
		}

		return nil
	})
	if err != nil {
		log.Fatal().Err(err).Str("path", projectDir).Msg("failed to get directory list")
	}

	// sorting
	sort.Slice(files, func(i, j int) bool {
		return len(files[i]) < len(files[j])
	})

	// result
	return AnalyzerContext{
		ProjectDir:            projectDir,
		Files:                 files,
		FilesByExtension:      filesByExtension,
		FilesWithoutExtension: filesWithoutExtension,
	}
}

func FindParentModule(modules *[]*ProjectModule, module *ProjectModule) *ProjectModule {
	for _, m := range *modules {
		if strings.HasPrefix(module.Directory, m.Directory+"/") || strings.HasPrefix(module.Directory, m.Directory+"\\") {
			return m
		}
	}

	return nil
}

func GetSingleLanguageMap(language ProjectLanguage, version *string) map[ProjectLanguage]*string {
	languageMap := make(map[ProjectLanguage]*string)
	languageMap[language] = version
	return languageMap
}

func AddModuleToResult(result *[]*ProjectModule, module *ProjectModule) {
	parent := FindParentModule(result, module)
	if parent != nil {
		module.Name = parent.Name + "-" + module.Name
		module.Slug = parent.Slug + "-" + module.Slug
		parent.Submodules = append(parent.Submodules, module)
	} else {
		*result = append(*result, module)
	}
}

func ProcessIgnoreFiles(files []string) *ignore.GitIgnore {
	var ignoreLines []string

	for _, file := range files {
		if _, err := os.Stat(file); err == nil {
			content, contentErr := os.ReadFile(file)
			if contentErr == nil {
				for _, l := range strings.Split(string(content), "\n") {
					ignoreLines = append(ignoreLines, l)
				}
			}
		}
	}

	return ignore.CompileIgnoreLines(ignoreLines...)
}
