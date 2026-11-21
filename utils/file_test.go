package utils_test

import (
	"io/fs"
	"os"
	"path/filepath"

	Utils "ckg/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("File", func() {
	Describe("Contains", func() {
		It("returns true when a DirEntry matches the name", func() {
			dir := GinkgoT().TempDir()
			gitPath := filepath.Join(dir, ".git")
			Expect(os.Mkdir(gitPath, 0o755)).To(Succeed())
			entries, err := os.ReadDir(dir)
			Expect(err).NotTo(HaveOccurred())
			Expect(Utils.Contains(entries, ".git")).To(BeTrue())
		})

		It("returns false when no DirEntry matches", func() {
			entries := []fs.DirEntry{}
			Expect(Utils.Contains(entries, ".git")).To(BeFalse())
		})
	})

	Describe("LocalizeFolder", func() {
		var root string

		BeforeEach(func() {
			root = GinkgoT().TempDir()
		})

		It("returns the path when folder exists as absolute or direct child", func() {
			target := filepath.Join(root, "myproject")
			Expect(os.Mkdir(target, 0o755)).To(Succeed())
			Expect(Utils.LocalizeFolder("myproject", root, 0)).To(Equal(root))
		})

		It("finds a nested folder within depth limit", func() {
			nested := filepath.Join(root, "work", "repo")
			Expect(os.MkdirAll(nested, 0o755)).To(Succeed())
			Expect(os.Mkdir(filepath.Join(nested, ".git"), 0o755)).To(Succeed())
			found := Utils.LocalizeFolder("repo", root, 0)
			Expect(found).NotTo(BeEmpty())
			Expect(filepath.Join(found, "repo")).To(BeADirectory())
		})

		It("skips hidden directories when searching", func() {
			hidden := filepath.Join(root, ".hidden", "repo")
			Expect(os.MkdirAll(hidden, 0o755)).To(Succeed())
			Expect(Utils.LocalizeFolder("repo", root, 0)).To(Equal(""))
		})

		It("returns empty string when level exceeds 2", func() {
			deep := filepath.Join(root, "a", "b", "c", "target")
			Expect(os.MkdirAll(deep, 0o755)).To(Succeed())
			Expect(Utils.LocalizeFolder("target", root, 3)).To(Equal(""))
		})

		It("returns folder path when it exists on disk", func() {
			existing := filepath.Join(root, "exists")
			Expect(os.Mkdir(existing, 0o755)).To(Succeed())
			Expect(Utils.LocalizeFolder(existing, root, 0)).To(Equal(existing))
		})
	})
})
