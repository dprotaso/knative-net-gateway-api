module knative.dev/net-gateway-api

go 1.16

require (
	github.com/google/go-cmp v0.5.6
	github.com/gorilla/websocket v1.4.2
	go.uber.org/zap v1.19.1
	google.golang.org/grpc v1.42.0
	k8s.io/api v0.23.5
	k8s.io/apimachinery v0.23.5
	k8s.io/client-go v0.23.5
	k8s.io/code-generator v0.23.5
	k8s.io/utils v0.0.0-20220210201930-3a6ce19ff2f9
	knative.dev/hack v0.0.0-20220506160929-a8076b0510ed
	knative.dev/networking v0.0.0-20220505013701-91dcb36802a9
	knative.dev/pkg v0.0.0-20220505013700-a8b7d99374a3
	sigs.k8s.io/gateway-api v0.4.0
	sigs.k8s.io/yaml v1.3.0
)
