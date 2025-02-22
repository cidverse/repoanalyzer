package analyzerapi

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/cidverse/cidverseutils/filesystem"
	"github.com/gosimple/slug"
	ignore "github.com/sabhiram/go-gitignore"
	"golang.org/x/exp/slog"
)

func GetAnalyzerContext(projectDir string) AnalyzerContext {
	// respect gitignore
	ignoreMatcher := ProcessIgnoreFiles([]string{filepath.Join(projectDir, ".gitignore"), filepath.Join(projectDir, ".cidignore")}, []string{".git", ".svn", ".hg"})

	files, err := filesystem.FindFiles(projectDir, func(absPath string, name string) bool {
		return ignoreMatcher.MatchesPath(absPath)
	}, func(absPath string, name string) bool {
		return true
	})
	if err != nil {
		slog.Error("failed to get directory list", err, slog.String("path", projectDir))
	}
	filesByExtension := filesystem.GenerateFileMapByDeepExtension(files)

	// sorting
	sort.Slice(files, func(i, j int) bool {
		return len(files[i]) < len(files[j])
	})

	// result
	return AnalyzerContext{
		ProjectDir:       projectDir,
		Files:            files,
		FilesByExtension: filesByExtension,
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

func GetSingleLanguageMap(language ProjectLanguage, version string) map[ProjectLanguage]string {
	languageMap := make(map[ProjectLanguage]string)
	languageMap[language] = version
	return languageMap
}

func AddModuleToResult(result *[]*ProjectModule, module *ProjectModule) {
	parent := FindParentModule(result, module)
	if parent != nil {
		module.Name = parent.Name + "-" + module.Name
		module.Slug = parent.Slug + "-" + module.Slug

		// inherit language from parent if not set
		if module.Language == nil || len(module.Language) == 0 {
			module.Language = parent.Language
		}

		parent.Submodules = append(parent.Submodules, module)
	} else {
		*result = append(*result, module)
	}
}

func ProcessIgnoreFiles(files []string, ignoreLines []string) *ignore.GitIgnore {
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

func GetSlugFromPath(projectDir, file, analyzerName string) string {
	relativePath, err := filepath.Rel(projectDir, file)
	if err != nil {
		return slug.Make(analyzerName + "-" + file)
	}
	return slug.Make(analyzerName + "-" + relativePath)
}
