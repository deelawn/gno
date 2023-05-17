// Mostly copied from go source at tip, commit d922c0a.
//
// Copyright 2015 The Go Authors. All rights reserved.

package doc

import (
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// A bfsDir describes a directory holding code by specifying
// the expected import path and the file system directory.
type bfsDir struct {
	importPath string // import path for that dir
	dir        string // file system directory
}

// dirs is a structure for scanning the directory tree.
// Its Next method returns the next Go source directory it finds.
// Although it can be used to scan the tree multiple times, it
// only walks the tree once, caching the data it finds.
type bfsDirs struct {
	scan   chan bfsDir // Directories generated by walk.
	hist   []bfsDir    // History of reported Dirs.
	offset int         // Counter for Next.
}

// newDirs begins scanning the given stdlibs directory.
func newDirs(dirs ...string) *bfsDirs {
	d := &bfsDirs{
		hist: make([]bfsDir, 0, 256),
		scan: make(chan bfsDir),
	}
	go d.walk(dirs)
	return d
}

// Reset puts the scan back at the beginning.
func (d *bfsDirs) Reset() {
	d.offset = 0
}

// Next returns the next directory in the scan. The boolean
// is false when the scan is done.
func (d *bfsDirs) Next() (bfsDir, bool) {
	if d.offset < len(d.hist) {
		dir := d.hist[d.offset]
		d.offset++
		return dir, true
	}
	dir, ok := <-d.scan
	if !ok {
		return bfsDir{}, false
	}
	d.hist = append(d.hist, dir)
	d.offset++
	return dir, ok
}

// walk walks the trees in the given roots.
func (d *bfsDirs) walk(roots []string) {
	for _, root := range roots {
		d.bfsWalkRoot(root)
	}
	close(d.scan)
}

// bfsWalkRoot walks a single directory hierarchy in breadth-first lexical order.
// Each Go source directory it finds is delivered on d.scan.
func (d *bfsDirs) bfsWalkRoot(root string) {
	root = filepath.Clean(root)

	// this is the queue of directories to examine in this pass.
	this := []string{}
	// next is the queue of directories to examine in the next pass.
	next := []string{root}

	for len(next) > 0 {
		this, next = next, this[:0]
		for _, dir := range this {
			fd, err := os.Open(dir)
			if err != nil {
				log.Print(err)
				continue
			}
			entries, err := fd.Readdir(0)
			fd.Close()
			if err != nil {
				log.Print(err)
				continue
			}
			hasGnoFiles := false
			for _, entry := range entries {
				name := entry.Name()
				// For plain files, remember if this directory contains any .gno
				// source files, but ignore them otherwise.
				if !entry.IsDir() {
					if !hasGnoFiles && strings.HasSuffix(name, ".gno") {
						hasGnoFiles = true
					}
					continue
				}
				// Entry is a directory.

				// Ignore same directories ignored by the go tool.
				if name[0] == '.' || name[0] == '_' || name == "testdata" {
					continue
				}
				// Remember this (fully qualified) directory for the next pass.
				next = append(next, filepath.Join(dir, name))
			}
			if hasGnoFiles {
				// It's a candidate.
				var importPath string
				if len(dir) > len(root) {
					importPath = filepath.ToSlash(dir[len(root)+1:])
				}
				d.scan <- bfsDir{importPath, dir}
			}
		}
	}
}

// findPackage finds a package iterating over d where the import path has
// name as a suffix (which may be a package name or a fully-qualified path).
// returns a list of possible directories. If a directory's import path matched
// exactly, it will be returned as first.
func (d *bfsDirs) findPackage(name string) []bfsDir {
	d.Reset()
	candidates := make([]bfsDir, 0, 4)
	for dir, ok := d.Next(); ok; dir, ok = d.Next() {
		// want either exact matches or suffixes
		if dir.importPath == name || strings.HasSuffix(dir.importPath, "/"+name) {
			candidates = append(candidates, dir)
		}
	}
	sort.Slice(candidates, func(i, j int) bool {
		// prefer exact matches with name
		if candidates[i].importPath == name {
			return true
		} else if candidates[j].importPath == name {
			return false
		}
		return candidates[i].importPath < candidates[j].importPath
	})
	return candidates
}

// findDir determines if the given absdir is present in the Dirs.
// If not, the nil slice is returned. It returns always at most one dir.
func (d *bfsDirs) findDir(absdir string) []bfsDir {
	d.Reset()
	for dir, ok := d.Next(); ok; dir, ok = d.Next() {
		if dir.dir == absdir {
			return []bfsDir{dir}
		}
	}
	return nil
}