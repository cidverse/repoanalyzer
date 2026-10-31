package analyzerapi

import (
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/charlievieth/fastwalk"
	"github.com/cidverse/cidverseutils/filesystem"
	"github.com/gosimple/slug"
	"golang.org/x/exp/slog"
)

func GetAnalyzerContext(projectDir string) AnalyzerContext {
	// respect gitignore, including nested ignore files
	ignoreMatcher := ProcessIgnoreFiles(projectDir)

	files, err := findFiles(projectDir, func(absPath string, isDir bool) bool {
		return ignoreMatcher.MatchesPath(absPath, isDir)
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

func findFiles(rootPath string, ignore func(absPath string, isDir bool) bool) ([]string, error) {
	var (
		mu    sync.Mutex
		files []string
	)

	conf := fastwalk.Config{
		Follow: false,
	}
	err := fastwalk.Walk(&conf, rootPath, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if ignore(path, info.IsDir()) {
			if info.IsDir() {
				return filepath.SkipDir
			}

			return nil
		}

		if !info.IsDir() {
			mu.Lock()
			files = append(files, path)
			mu.Unlock()
		}
		return nil
	})

	sort.Strings(files)

	return files, err
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
	// ensure dependencies are ordered
	if module.Dependencies != nil {
		sort.Slice(module.Dependencies, func(i, j int) bool {
			return module.Dependencies[i].ID < module.Dependencies[j].ID
		})
	}

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

func GetSlugFromPath(projectDir, file, analyzerName string) string {
	relativePath, err := filepath.Rel(projectDir, file)
	if err != nil {
		return slug.Make(analyzerName + "-" + file)
	}
	return slug.Make(analyzerName + "-" + relativePath)
}
