package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/opensourceways/server-common-lib/logrusutil"
	liboptions "github.com/opensourceways/server-common-lib/options"
	"github.com/sirupsen/logrus"

	"github.com/opensourceways/software-package-github-server/config"
	"github.com/opensourceways/software-package-github-server/message-server"
	"github.com/opensourceways/software-package-github-server/mq"
	"github.com/opensourceways/software-package-github-server/softwarepkg/app"
	"github.com/opensourceways/software-package-github-server/softwarepkg/infrastructure/codeimpl"
	"github.com/opensourceways/software-package-github-server/softwarepkg/infrastructure/messageimpl"
)

type options struct {
	service liboptions.ServiceOptions
}

func (o *options) Validate() error {
	return o.service.Validate()
}

func gatherOptions(fs *flag.FlagSet, args ...string) options {
	var o options

	o.service.AddFlags(fs)

	fs.Parse(args)
	return o
}

func main() {
	logrusutil.ComponentInit("software-package")
	log := logrus.NewEntry(logrus.StandardLogger())

	o := gatherOptions(flag.NewFlagSet(os.Args[0], flag.ExitOnError), os.Args[1:]...)
	if err := o.Validate(); err != nil {
		logrus.Fatalf("Invalid options: %v", err)
	}

	cfg, err := config.LoadConfig(o.service.ConfigFile)
	if err != nil {
		logrus.Fatalf("load config file failed: %v", err)
	}

	if err = mq.Init(&cfg.MQ, log); err != nil {
		logrus.Fatalf("initialize mq failed, err:%v", err)
	}

	defer mq.Exit()

	pkgService := app.NewPkgService(
		codeimpl.NewCodeImpl(cfg.Code),
		messageimpl.NewMessageImpl(cfg.MessageServer.Message),
	)

	ms := messageserver.Init(pkgService, cfg.MessageServer)

	run(ms)
}

func run(ms *messageserver.MessageServer) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	var wg sync.WaitGroup
	defer wg.Wait()

	called := false
	ctx, done := context.WithCancel(context.Background())

	defer func() {
		if !called {
			called = true
			done()
		}
	}()

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()

		select {
		case <-ctx.Done():
			logrus.Info("receive done. exit normally")

			return

		case <-sig:
			logrus.Info("receive exit signal")
			done()
			called = true

			return
		}
	}(ctx)

	if err := ms.Run(ctx); err != nil {
		logrus.Errorf("run message server, err:%v", err)
	}
}
