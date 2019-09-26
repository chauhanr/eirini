package k8s_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"

	"code.cloudfoundry.org/eirini/k8s"
	. "code.cloudfoundry.org/eirini/k8s"
	"code.cloudfoundry.org/eirini/k8s/k8sfakes"
	"code.cloudfoundry.org/eirini/metrics"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metricsv1beta1api "k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

var _ = Describe("Metrics", func() {

	const (
		podName1 = "thor-thunder-9000"
		podName2 = "loki-thunder-8000"
	)

	Describe("Collect", func() {

		var (
			podClient        *k8sfakes.FakePodInterface
			podMetricsClient *k8sfakes.FakePodMetricsInterface
			collector        k8s.MetricsCollector
		)

		BeforeEach(func() {
			podClient = new(k8sfakes.FakePodInterface)
			podMetricsClient = new(k8sfakes.FakePodMetricsInterface)
			collector = NewMetricsCollector(podMetricsClient, podClient)
		})

		When("all metrics are valid", func() {
			It("should return them as messages", func() {
				podMetrics := &metricsv1beta1api.PodMetricsList{
					Items: []metricsv1beta1api.PodMetrics{
						createMetrics(podName1),
						createMetrics(podName2),
					},
				}
				podMetricsClient.ListReturns(podMetrics, nil)

				podList := &v1.PodList{
					Items: []v1.Pod{*createPod(podName1), *createPod(podName2)},
				}
				podClient.ListReturns(podList, nil)

				collected, err := collector.Collect()
				Expect(err).ToNot(HaveOccurred())
				Expect(collected).To(ConsistOf(
					metrics.Message{
						AppID:       podName1,
						IndexID:     "9000",
						CPU:         420.5,
						Memory:      430080,
						MemoryQuota: 819200,
						Disk:        42000000,
						DiskQuota:   10,
					},
					metrics.Message{
						AppID:       podName2,
						IndexID:     "8000",
						CPU:         420.5,
						Memory:      430080,
						MemoryQuota: 819200,
						Disk:        42000000,
						DiskQuota:   10,
					},
				))
			})
		})

		When("there are no pods", func() {
			It("should return empty list", func() {
				podClient.ListReturns(&v1.PodList{Items: []v1.Pod{}}, nil)

				podMetrics := metricsv1beta1api.PodMetricsList{
					Items: []metricsv1beta1api.PodMetrics{createMetrics(podName1)},
				}
				podMetricsClient.ListReturns(&podMetrics, nil)

				collected, err := collector.Collect()
				Expect(err).ToNot(HaveOccurred())
				Expect(collected).To(BeEmpty())
			})
		})

		When("there are no metrics", func() {
			It("should return empty list", func() {
				podClient.ListReturns(&v1.PodList{Items: []v1.Pod{*createPod(podName1)}}, nil)

				podMetrics := metricsv1beta1api.PodMetricsList{
					Items: []metricsv1beta1api.PodMetrics{},
				}
				podMetricsClient.ListReturns(&podMetrics, nil)

				collected, err := collector.Collect()
				Expect(err).ToNot(HaveOccurred())
				Expect(collected).To(BeEmpty())
			})
		})

		When("there are no container metrics for a pod", func() {
			It("should skip such pods", func() {
				podClient.ListReturns(&v1.PodList{Items: []v1.Pod{*createPod(podName1), *createPod(podName2)}}, nil)

				podMetrics := metricsv1beta1api.PodMetricsList{
					Items: []metricsv1beta1api.PodMetrics{
						{ObjectMeta: metav1.ObjectMeta{Name: podName1}},
						createMetrics(podName2),
					},
				}
				podMetricsClient.ListReturns(&podMetrics, nil)

				collected, err := collector.Collect()
				Expect(err).ToNot(HaveOccurred())
				Expect(collected).To(ConsistOf(metrics.Message{
					AppID:       podName2,
					IndexID:     "8000",
					CPU:         420.5,
					Memory:      430080,
					MemoryQuota: 819200,
					Disk:        42000000,
					DiskQuota:   10,
				}))
			})
		})

		When("a pod name doesn't have an index (e.g. staging tasks)", func() {
			It("should skip such pod", func() {
				aPodHasNoIndex := "i-dont-have-an-index"

				podClient.ListReturns(&v1.PodList{Items: []v1.Pod{*createPod(aPodHasNoIndex), *createPod(podName2)}}, nil)

				podMetrics := metricsv1beta1api.PodMetricsList{
					Items: []metricsv1beta1api.PodMetrics{createMetrics(aPodHasNoIndex), createMetrics(podName2)},
				}
				podMetricsClient.ListReturns(&podMetrics, nil)

				collected, err := collector.Collect()
				Expect(err).ToNot(HaveOccurred())
				Expect(collected).To(ConsistOf(metrics.Message{
					AppID:       podName2,
					IndexID:     "8000",
					CPU:         420.5,
					Memory:      430080,
					MemoryQuota: 819200,
					Disk:        42000000,
					DiskQuota:   10,
				}))
			})
		})

		When("metrics client returns an error", func() {
			It("should return a wrapped error", func() {
				podMetricsClient.ListReturns(&metricsv1beta1api.PodMetricsList{}, errors.New("oopsie"))

				collected, err := collector.Collect()
				Expect(collected).To(BeEmpty())
				Expect(err).To(MatchError(ContainSubstring("failed to list metrics: oopsie")))
			})
		})
	})

})

var _ = Describe("ForwardMetricsToChannel", func() {
	It("should forward the messages when collector returns them", func() {
		collector := new(k8sfakes.FakeMetricsCollector)
		collector.CollectReturns([]metrics.Message{{AppID: "metric"}}, nil)
		work := make(chan []metrics.Message, 1)
		defer close(work)

		Expect(ForwardMetricsToChannel(collector, work)).To(Succeed())
		Expect(work).To(Receive())
		Expect(work).ToNot(BeClosed())
	})

	It("should return error if collector returns error", func() {
		collector := new(k8sfakes.FakeMetricsCollector)
		collector.CollectReturns(nil, errors.New("oopsie"))
		work := make(chan []metrics.Message, 1)
		defer close(work)

		Expect(ForwardMetricsToChannel(collector, work)).To(MatchError("oopsie"))
		Expect(work).ToNot(Receive())
		Expect(work).ToNot(BeClosed())
	})
})

func createMetrics(podName string) metricsv1beta1api.PodMetrics {
	return metricsv1beta1api.PodMetrics{
		ObjectMeta: metav1.ObjectMeta{Name: podName, Namespace: "opi", ResourceVersion: "10", Labels: map[string]string{"key": "value"}},
		Containers: []metricsv1beta1api.ContainerMetrics{
			{
				Usage: v1.ResourceList{
					v1.ResourceCPU:    resource.MustParse("4205m"),
					v1.ResourceMemory: resource.MustParse("420Ki"),
				},
			},
		},
	}
}

func createPod(podName string) *v1.Pod {
	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: podName,
			Labels: map[string]string{
				"guid": podName,
			},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Resources: v1.ResourceRequirements{
						Limits: v1.ResourceList{
							v1.ResourceMemory: resource.MustParse("800Ki"),
						},
					},
				},
			},
		},
	}
}
