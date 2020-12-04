package appmesh

import (
	"context"
	"fmt"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/input"
	"github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/output/appmesh"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-networking/reporting"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-networking/translation/appmesh/traffictarget"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-networking/translation/appmesh/traffictarget/mesh"
)

// the appmesh translator translates an input networking snapshot to an output snapshot of appmesh resources
type Translator interface {
	// Translate translates the appropriate resources to apply input configuration resources for all appmesh meshes contained in the input snapshot.
	// Output resources will be added to the output.Builder
	// Errors caused by invalid user config will be reported using the Reporter.
	Translate(
		ctx context.Context,
		in input.Snapshot,
		appmeshOutputs appmesh.Builder,
		reporter reporting.Reporter,
	)
}

type appmeshTranslator struct {
	trafficTargetTranslator traffictarget.Translator
	meshTranslator          mesh.Translator

	totalTranslates int // TODO(ilackarms): metric
}

func NewAppmeshTranslator() Translator {
	return &appmeshTranslator{
		trafficTargetTranslator: traffictarget.NewTranslator(),
		meshTranslator:          mesh.NewTranslator(),
	}
}

func (t *appmeshTranslator) Translate(
	ctx context.Context,
	in input.Snapshot,
	appmeshOutputs appmesh.Builder,
	reporter reporting.Reporter,
) {
	ctx = contextutils.WithLogger(ctx, fmt.Sprintf("appmesh-translator-%v", t.totalTranslates))

	for _, trafficTarget := range in.TrafficTargets().List() {
		trafficTarget := trafficTarget

		t.trafficTargetTranslator.Translate(ctx, in, trafficTarget, appmeshOutputs, reporter)
	}

	for _, discoveryMesh := range in.Meshes().List() {
		discoveryMesh := discoveryMesh

		t.meshTranslator.Translate(ctx, in, discoveryMesh, appmeshOutputs, reporter)
	}

	t.totalTranslates++
}
