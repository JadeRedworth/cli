package server

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/fnproject/cli/common"
	"github.com/go-yaml/yaml"
	"github.com/urfave/cli"
)

func start(c *cli.Context) error {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln("Getwd failed:", err)
	}
	args := []string{"run", "--rm", "-i",
		"--name", "fnserver",
		"-v", fmt.Sprintf("%s/data:/app/data", wd),
		"-v", "/var/run/docker.sock:/var/run/docker.sock",
		"--privileged",
		"-p", fmt.Sprintf("%d:8080", c.Int("port")),
		"--entrypoint", "./fnserver",
	}
	if c.String("log-level") != "" {
		args = append(args, "-e", fmt.Sprintf("FN_LOG_LEVEL=%v", c.String("log-level")))
	}
	if c.String("env-file") != "" {
		args = append(args, "--env-file", c.String("env-file"))
	}
	if c.Bool("detach") {
		args = append(args, "-d")
	}
	args = append(args, common.FunctionsDockerImage)
	cmd := exec.Command("docker", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		log.Fatalln("starting command failed:", err)
	}

	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	// catch ctrl-c and kill
	sigC := make(chan os.Signal, 2)
	signal.Notify(sigC, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-sigC:
			log.Println("interrupt caught, exiting")
			err = cmd.Process.Signal(syscall.SIGTERM)
			if err != nil {
				log.Println("error: could not kill process:", err)
				return err
			}
		case err := <-done:
			if err != nil {
				log.Println("error: processed finished with error", err)
			}
		}
		return err
	}
}

func update(c *cli.Context) error {
	args := []string{"pull",
		common.FunctionsDockerImage,
	}
	cmd := exec.Command("docker", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		log.Fatalln("starting command failed:", err)
	}

	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	// catch ctrl-c and kill
	sigC := make(chan os.Signal, 2)
	signal.Notify(sigC, os.Interrupt, syscall.SIGTERM)
	select {
	case <-sigC:
		log.Println("interrupt caught, exiting")
		err = cmd.Process.Kill()
		if err != nil {
			log.Println("error: could not kill process")
		}
	case err := <-done:
		if err != nil {
			log.Println("processed finished with error:", err)
		} else {
			log.Println("process finished gracefully")
		}
	}
	return nil
}

// steps:
// • Yaml file with extensions listed
// • NO‎TE: All extensions should use env vars for config
// • ‎Generating main.go with extensions
// * Generate a Dockerfile that gets all the extensions (using dep)
// • ‎then generate a main.go with extensions
// • ‎compile, throw in another container like main dockerfile
func (b *buildServerCmd) Build(c *cli.Context) error {

	if c.String("tag") == "" {
		return errors.New("docker tag required")
	}

	// path, err := os.Getwd()
	// if err != nil {
	// 	return err
	// }
	fpath := "ext.yaml"
	bb, err := ioutil.ReadFile(fpath)
	if err != nil {
		return fmt.Errorf("could not open %s for parsing. Error: %v", fpath, err)
	}
	ef := &extFile{}
	err = yaml.Unmarshal(bb, ef)
	if err != nil {
		return err
	}

	err = os.MkdirAll("tmp", 0777)
	if err != nil {
		return err
	}
	err = os.Chdir("tmp")
	if err != nil {
		return err
	}
	err = generateMain(ef)
	if err != nil {
		return err
	}
	err = generateDockerfile()
	if err != nil {
		return err
	}
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	err = common.RunBuild(c, dir, c.String("tag"), "Dockerfile", nil, b.noCache)
	if err != nil {
		return err
	}
	fmt.Printf("Custom Fn server built successfully.\n")
	return nil
}

func generateMain(ef *extFile) error {
	tmpl, err := template.New("main").Parse(mainTmpl)
	if err != nil {
		return err
	}
	f, err := os.Create("main.go")
	if err != nil {
		return err
	}
	defer f.Close()
	err = tmpl.Execute(f, ef)
	if err != nil {
		return err
	}
	return nil
}

func generateDockerfile() error {
	if err := ioutil.WriteFile("Dockerfile", []byte(dockerFileTmpl), os.FileMode(0644)); err != nil {
		return err
	}
	return nil
}

type extFile struct {
	Extensions []*extInfo `yaml:"extensions"`
}

type extInfo struct {
	Name string `yaml:"name"`
	// will have version and other things down the road
}

var mainTmpl = `package main

import (
	"context"

	"github.com/fnproject/fn/api/server"
	
	{{- range .Extensions }}
		_ "{{ .Name }}"
	{{- end}}
)

func main() {
	ctx := context.Background()
	funcServer := server.NewFromEnv(ctx)
	{{- range .Extensions }}
		funcServer.AddExtensionByName("{{ .Name }}")
	{{- end}}
	funcServer.Start(ctx)
}
`

// NOTE: Getting build errors with dep, probably because our vendor dir is wack. Might work again once we switch to dep.
// vendor/github.com/fnproject/fn/api/agent/drivers/docker/registry.go:93: too many arguments in call to client.NewRepository
// have ("context".Context, reference.Named, string, http.RoundTripper) want (reference.Named, string, http.RoundTripper)
// go build github.com/x/y/vendor/github.com/rdallman/migrate/database/mysql: no buildable Go source files in /go/src/github.com/x/y/vendor/github.com/rdallman/migrate/database/mysql
// # github.com/x/y/vendor/github.com/openzipkin/zipkin-go-opentracing/thrift/gen-go/scribe
// vendor/github.com/openzipkin/zipkin-go-opentracing/thrift/gen-go/scribe/scribe.go:210: undefined: thrift.TClient
var dockerFileTmpl = `# build stage
FROM golang:1.10-alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
# RUN go get -u github.com/golang/dep/cmd/dep
ENV D=/go/src/github.com/x/y
ADD main.go $D/
RUN cd $D && go get
# RUN cd $D && dep init && dep ensure
RUN cd $D && go build -o fnserver && cp fnserver /tmp/

# final stage
FROM fnproject/dind
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=build-env /tmp/fnserver /app/fnserver
CMD ["./fnserver"]
`
