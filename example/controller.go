package example

import (
	"github.com/golang/glog"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
)

type Controller struct {
	// kubeclientset is a standard kubernetes clientset
	kubeclientset kubernetes.Interface
	// workqueue is a rate limited work queue. This is used to queue work to be
	// processed instead of performing it as soon as a change happens. This
	// means we can ensure we only process a fixed amount of resources at a
	// time, and makes it easy to ensure we are never processing the same item
	// simultaneously in two different workers.
	workqueue workqueue.RateLimitingInterface
	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	recorder record.EventRecorder
}

const controllerAgentName = "network-controller"

const (
	// SuccessSynced is used as part of the Event 'reason' when a Network is synced
	SuccessSynced = "Synced"

	// MessageResourceSynced is the message used for an Event fired when a Network
	// is synced successfully
	MessageResourceSynced = "Network synced successfully"
)


func NewController(kubeclientset kubernetes.Interface) *Controller {
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(glog.Infof)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeclientset.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})
	controller := &Controller{
		kubeclientset: kubeclientset,
		workqueue:     workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "Networks"),
		recorder:      recorder,
	}
	glog.Info("Setting up event handlers")
	//networkInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
	//	AddFunc: controller.enqueueNetwork,
	//	UpdateFunc: func(old, new interface{}) {
	//		oldNetwork := old.(*samplecrdv1.Network)
	//		newNetwork := new.(*samplecrdv1.Network)
	//		if oldNetwork.ResourceVersion == newNetwork.ResourceVersion {
	//			// Periodic resync will send update events for all known Networks.
	//			// Two different versions of the same Network will always have different RVs.
	//			return
	//		}
	//		controller.enqueueNetwork(new)
	//	},
	//	DeleteFunc: controller.enqueueNetworkForDelete,
	//})
	return controller
}
