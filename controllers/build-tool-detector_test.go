package controllers_test

import (
	"build-tool-detector/app/test"
	controllers "build-tool-detector/controllers"
	"github.com/goadesign/goa"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = Describe("BuildToolDetector", func() {
	Context("Internal Server Error", func() {
		It("Non-existent owner name -- 500 Internal Server Error", func() {
			service := goa.New("build-tool-detector")
			test.ShowBuildToolDetectorInternalServerError(GinkgoT(), nil, nil, controllers.NewBuildToolDetectorController(service), "https://github.com/fabric8-launcherz/launcher-backend", nil)
		})

		It("Non-existent repository name -- 500 Internal Server Error", func() {
			service := goa.New("build-tool-detector")
			test.ShowBuildToolDetectorInternalServerError(GinkgoT(), nil, nil, controllers.NewBuildToolDetectorController(service), "https://github.com/fabric8-launcher/launcher-backendz", nil)
		})

		It("Non-existent branch name -- 500 Internal Server Error", func() {
			service := goa.New("build-tool-detector")
			branch := "masterz"
			test.ShowBuildToolDetectorInternalServerError(GinkgoT(), nil, nil, controllers.NewBuildToolDetectorController(service), "https://github.com/fabric8-launcher/launcher-backend", &branch)
		})

		It("Build tool type expected to be Unknown -- 500 Internal Server Error", func() {
			service := goa.New("build-tool-detector")
			branch := "master"
			test.ShowBuildToolDetectorInternalServerError(GinkgoT(), nil, nil, controllers.NewBuildToolDetectorController(service), "https://github.com/fabric8-services/fabric8-wit", &branch)
		})
	})

	Context("Okay", func() {
		It("Okay response -- 200 Okay", func() {
			service := goa.New("build-tool-detector")
			branch := "master"
			test.ShowBuildToolDetectorOK(GinkgoT(), nil, nil, controllers.NewBuildToolDetectorController(service), "https://github.com/fabric8-launcher/launcher-backend", &branch)
		})

		It("Non-nil response -- 200 Okay", func() {
			service := goa.New("build-tool-detector")
			branch := "master"
			_, buildTool := test.ShowBuildToolDetectorOK(GinkgoT(), nil, nil, controllers.NewBuildToolDetectorController(service), "https://github.com/fabric8-launcher/launcher-backend", &branch)
			gomega.Expect(buildTool).ShouldNot(gomega.BeNil(), "buildTool should not be empty")
		})

		It("Build tool type to be Maven -- 200 Okay", func() {
			service := goa.New("build-tool-detector")
			branch := "master"
			_, buildTool := test.ShowBuildToolDetectorOK(GinkgoT(), nil, nil, controllers.NewBuildToolDetectorController(service), "https://github.com/fabric8-launcher/launcher-backend", &branch)
			gomega.Expect(buildTool.BuildToolType).Should(gomega.BeEquivalentTo("maven"), "build type should be equivalent to 'Maven'")
		})
	})
})