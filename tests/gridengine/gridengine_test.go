package e2e

import (
	"context"
	"github.com/ohsu-comp-bio/funnel/logger"
	"github.com/ohsu-comp-bio/funnel/proto/tes"
	"github.com/ohsu-comp-bio/funnel/tests"
	"os"
	"testing"
)

var fun *tests.Funnel

func TestMain(m *testing.M) {
	conf := tests.DefaultConfig()

	if conf.Compute != "gridengine" {
		logger.Debug("Skipping grid engine e2e tests...")
		os.Exit(0)
	}

	fun = tests.NewFunnel(conf)
	fun.StartServerInDocker("ohsucompbio/gridengine:latest", []string{})
	defer fun.CleanupTestServerContainer()

	m.Run()
	return
}

func TestHelloWorld(t *testing.T) {
	id := fun.Run(`
    --sh 'echo hello world'
  `)
	task := fun.Wait(id)

	if task.State != tes.State_COMPLETE {
		t.Fatal("expected task to complete")
	}

	if task.Logs[0].Logs[0].Stdout != "hello world\n" {
		t.Fatal("Missing stdout")
	}
}

func TestResourceRequest(t *testing.T) {
	id := fun.Run(`
    --sh 'echo I need resources!' --cpu 1 --ram 2 --disk 5
  `)
	task := fun.Wait(id)

	if task.State != tes.State_COMPLETE {
		t.Fatal("expected task to complete")
	}

	if task.Logs[0].Logs[0].Stdout != "I need resources!\n" {
		t.Fatal("Missing stdout")
	}
}

func TestCancel(t *testing.T) {
	id := fun.Run(`
    --sh 'echo I wont ever run!' --cpu 1000
  `)

	_, err := fun.HTTP.CancelTask(context.Background(), &tes.CancelTaskRequest{Id: id})
	if err != nil {
		t.Fatal("unexpected error")
	}

	task := fun.Wait(id)

	if task.State != tes.State_CANCELED {
		t.Fatal("expected task to get canceled")
	}
}
