package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"
	"github.com/spf13/pflag"
	k8s_client "k8s.io/kubernetes/pkg/client/unversioned"
	kubectl_util "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

var (
	flags       = pflag.NewFlagSet("", pflag.ExitOnError)
	showVersion = flags.Bool(
		"version", false,
		"Print version information",
	)
	listenAddress = flags.String(
		"web.listen-address", ":9102",
		"Address to listen on for web interface and telemetry.",
	)
	metricPath = flags.String(
		"web.telemetry-path", "/metrics",
		"Path under which to expose metrics.",
	)
	inCluster = flags.Bool(
		"running-in-cluster", true,
		`Optional, if this controller is running in a kubernetes cluster, use the
		pod secrets for creating a Kubernetes client.`,
	)
)

var landingPage = []byte(`<html>
<head><title>Hierarchy exporter</title></head>
<body>
<h1>Hierarchy exporter</h1>
<p><a href='` + *metricPath + `'>Metrics</a></p>
</body>
</html>
`)

func main() {
	flags.AddGoFlagSet(flag.CommandLine)
	flags.Parse(os.Args)
	var client *k8s_client.Client
	clientConfig := kubectl_util.DefaultClientConfig(flags)

	// Workaround of noisy log, see https://github.com/kubernetes/kubernetes/issues/17162
	flag.CommandLine.Parse([]string{})

	if *showVersion {
		fmt.Fprintln(os.Stdout, version.Print("hierarchy_exporter"))
		os.Exit(0)
	}
	var err error
	if *inCluster {
		client, err = k8s_client.NewInCluster()
	} else {
		config, connErr := clientConfig.ClientConfig()
		if connErr != nil {
			log.Fatalln("error connecting to the client:", err)
		}
		client, err = k8s_client.New(config)
	}
	if err != nil {
		log.Fatalln("failed to create client:", err)
	}
	controller, err := NewController(client)
	if err != nil {
		log.Fatalln("error create hierarchy store:", err)
	}
	go controller.Run()
	exporter := NewExporter(controller)
	prometheus.MustRegister(exporter)
	log.Infoln("Starting hierarchy_exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())

	http.Handle(*metricPath, prometheus.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(landingPage)
	})

	log.Infoln("Listening on", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
