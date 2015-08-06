package lowprofile

import (
	. "github.com/DualSpark/lowprofile/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/DualSpark/lowprofile/Godeps/_workspace/src/github.com/onsi/gomega"
	"flag"
	"github.com/DualSpark/lowprofile/Godeps/_workspace/src/github.com/codegangsta/cli"
	"io/ioutil"
	"os"
	"strings"
	"errors"
)

var _ = Describe("Activate", func() {
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
    It("should error", func() {
				os.Setenv("SHELL", "not_supported_shell")
				Expect(BeforeActivateProfile(context)).Should(Equal(errors.New("Sorry, not_supported_shell is not a supported shell")))
    })
  })

	Context("When the resource file doesn't exist", func() {
    It("should error", func() {
				os.Setenv("SHELL", "/bin/bash")
				os.Setenv("HOME", "/tmp")
        Expect(BeforeActivateProfile(context)).Should(Equal(errors.New("File ~/.bash_profile not found")))
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

    It("should activate", func() {
        Expect(func(){ActivateProfile(context)}).ShouldNot(Panic())
				contents, _ := ioutil.ReadFile(profilePath)
				Expect(strings.TrimSpace(string(contents))).To(Equal("export AWS_PROFILE=some-profile"))
    })

		It("should reactivate", func() {
				set.Set("profile", "some-other-profile")

				Expect(func(){ActivateProfile(context)}).ShouldNot(Panic())
				contents, _ := ioutil.ReadFile(profilePath)
				Expect(strings.TrimSpace(string(contents))).To(Equal("export AWS_PROFILE=some-other-profile"))
		})

		It("should reactivate after deactivation", func() {
				set.Set("profile", "some-other-profile")

				Expect(func(){ActivateProfile(context)}).ShouldNot(Panic())
				Expect(func(){DeactivateProfile(context)}).ShouldNot(Panic())
				Expect(func(){ActivateProfile(context)}).ShouldNot(Panic())
				contents, _ := ioutil.ReadFile(profilePath)
				Expect(strings.TrimSpace(string(contents))).To(Equal("export AWS_PROFILE=some-other-profile"))
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

		It("should activate", func() {
				Expect(func(){ActivateProfile(context)}).ShouldNot(Panic())
				contents, _ := ioutil.ReadFile(profilePath)
				Expect(strings.TrimSpace(string(contents))).To(Equal("export AWS_PROFILE=some-profile"))
		})

		It("should reactivate", func() {
				set.Set("profile", "some-other-profile")

				Expect(func(){ActivateProfile(context)}).ShouldNot(Panic())
				contents, _ := ioutil.ReadFile(profilePath)
				Expect(strings.TrimSpace(string(contents))).To(Equal("export AWS_PROFILE=some-other-profile"))
		})

		It("should reactivate after deactivation", func() {
				set.Set("profile", "some-other-profile")

				Expect(func(){ActivateProfile(context)}).ShouldNot(Panic())
				Expect(func(){DeactivateProfile(context)}).ShouldNot(Panic())
				Expect(func(){ActivateProfile(context)}).ShouldNot(Panic())
				contents, _ := ioutil.ReadFile(profilePath)
				Expect(strings.TrimSpace(string(contents))).To(Equal("export AWS_PROFILE=some-other-profile"))
		})

		AfterEach(func(){
			os.Remove(profilePath)
		})
	})

	AfterEach(func() {

  })
})
