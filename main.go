package main

import (
	"log"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type resume struct {
	app.Compo
	isAppInstallable bool
}

func (r *resume) OnMount(ctx app.Context) {
	r.isAppInstallable = ctx.IsAppInstallable()
}
func (r *resume) onInstallButtonClicked(ctx app.Context, e app.Event) {
	ctx.ShowAppInstallPrompt()
}
func (r *resume) OnAppInstallChange(ctx app.Context) {
	r.isAppInstallable = ctx.IsAppInstallable()
}

func (r *resume) Render() app.UI {
	return app.Div().Body(
		app.If(r.isAppInstallable, func() app.UI {
			return app.Button().
				Text("Install App").
				OnClick(r.onInstallButtonClicked)
		}),
		app.H1().Text("Brandon Carpenter"),
		app.P().Text("City, ST | (555) 555-5555 | email@email.com"),
		app.A().Href("https://www.linkedin.com/in/brandon-c-40380ab8/").Text("LinkedIn"),
		app.H2().Text("Skills"),
		app.Ul().Body(
			app.Li().Text("Programming languages (Bash, Groovy, Ruby, Go, Python)"),
			app.Li().Text("Build tools (Fastlane, Xcode, Gradle)"),
			app.Li().Text("Cloud (GCP, Azure, AWS)"),
			app.Li().Text("IaC (Terraform, Chef)"),
			app.Li().Text("CI/CD (Jenkins, Azure DevOps, GHA)"),
			app.Li().Text("Unix, Linux, macOS, Windows"),
			app.Li().Text("Containers (Kubernetes, Docker)"),
		),
		app.H2().Text("Experience"),
		app.H3().Text("Jack Henry (2/18-Present)"),
		app.H4().Text("Mobile Release Engineer"),
		app.P().Text("Staff Engineer/Manager(6/22 - Present), Senior(6/20 - 5/22)"),
		app.Ul().Body(
			app.Li().Text("Maintain and improve automated build/deploy pipeline in CI/CD environments for both native Android and iOS applications using Fastlane to ensure continued scalability and stability."),
			app.Li().Text("Troubleshoot and triage mobile CI/CD build/deploy issues."),
			app.Li().Text("Build and deploy iOS and Android monthly release for 1000+ white-label financial apps and hotfixes as needed, totaling more than 70,000 production releases each to both Apple App Store and Google Play over 7 years."),
			app.Li().Text("Manage 27 developer accounts housing a total of 1000+ client developer accounts."),
			app.Li().Text("Configure initial applications and app stores for new clients. Personally responsible for 600+ initial app configurations."),
			app.Li().Text("Automated the creation of first time app deployments for new clients taking initial setup time from 30+ hours to 4 hours."),
			app.Li().Text("Research and implement new integrations along with client engineers to deliver new features and products."),
			app.Li().Text("Work directly with engineering, support, and implementation teams to identify deficiencies and harden the mobile build/deploy process."),
			app.Li().Text("Mentored a small agile team to ensure applications are delivered consistently and on time."),
		),
		app.H4().Text("Devops Engineer"),
		app.P().Text("Advanced(6/18 - 5/20), Intern(2/18-5/18)"),
		app.Ul().Body(
			app.Li().Text("Maintained production CI/CD environments as well as lower environments."),
			app.Li().Text("Deployed and rolled back production level code across the company's entire technology stack."),
			app.Li().Text("Maintained, hardened, and created new reverse-proxy configs."),
			app.Li().Text("Designed automations for the deployment of mobile client applications using Fastlane and custom automations dropping the time between code release to production deployment from several days to several hours."),
			app.Li().Text("Setup and install open source and custom infrastructure tooling."),
		),
		app.H2().Text("Education"),
		app.H3().Text("University of Evansville"),
		app.P().Text("BS, Sociology"),
	)
}

func main() {

	app.Route("/", func() app.Composer { return &resume{} })

	app.RunWhenOnBrowser()

	err := app.GenerateStaticWebsite(".", &app.Handler{
		Name:        "Brandon Carpenter - Resume",
		Description: "A simple PWA version of my resume",
		Resources:   app.GitHubPages("pwa-resume"),
		Icon: app.Icon{
			Default: "web/icon.png",
			SVG:     "web/icon.svg",
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
