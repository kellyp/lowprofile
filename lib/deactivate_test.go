package lowprofile

import (
	. "github.com/dualspark/lowprofile/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/dualspark/lowprofile/Godeps/_workspace/src/github.com/onsi/gomega"
	"flag"
	"github.com/dualspark/lowprofile/Godeps/_workspace/src/github.com/codegangsta/cli"
	"io/ioutil"
	"os"
	"strings"
)

var _ = Describe("Deactivate", func() {
	var (
		context *cli.Context
		set *flag.FlagSet
	)

	BeforeEach(func() {
		os.Clearenv()

		set = flag.NewFlagSet("test", 0)
		set.String("profile", "profile-name", "doc")
		command := cli.Command{Name: "ap"}
		context = cli.NewContext(nil, set, nil)
		context.Command = command

		set.Set("profile", "some-profile")
  })


	Context("When the shell is not supported", func() {
    It("should panic", func() {
				os.Setenv("SHELL", "not_supported_shell")
        Expect(func(){DeactivateProfile(context)}).Should(Panic())
    })
  })

	Context("When the resource file doesn't exist", func() {
    It("should panic", func() {
				os.Setenv("SHELL", "bash")
				os.Setenv("HOME", "/tmp")
        Expect(func(){DeactivateProfile(context)}).Should(Panic())
    })
  })

	Context("When the bash resource file exists", func() {
		var (
			profilePath string = "/tmp/.bash_profile"
			shellName string = "/bin/bash"
			homePath string = "/tmp"
		)

		BeforeEach(func(){
			os.Setenv("SHELL", shellName)
			os.Setenv("HOME", homePath)

			var bytes []byte
			ioutil.WriteFile(profilePath, bytes, 0660)
		})

		It("should noop when not activated", func() {
				Expect(func(){DeactivateProfile(context)}).ShouldNot(Panic())
				contents, _ := ioutil.ReadFile(profilePath)

				Expect(strings.TrimSpace(string(contents))).To(Equal(""))
		})

		It("should deactivate when activated", func() {
				Expect(func(){ActivateProfile(context)}).ShouldNot(Panic())
				Expect(func(){DeactivateProfile(context)}).ShouldNot(Panic())
				contents, _ := ioutil.ReadFile(profilePath)

				Expect(strings.TrimSpace(string(contents))).To(Equal("# export AWS_DEFAULT_PROFILE=some-profile"))
		})

		AfterEach(func(){
			os.Remove(profilePath)
		})
  })

	Context("When the zsh resource file exists", func() {
		var (
			profilePath string = "/tmp/.zshrc"
			shellName string = "/bin/zsh"
			homePath string = "/tmp"
		)

		BeforeEach(func(){
			os.Setenv("SHELL", shellName)
			os.Setenv("HOME", homePath)

			var bytes []byte
			ioutil.WriteFile(profilePath, bytes, 0660)
		})

		It("should noop when not activated", func() {
				Expect(func(){DeactivateProfile(context)}).ShouldNot(Panic())
				contents, _ := ioutil.ReadFile(profilePath)

				Expect(strings.TrimSpace(string(contents))).To(Equal(""))
		})

		It("should deactivate when activated", func() {
				Expect(func(){ActivateProfile(context)}).ShouldNot(Panic())
				Expect(func(){DeactivateProfile(context)}).ShouldNot(Panic())
				contents, _ := ioutil.ReadFile(profilePath)

				Expect(strings.TrimSpace(string(contents))).To(Equal("# export AWS_DEFAULT_PROFILE=some-profile"))
		})

		AfterEach(func(){
			os.Remove(profilePath)
		})
	})

	AfterEach(func() {

  })
})
