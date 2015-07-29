package lowprofile

import (
	. "github.com/kellyp/lowprofile/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/kellyp/lowprofile/Godeps/_workspace/src/github.com/onsi/gomega"
	"flag"
	"github.com/kellyp/lowprofile/Godeps/_workspace/src/github.com/codegangsta/cli"
	"io/ioutil"
	"os"
)

var _ = Describe("Describe", func() {
	var (
		context *cli.Context
		set *flag.FlagSet
	)

	BeforeEach(func() {
		os.Clearenv()

		set = flag.NewFlagSet("test", 0)
		command := cli.Command{Name: "describe-active-profile"}
		context = cli.NewContext(nil, set, nil)
		context.Command = command
  })


	Context("When the resource file doesn't exist", func() {
    It("should panic", func() {
				os.Setenv("HOME", "/tmp")
        Expect(func(){DescribeProfiles(context)}).Should(Panic())
    })
  })

	Context("When the aws resource file exists", func() {
		var (
			homePath string = "/tmp"
			awsPath string = homePath + "/.aws"
			resourcePath string = awsPath + "/credentials"
			shellName string = "/bin/bash"
		)

		BeforeEach(func(){
			os.Setenv("SHELL", shellName)
			os.Setenv("HOME", homePath)

			os.Mkdir(awsPath, 0777)
			var bytes []byte
			ioutil.WriteFile(resourcePath, bytes, 0777)
		})

		It("should describe profiles", func() {
				Expect(func(){DescribeProfiles(context)}).ShouldNot(Panic())
		})

		It("should describe active profiles", func() {
				os.Setenv("AWS_DEFAULT_PROFILE", "some-profile")
				Expect(func(){DescribeActiveProfile(context)}).ShouldNot(Panic())
				os.Setenv("AWS_DEFAULT_PROFILE", "some-other-profile")
				Expect(func(){DescribeActiveProfile(context)}).ShouldNot(Panic())
				os.Setenv("AWS_DEFAULT_PROFILE", "")
				Expect(func(){DescribeActiveProfile(context)}).ShouldNot(Panic())
		})

		AfterEach(func(){
			os.RemoveAll(awsPath)
		})
  })

	AfterEach(func() {

  })
})
