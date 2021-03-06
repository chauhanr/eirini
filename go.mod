module code.cloudfoundry.org/eirini

go 1.15

replace (
	github.com/coreos/bbolt => github.com/coreos/bbolt v1.3.0
	google.golang.org/grpc => google.golang.org/grpc v1.29.1
	k8s.io/client-go => k8s.io/client-go v0.19.2
)

require (
	cloud.google.com/go v0.70.0 // indirect
	code.cloudfoundry.org/bbs v0.0.0-20200615191359-7b6fa295fa8d // indirect
	code.cloudfoundry.org/cfhttp/v2 v2.0.0
	code.cloudfoundry.org/clock v1.0.0 // indirect
	code.cloudfoundry.org/consuladapter v0.0.0-20200131002136-ac1daf48ba97 // indirect
	code.cloudfoundry.org/diego-logging-client v0.0.0-20200130234554-60ef08820a45 // indirect
	code.cloudfoundry.org/eirinix v0.4.0
	code.cloudfoundry.org/executor v0.0.0-20200629205945-23d8d6f82636 // indirect
	code.cloudfoundry.org/garden v0.0.0-20200813151451-b404ff2d61e6 // indirect
	code.cloudfoundry.org/go-diodes v0.0.0-20190809170250-f77fb823c7ee // indirect
	code.cloudfoundry.org/go-loggregator v7.4.0+incompatible
	code.cloudfoundry.org/lager v2.0.0+incompatible
	code.cloudfoundry.org/locket v0.0.0-20200509160055-68bb3033b039 // indirect
	code.cloudfoundry.org/rep v0.0.0-20200325195957-1404b978e31e // indirect
	code.cloudfoundry.org/rfc5424 v0.0.0-20180905210152-236a6d29298a // indirect
	code.cloudfoundry.org/runtimeschema v0.0.0-20180622184205-c38d8be9f68c
	code.cloudfoundry.org/tlsconfig v0.0.0-20200131000646-bbe0f8da39b3
	code.cloudfoundry.org/tps v0.0.0-20190724214151-ce1ef3913d8e
	code.cloudfoundry.org/urljoiner v0.0.0-20170223060717-5cabba6c0a50 // indirect
	github.com/Azure/go-autorest/autorest v0.11.10 // indirect
	github.com/cloudflare/cfssl v1.5.0 // indirect
	github.com/cockroachdb/apd v1.1.0 // indirect
	github.com/containerd/continuity v0.0.0-20200928162600-f2cc35102c2a // indirect
	github.com/containers/image v3.0.2+incompatible
	github.com/containers/storage v1.23.5 // indirect
	github.com/deckarep/golang-set v1.7.1
	github.com/docker/distribution v2.7.1+incompatible
	github.com/docker/docker-credential-helpers v0.6.3 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7 // indirect
	github.com/go-logr/logr v0.2.1
	github.com/go-test/deep v1.0.7 // indirect
	github.com/gofrs/flock v0.8.0
	github.com/gofrs/uuid v3.3.0+incompatible // indirect
	github.com/google/certificate-transparency-go v1.1.1 // indirect
	github.com/googleapis/gnostic v0.5.3 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.15.2 // indirect
	github.com/hashicorp/consul/api v1.7.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.0
	github.com/hashicorp/go-retryablehttp v0.6.7
	github.com/hashicorp/go-uuid v1.0.2
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
	github.com/jackc/pgx v3.6.2+incompatible // indirect
	github.com/jessevdk/go-flags v1.4.0
	github.com/jinzhu/copier v0.0.0-20201025035756-632e723a6687
	github.com/julienschmidt/httprouter v1.3.0
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/maxbrunsfeld/counterfeiter/v6 v6.3.0
	github.com/nats-io/jwt v1.1.0 // indirect
	github.com/nats-io/nats-server/v2 v2.1.8
	github.com/nats-io/nats.go v1.10.0
	github.com/nats-io/nkeys v0.2.0 // indirect
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d // indirect
	github.com/onsi/ginkgo v1.14.2
	github.com/onsi/gomega v1.10.3
	github.com/opencontainers/image-spec v1.0.1
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.8.0 // indirect
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/spf13/afero v1.4.1 // indirect
	github.com/spf13/cobra v1.1.1
	github.com/tedsuo/ifrit v0.0.0-20191009134036-9a97d0632f00 // indirect
	go.uber.org/multierr v1.6.0
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
	golang.org/x/net v0.0.0-20201027133719-8eef5233e2a1 // indirect
	golang.org/x/sys v0.0.0-20201027140754-0fcbb8f4928c // indirect
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/tools v0.0.0-20201027233111-0b86805daf58 // indirect
	gomodules.xyz/jsonpatch/v2 v2.1.0
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20201026171402-d4b8fe4fd877 // indirect
	google.golang.org/grpc v1.33.1
	gopkg.in/yaml.v2 v2.3.0
	k8s.io/api v0.19.3
	k8s.io/apiextensions-apiserver v0.19.3 // indirect
	k8s.io/apimachinery v0.19.3
	k8s.io/client-go v1.5.1
	k8s.io/code-generator v0.19.3
	k8s.io/klog v1.0.0
	k8s.io/metrics v0.19.3
	k8s.io/utils v0.0.0-20201027101359-01387209bb0d // indirect
	sigs.k8s.io/controller-runtime v0.6.3
)
