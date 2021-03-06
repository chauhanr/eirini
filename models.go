package eirini

import "errors"

const (
	// Environment Variable Names
	EnvEiriniNamespace                 = "EIRINI_NAMESPACE"
	EnvDownloadURL                     = "DOWNLOAD_URL"
	EnvBuildpacks                      = "BUILDPACKS"
	EnvDropletUploadURL                = "DROPLET_UPLOAD_URL"
	EnvAppID                           = "APP_ID"
	EnvStagingGUID                     = "STAGING_GUID"
	EnvCompletionCallback              = "COMPLETION_CALLBACK"
	EnvEiriniAddress                   = "EIRINI_ADDRESS"
	EnvBuildpackCacheUploadURI         = "BUILDPACK_CACHE_UPLOAD_URI"
	EnvBuildpackCacheDownloadURI       = "BUILDPACK_CACHE_DOWNLOAD_URI"
	EnvBuildpackCacheChecksum          = "BUILDPACK_CACHE_CHECKSUM"
	EnvBuildpackCacheChecksumAlgorithm = "BUILDPACK_CACHE_CHECKSUM_ALGORITHM"
	EnvBuildpackCacheArtifactsDir      = "EIRINI_BUILD_ARTIFACTS_CACHE_DIR"
	EnvBuildpackCacheOutputArtifact    = "EIRINI_OUTPUT_BUILD_ARTIFACTS_CACHE"

	EnvPodName              = "POD_NAME"
	EnvCFInstanceIP         = "CF_INSTANCE_IP"
	EnvCFInstanceIndex      = "CF_INSTANCE_INDEX"
	EnvCFInstanceGUID       = "CF_INSTANCE_GUID"
	EnvCFInstanceInternalIP = "CF_INSTANCE_INTERNAL_IP"
	EnvCFInstanceAddr       = "CF_INSTANCE_ADDR"
	EnvCFInstancePort       = "CF_INSTANCE_PORT"
	EnvCFInstancePorts      = "CF_INSTANCE_PORTS"

	RecipeBuildPacksDir    = "/var/lib/buildpacks"
	RecipeBuildPacksName   = "recipe-buildpacks"
	RecipeWorkspaceDir     = "/recipe_workspace"
	RecipeWorkspaceName    = "recipe-workspace"
	RecipeOutputName       = "staging-output"
	RecipeOutputLocation   = "/out"
	RecipePacksBuilderPath = "/packs/builder"
	BuildpackCacheDir      = "/tmp"
	BuildpackCacheName     = "buildpack-cache"

	AppMetricsEmissionIntervalInSecs = 15

	// Staging TLS:
	CertsMountPath   = "/etc/config/certs"
	CertsVolumeName  = "certs-volume"
	CACertName       = "internal-ca-cert"
	CCAPICertName    = "cc-server-crt"
	CCAPIKeyName     = "cc-server-crt-key"
	EiriniClientCert = "eirini-client-crt"
	EiriniClientKey  = "eirini-client-crt-key"

	RegistrySecretName = "default-image-pull-secret"

	// Certs
	TLSSecretKey  = "tls.key"
	TLSSecretCert = "tls.crt"
	TLSSecretCA   = "ca.crt"

	EiriniCAPath        = "/etc/eirini/certs/ca.crt"
	EiriniCrtPath       = "/etc/eirini/certs/tls.crt"
	EiriniKeyPath       = "/etc/eirini/certs/tls.key"
	CCCrtPath           = "/etc/cf-api/certs/tls.crt"
	CCKeyPath           = "/etc/cf-api/certs/tls.key"
	CCCAPath            = "/etc/cf-api/certs/ca.crt"
	LoggregatorCertPath = "/etc/loggregator/certs/tls.crt"
	LoggregatorKeyPath  = "/etc/loggregator/certs/tls.key"
	LoggregatorCAPath   = "/etc/loggregator/certs/ca.crt"

	CCUploaderSecretName   = "cc-uploader-certs"   //#nosec G101
	EiriniClientSecretName = "eirini-client-certs" //#nosec G101
)

