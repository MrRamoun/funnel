package gridengine

import (
	"github.com/ohsu-comp-bio/funnel/compute"
	"github.com/ohsu-comp-bio/funnel/config"
	"github.com/ohsu-comp-bio/funnel/events"
	"github.com/ohsu-comp-bio/funnel/proto/tes"
	"regexp"
)

// NewBackend returns a new Grid Engine HPCBackend instance.
func NewBackend(conf config.Config, reader tes.ReadOnlyServer, writer events.Writer) *compute.HPCBackend {
	return &compute.HPCBackend{
		Name:      "gridengine",
		SubmitCmd: "qsub",
		CancelCmd: "qdel",
		Conf:      conf,
		Template:  conf.GridEngine.Template,
		Event:     writer,
		Database:  reader,
		ExtractID: extractID,
	}
}

// extractID extracts the task id from the response returned by the `qsub` command.
// Example response:
// Your job 1 ("test_job") has been submitted
func extractID(in string) string {
	re := regexp.MustCompile("(Your job )([0-9]+)( \\(\".*\"\\) has been submitted)\n$")
	return re.ReplaceAllString(in, "$2")
}