var ErrNotFound = errors.New("not found")

var ErrInvalidInstanceIndex = errors.New("invalid instance index")

type Config struct {
	Properties Properties `yaml:"opi"`
}

type KubeConfig struct {
	Namespace                   string `yaml:"app_namespace"`
	EnableMultiNamespaceSupport bool   `yaml:"enable_multi_namespace_support"`
	ConfigPath                  string `yaml:"kube_config_path"`
}

type Properties struct { //nolint:maligned
	TLSPort       int `yaml:"tls_port"`
	PlaintextPort int `yaml:"plaintext_port"`

	RegistryAddress                  string `yaml:"registry_address"`
	RegistrySecretName               string `yaml:"registry_secret_name"`
	EiriniAddress                    string `yaml:"eirini_address"`
	DownloaderImage                  string `yaml:"downloader_image"`
	UploaderImage                    string `yaml:"uploader_image"`
	ExecutorImage                    string `yaml:"executor_image"`
	AppMetricsEmissionIntervalInSecs int    `yaml:"app_metrics_emission_interval_in_secs"`

	CCTLSDisabled bool `yaml:"cc_tls_disabled"`

	CCCertPath     string
	CCKeyPath      string
	CCCAPath       string
	ClientCAPath   string
	ServerCertPath string
	ServerKeyPath  string

	KubeConfig `yaml:",inline"`

	ApplicationServiceAccount string `yaml:"application_service_account"`
	StagingServiceAccount     string `yaml:"staging_service_account"`

	AllowRunImageAsRoot                     bool `yaml:"allow_run_image_as_root"`
	UnsafeAllowAutomountServiceAccountToken bool `yaml:"unsafe_allow_automount_service_account_token"`

	ServePlaintext bool `yaml:"serve_plaintext"`
}

type EventReporterConfig struct {
	CcInternalAPI string `yaml:"cc_internal_api"`
	CCTLSDisabled bool   `yaml:"cc_tls_disabled"`

	CCCertPath string
	CCKeyPath  string
	CCCAPath   string

	KubeConfig `yaml:",inline"`
}

type RouteEmitterConfig struct {
	NatsPassword        string `yaml:"nats_password"`
	NatsIP              string `yaml:"nats_ip"`
	NatsPort            int    `yaml:"nats_port"`
	EmitPeriodInSeconds uint   `yaml:"emit_period_in_seconds"`

	KubeConfig `yaml:",inline"`
}

type MetricsCollectorConfig struct {
	LoggregatorAddress string `yaml:"loggregator_address"`

	LoggregatorCertPath string
	LoggregatorKeyPath  string
	LoggregatorCAPath   string

	AppMetricsEmissionIntervalInSecs int `yaml:"app_metrics_emission_interval_in_secs"`

	KubeConfig `yaml:",inline"`
}

type TaskReporterConfig struct {
	CCTLSDisabled                bool   `yaml:"cc_tls_disabled"`
	CCCertPath                   string `yaml:"cc_cert_path"`
	CCKeyPath                    string `yaml:"cc_key_path"`
	CAPath                       string `yaml:"ca_path"`
	CompletionCallbackRetryLimit int    `yaml:"completion_callback_retry_limit"`
	TTLSeconds                   int    `yaml:"ttl_seconds"`

	KubeConfig `yaml:",inline"`
}

type StagingReporterConfig struct {
	EiriniCertPath string `yaml:"eirini_cert_path"`
	EiriniKeyPath  string `yaml:"eirini_key_path"`
	CAPath         string `yaml:"ca_path"`

	KubeConfig `yaml:",inline"`
}

type StagerConfig struct {
	EiriniAddress   string
	DownloaderImage string
	UploaderImage   string
	ExecutorImage   string
}

type InstanceIndexEnvInjectorConfig struct {
	ServiceName                string `yaml:"service_name"`
	ServiceNamespace           string `yaml:"service_namespace"`
	ServicePort                int32  `yaml:"service_port"`
	EiriniXOperatorFingerprint string

	KubeConfig `yaml:",inline"`
}
